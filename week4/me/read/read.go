package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Name struct stores the first and last name of a persion
type Name struct {
	firstName string
	lastName string
}


// this is how one can have multiple comments in golang.
/*
var names = []Name{
	Name{
		firstName: "vinay",
		lastName:  "kumar",
	},
}
 */

func addDataToSlice(input *[]Name, data Name) {
	*input = append(*input, data)
}

// getDataFromFile captures the data from the given file Name.
func getDataFromFile(name string) (result []Name, err error) {
	fileData, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("failed to read data from file with error: %v", err)
	}
	names := strings.Split(string(fileData), "\n")
	for _, each := range names {
		names := strings.Split(each, " ")
		firstName := names[0]
		lastName := names[1]
		data := Name{firstName: firstName, lastName:lastName}
		addDataToSlice(&result, data)
	}
	return result, nil
}

// getFileNameFromUser obtains the file Name from user.
func getFileNameFromUser() (fileName string, err error) {
	fmt.Println("Enter the input file")
	reader := bufio.NewReader(os.Stdin)
	input, _, err := reader.ReadLine()
	fileName = string(input)
	if err != nil {
		return "", fmt.Errorf("failed to get the input from user with error: %v", err)
	}
	return fileName, nil
}

func main() {
	fileName, err := getFileNameFromUser()
	if err != nil {
		os.Exit(1)
	}
	result, err := getDataFromFile(fileName)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(result)
}
