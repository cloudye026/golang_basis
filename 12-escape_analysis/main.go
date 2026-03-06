package main

import "fmt"

func main() {
	foo()
}

func foo() {
	a := 10
	go func() {
		fmt.Println(a)
	}()
}
