package main

import (
	"fmt"
)
import (
	"math"
)
func main() {
	var num float64;
	fmt.Println("Enter a number: ");
	fmt.Scan(&num);
	var secondNum float64;
	fmt.Println("Enter a second number: ");
	fmt.Scan(&secondNum);
	var operand string;
	fmt.Println("Enter an operand: ");
	fmt.Scan(&operand);
	if (operand == "+") {
		fmt.Println(num + secondNum);
	} else if (operand == "-") {
		fmt.Println(num - secondNum);
	} else if (operand == "*") {
		fmt.Println(num * secondNum);
	} else if (operand == "/") {
		fmt.Println(num / secondNum);
	} else if (operand == "%") {
		result := math.Mod(num, secondNum);
		fmt.Println(result);
	} else {
		fmt.Println("Invalid operand");
	}
	
}