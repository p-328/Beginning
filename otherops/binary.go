package main

import (
	"fmt"
)

func main() {
	a := 44 // 101101
	b := 3 // 100001
	fmt.Printf("%d\n", a & b) // 100000
	fmt.Printf("%d\n", a | b) // 101101
	fmt.Printf("%d\n", a ^ b) // 001100
	fmt.Printf("%d\n", a &^ b) // 010010
	fmt.Printf("%d\n", a >> b)
	fmt.Printf("%d\n", a << b)
	fmt.Printf("%d\n", b >> a)
	fmt.Printf("%d\n", b << a)
}