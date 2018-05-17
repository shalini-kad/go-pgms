
package main

import "fmt"


func counter() func() int {
    c := 0
    return func() int {
		c++
		return c
    }
}

func main() {

    increment := counter()
    fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println("New counter")
	newc:=counter()
	fmt.Println(newc())
	fmt.Println(newc())
}