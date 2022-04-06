package main

import (
	"fmt"
)

import (
	"io/ioutil"
)
func main() {
	html, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(html))
	}
}