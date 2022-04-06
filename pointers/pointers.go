package main

import (
	"fmt"
)

func main() {
	fmt.Println("");
	a := 33
	b := &a
	c := &b
	d := &c
	e := &d
	f := &a
	fmt.Println("How Pointers Work")
	for i := 0; i < 50; i++ {
		fmt.Print("🗿")
	}
	fmt.Print("\n")
	fmt.Printf("a = %d\n", a)
	fmt.Printf("&a = %p\n", &a)
	fmt.Printf("b = %p\n", b)
	fmt.Printf("*b = %d\n", *b)
	fmt.Printf("c = %p\n",c)
	fmt.Printf("*c = %p\n", *c)
	fmt.Printf("**c = %d\n",**c)
	fmt.Printf("d = %p\n", d)
	fmt.Printf("*d = %p\n",*d)
	fmt.Printf("**d = %p\n", **d)
	fmt.Printf("***d = %d\n", ***d)
	fmt.Printf("e = %p\n", e)
	fmt.Printf("*e = %p\n", *e)
	fmt.Printf("**e = %p\n", **e)
	fmt.Printf("***e = %p\n", ***e)
	fmt.Printf("****e = %d\n", ****e)
	fmt.Printf("f = %p\n", f);
	fmt.Printf("*f = %d\n", *f);
}