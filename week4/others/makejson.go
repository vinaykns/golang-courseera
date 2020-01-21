package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	user := make(map[string]string)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Introduce your name:")
	name, _, err := reader.ReadLine()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	user["name"] = string(name)

	fmt.Println("Introduce your address:")
	address, _, err := reader.ReadLine()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	user["address"] = string(address)

	userJSON, err := json.Marshal(user)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(userJSON))
}