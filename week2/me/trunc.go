package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// truncateFpNumber truncates the floating point number
// to display only the integer part in it.
func truncateFpNumber() (err error) {
	fmt.Printf("Enter the floating point number \n")
	reader := bufio.NewReader(os.Stdin)
	userData, _, err := reader.ReadLine()
	if err != nil {
		return fmt.Errorf("failed to read input from user %v", err)
	}
	if floatData, err := strconv.ParseFloat(string(userData), 64); err == nil {
		intData  := int(floatData)
		fmt.Printf("The value after truncation is %d \n", intData)
		return nil
	} else {
		return fmt.Errorf("failed to truncate the input with error: %v", err)
	}
}

func main()  {
	err := truncateFpNumber()
	if err != nil {
		fmt.Printf("failed with error: %v", err)
		os.Exit(1)
	}
}