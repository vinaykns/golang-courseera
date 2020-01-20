package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Enter the string")
	var data string
	fmt.Scanf("%q", &data)
	data = strings.ToLower(data)
	if string(data[0]) == "i" &&  string(data[len(data)-1]) == "n" &&
		strings.Contains(data[1:len(data)-1], "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}