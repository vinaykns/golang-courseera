package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// createJsonObject Marshals into a Json Object with the fields name and address
func createJsonObject(name, address string) ([]byte, error) {
	m := make(map[string]string)
	m["name"] = name
	m["address"] = address
	result, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("failed to Marshal the Json data with error: %v", err)
	}
	return result, nil
}

// takeInput interacts with user to take name and address data
// TODO: can talk about how the modifications are done to take
//  input including with space.(i.e adding deliminator new line)
func takeInput() (name, address string, err error) {
	fmt.Println("Enter the name")
	in := bufio.NewReader(os.Stdin)
	name, err = in.ReadString('\n')
	if err != nil {
		return "", "", fmt.Errorf("failed to read the string: %v", err)
	}
	name = strings.Trim(name, "\n")
	fmt.Println("Enter the address")
	address, err = in.ReadString('\n')
	if err != nil {
		return "", "", fmt.Errorf("failed to read the string: %v", err)
	}
	address = strings.Trim(address, "\n")
	return
}

// this will return a json object for the fields name and address.
func main() {
	name, address, err := takeInput()
	if err != nil {
		os.Exit(1)
	}
	result, err := createJsonObject(name, address)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(result))
}
