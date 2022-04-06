package main

import (
	"fmt"
)
func main() {
	gradeAvg := 89
 
	if gradeAvg >= 90 {
	fmt.Println("You've earned an A!")
	} else if gradeAvg >= 80 {
	fmt.Println("You've earned a B.")
	} else if gradeAvg >= 70 {
	fmt.Println("You've earned a C?") 
	} else if gradeAvg >= 60 { 
	fmt.Println("You've earned a D...")
	} else {
	fmt.Println("An F???")
	}
}