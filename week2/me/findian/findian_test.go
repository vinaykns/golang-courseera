package main

import (
	"fmt"
	"testing"
	"strings"
)

var (
	// input strings that need to be tested.
	inputs = []string{"ianAIN", "ian", "Ian", "iuiygaygn", "I d skd a efju N", `d
skd
a
efju
N`, `i
a
N`, "iaaaan", "ihhhhhn", "ina", "xian"}

	// outputs contains the results of the corresponding input strings.
	outputs = []string{"Found!", "Found!", "Found!", "Found!", "Found!", "Not Found!",
		"Found!", "Found!", "Not Found!", "Not Found!", "Not Found!"}
)

// TestFindian runs to test the function findian.
func TestFindian(t *testing.T) {
	for i, each := range inputs {
		result := findian(strings.ToLower(each))
		if outputs[i] == result {
			continue
		} else {
			fmt.Println("tests failed")
		}
	}
	fmt.Println("tests passed")
}