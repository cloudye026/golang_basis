package main

import (
	"fmt"

	"08-go_docs/hello/morestrings"

	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("hello,world")

	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))

	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
