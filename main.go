package main

import (
	"fmt" // biblioteca usada para entrada/saida
	test "helloWorld/src"
	s "./anotherTest"
)

func main() {
 	fmt.Println("Hello world from Go!")

    a := 12 // inferencia de tipos
    b := false
    c := 12.5
    d := "andressa"
    e := 'A'

    fmt.Printf("%v \n", a) // imprimindo valor da variável
    fmt.Printf("%v \n", b)
    fmt.Printf("%v \n", c)
    fmt.Printf("%v \n", d)

    fmt.Printf("%T \n", a) // imprimindo tipo da variável
    fmt.Printf("%T \n", b)
    fmt.Printf("%T \n", c)
    fmt.Printf("%T \n", d)
    fmt.Printf("%T \n", e)

	test.TestExample()
    SayHello()
}
