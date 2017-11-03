package main

import "fmt"

func main() {
	var a int = 12
	var b int = 45
	swap(&a, &b)
	fmt.Printf("a=%d\n", a)
	fmt.Printf("b=%d\n", b)
}

func swap(a *int, b *int) {
	var t int
	t = *a
	*a = *b
	*b = t
}
