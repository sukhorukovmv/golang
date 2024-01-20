package main

import (
	"fmt"
)

type Runner interface {
    Run() string
}

type Swimmer interface {
    Swim() string
}

type Flyer interface {
    Fly() string
}

type Ducker interface {
    Runner
    Swimmer
    Flyer
}

type Human struct { //имплементирует интерфейс Runner
    Name string 
}

// т.е. значения типа Human можем подставить наши интерфейсные значения

func (h Human) Run() string { // этого достаточно чтобы заимплементить интерфейс раннер
    return fmt.Sprintf("Human %s is run\n", h.Name)
}
func (h Human) WriteCode() string { // этого достаточно чтобы заимплементить интерфейс раннер
    return fmt.Sprintf("\n\nHuman %s write code\n\n", h.Name)
}

type Duck struct {
    Name, Surname string
}

func (d Duck) Run() string {
    return "The duck is run\n"
}

func (d Duck) Swim() string {
    return "The duck is swim\n"
}

func (d Duck) Fly() string {
    return "The duck is flyer\n"
}

func main() {
    interfaceValues()
}

func polymorphism(runner Runner){
    fmt.Printf(runner.Run())
}

func typeAssertion(runner Runner) {
    fmt.Printf("Type: %T Value: %#v\n", runner, runner)
    if human, ok := runner.(*Human); ok {
        fmt.Printf("Type: %T Value: %#v\n", human, human)
        answer := human.WriteCode()
        println(answer)
    }

    if duck, ok := runner.(*Duck); ok {
        fmt.Printf("Type: %T Value: %#v\n", duck, duck)
        answer := duck.Fly()
        println(answer)
    }

    switch v := runner.(type) {
    case *Human:
        fmt.Println(v.WriteCode())
        fmt.Println(v.Run())
    case *Duck:
        fmt.Println(v.Run())
    default:
        fmt.Printf("Type: %T Value: %#v\n", v, v)
    }
}
func interfaceValues() {
    var runner Runner
    fmt.Printf("Type: %T Value: %#v\n", runner, runner)

    if runner == nil {
        fmt.Printf("Runner is nil\n")
    }
    // runner = int64(1) // error type int64 not implemet interface Runner

    var unnamedRunner *Human
    fmt.Printf("Print * указатель на HUMAN: %T Value: %#v\n", unnamedRunner, unnamedRunner)

    runner = unnamedRunner
    fmt.Printf("Присвоим интерфейсу runner указатель на human: %T Value: %#v\n", runner, runner)
    if runner == nil {
        fmt.Printf("Runner is nil\n")
    }

    namedRunner := &Human{Name: "Jack"}
    fmt.Printf("Named Runner is: %T Value: %#v\n", namedRunner, namedRunner)

    // Для чего вообще нунжны значения интерфейсного типа, что с ними делать ? 

    var runner1 Runner
    fmt.Printf("Type: %T Value: %#v\n", runner1, runner1)

    john := &Human{"John"}
    runner1 = john
    polymorphism(john)

    println("\n\n")
    typeAssertion(john)
    println("\n\n")

    donald := &Duck{Name: "Donald", Surname: "Duck"}
    runner1 = donald
    polymorphism(donald)

    println("\n\n")
    typeAssertion(donald)
    println("\n\n")

}
