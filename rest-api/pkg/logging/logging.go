package logging

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

type writerHook struct {
	Writer    []io.Writer
	LogLevels []logrus.Level
}

func (hook *writerHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	for _, w := range hook.Writer {
		w.Write([]byte(line))
	}
	return err
}

func (hook *writerHook) Levels() []logrus.Level {
	return hook.LogLevels
}

var e *logrus.Entry

type Logger struct {
	*logrus.Entry
}

func GetLogger() *Logger {
	return &Logger{e}
}

func (l *Logger) GetLoggerWithField(k string, v interface{}) *Logger {
	return &Logger{l.WithField(k, v)}
}

func init() { // есть в logrus entry
	l := logrus.New()
	l.SetReportCaller(true)
	l.Formatter = &logrus.TextFormatter{
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			filename := path.Base(frame.File)
			return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		},
		FullTimestamp: true,
	}

	err := os.MkdirAll("logs", 0o744) // создаем директорию для записи лога
	if err != nil {
		panic(err)
	}

	allFile, err := os.OpenFile(
		"logs/all.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0o740,
	) // создаем файл для записи
	if err != nil {
		panic(err)
	}

	// для записи в несколько мест
	l.SetOutput(io.Discard) // по умолчанию логи никуда не уходят /dev/null
	l.AddHook(
		&writerHook{ // добавление хука, хук это ф-ция которая вызывается до или после создания/удаления/обновления
			Writer:    []io.Writer{allFile, os.Stdout},
			LogLevels: logrus.AllLevels,
		},
	)

	l.SetLevel(logrus.TraceLevel)

	e = logrus.NewEntry(l)
}
