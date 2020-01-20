package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Printf("Enter the floating point number \n")
	reader := bufio.NewReader(os.Stdin)
	userData, err := reader.ReadString('\n')
	userData = strings.Trim(userData, "\n")
	if err != nil {
		fmt.Errorf("failed to reader input from user %v", err)
	}
	if floatData, err := strconv.ParseFloat(userData, 64); err == nil {
		intData  := int(floatData)
		fmt.Printf("The value after truncation is %d \n", intData)
	} else {
		fmt.Printf("failed to convert to integer %v", err)
	}
}
