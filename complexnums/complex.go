package main

import (
	"fmt"
)
func main() {
	var complexNum complex64 = complex(3, -2);
	fmt.Printf("complex number: %v \ntype: %T\n", complexNum, complexNum);
	var ru rune = '🗿';
	
	fmt.Println("rune: ", ru);
	fmt.Println("rune content: ", string(ru));
}