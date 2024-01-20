Какая-то функция возвращает критичные и некритичные ошибки:

// некритичная ошибка валидации
type NonCriticalError struct{}
 
func (e NonCriticalError) Error() string {
    return "validation error"
}
 
// критичные ошибки
var (
    ErrBadConnection = errors.New("bad connection")
    ErrBadRequest    = errors.New("bad request")
)
Реализуйте функцию GetErrorMsg(err error) string, которая возвращает текст ошибки, если она критичная. Если ошибка некритичная, то возвращается пустая строка. В случае неизвестной ошибки возвращается строка unknown error:

GetErrorMsg(errors.New("bad connection")) // "bad connection"
GetErrorMsg(errors.New("bad request")) // "bad request"
GetErrorMsg(NonCriticalError{}) // ""
GetErrorMsg(errors.New("random error")) // "unknown error"
