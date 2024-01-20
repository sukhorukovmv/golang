package main

import (
	"fmt"
)

func main() {
    fmt.Println("test")
    var intPointer *int //объявить переменную которая имеет тип указатель на инт, default value for pointer is nil
    fmt.Printf("%T %#v \n", intPointer, intPointer)
    // Лучше не объявлять указатели без значения
    var a int64 = 7
    fmt.Printf("%T %#v \n", a, a)

    var pointerA *int64 = &a //& получить адрес переменной а. // *int64 можно не указывать
    fmt.Printf("%T %#v %#v \n", pointerA, pointerA, *pointerA)

    //get pointer via new keyword
    var newPointer = new(float32) // new позволяет создать указатель с не nil значением, начальное значение будет на основе типа указателя
    fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)
    *newPointer = 3
    fmt.Printf("%T %#v %#v \n", newPointer, newPointer, *newPointer)

    println()
    num := 3
    squarePointer(&num)
    println(num)

}


func squarePointer( n *int ) {
    *n *= *n

}
