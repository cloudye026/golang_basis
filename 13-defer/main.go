package main

import (
	"os"
	"fmt"
)

func main() {
	f, err := os.Open("./1.text")
	if err != nil {
		panic(err)
	}
	if f != nil {
		defer f.Close()
	}
	fmt.Println(f, err)
}