1. Интерфейс - спец. тип в go, представляющий из себя набор сигнатур методов,
которые нужно реализовать в его имплементации

2. Встраивание интерфейсов (interface embedding)

3. Явной имплементации нет, нужно просто иметь такие же методы.
Утиная типизация (duck typing). Т.е. в коде в реализации интерфейса не нужно явно 
указывать интерфейс, который реализизуется. Связывание реализации и интерфейса 
происходит во время компиляции.
// объявление интерфейса
type Printer interface {
    Print()
}

// нигде не указано, что User реализует интерфейс Printer
type User struct {
    email string
}

// структура User имеет метод Print, как в интерфейсе Printer
// Следовательно, во время компиляции запишется связь между User и Printer
func (u *User) Print() {
    fmt.Println("My email is", u.email)
}

// функция принимает как аргумент интерфейс Printer
func TestPrint(p Printer) {
    p.Print()
}

func main() {
    // в функцию TestPrint передается структура User,
    // и так как она реализует интерфейс Printer, все работает без ошибок
    TestPrint(&User{email: "test@test.com"})
}

4. Интерфейсное значение внутри хранит информацию о конкретном (неинтерфейсном) типе
и его значении
                                     ConcreateType              ConcreateType
                                    /                           /
interfaceValue = nil    interfaceValue             interfaceValue
                                    \                           \
                                    nil                         Concreate Value 

5. Nil interface: нет ни типа, ни значения. Паникует (падает) при вызове методов

6. Пустой интерфейс: interface{} - содержит любое значение, т.к. каждый реализует
набор из 0 методов

7. Полиморфизм

8. Утверждение типа (type assertion) - пытаемся получить значение конкретного типа
concreateValue := interfaceValue.(concreateType)

Заполните пропуск, чтобы структура OverflowErr реализовывала интерфейс error

type OverflowErr struct {
    msg string
}

func (e *OverflowErr) error() string {
    return e.msg
}
