package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Program is ending")
	fmt.Println("Program booting up...")
	fmt.Println("Hello user! What's your name?")
	var name string
	fmt.Scan(&name)
	fmt.Println("Enter the amount of money you have")
	var money float64
	fmt.Scan(&money)
	if money < 2.50 {
		fmt.Println("I was going to ask you if you if you wanted to buy an ice cream, but nvm, you're too broke.")
	} else {
		fmt.Print("Do you want to buy an ice cream, ", name, "? Cost: $2.50\n")
		var decision string
		fmt.Scan(&decision)
		if decision == "Yes" {
			fmt.Println("Here you go!")
			money -= 2.50
			fmt.Println("Remaining money: ", money)
		}
	}
	
}