package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Enter the string")
	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadLine()
	data := string(input)
	if err != nil {
		os.Exit(1)
	}
	if len(data) > 0 && data[0] == '"' {
		data = data[1:]
	}
	if len(data) > 0 && data[len(data)-1] == '"' {
		data = data[:len(data)-1]
	}
	data = strings.ToLower(data)
	fmt.Println(findian(data))
}

// findian finds if string starting with i, ending with n and contains a in between.
func findian(data string) (result string) {
	if strings.HasPrefix(data, "i") && strings.HasSuffix(data, "n") &&
		strings.Contains(data, "a") {
		return "Found!"
	} else {
		return "Not Found!"
	}
}
