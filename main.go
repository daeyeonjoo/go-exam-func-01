package main

import (
	"fmt"
)

type fooFunc func()

func f1() {
	fmt.Println("f1 fired")
}

func f2() {
	fmt.Println("f2 fired")
}

func main() {
	var l = []fooFunc{
		f1, f2,
	}

	for i := 0; i < len(l); i++ {
		l[i]()
	}
}
