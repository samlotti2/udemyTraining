package main

import "fmt"

func b(v int) {
	count := v
	for bug := 0; bug < count; bug++ {
		fmt.Println(v)
	}
}

func main() {
	var a = 2
	b(a)
	fmt.Println("Hello world")
}
