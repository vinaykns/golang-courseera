package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	input = flag.String("input", "empty", ".go run -input ian")
	test  = flag.Bool("test", false, "go run -test")
)

func main() {
	flag.Parse()

	if *test == true {
		testInputs()
		os.Exit(0)
	}

	if *input == "empty" {
		fmt.Println("Please make sure to enter -input ian or similar next run")
		os.Exit(1)
	}

	str := strings.ToLower(*input)
	fmt.Println(findian(str))
}

// findian finds if input str starts with prefix i, ends with suffix n and has a character a in the middle
func findian(str string) (result string) {
	if strings.HasPrefix(str, "i") && strings.Contains(str, "a") && strings.HasSuffix(str, "n") {
		return "Found!"
	} else {
		return "Not Found!"
	}
}

func testInputs() {
	inputs := []string{"ianAIN", "ian", "Ian", "iuiygaygn", "I d skd a efju N", `d
skd
a
efju
N`, `i
a
N`, "iaaaan", "ihhhhhn", "ina", "xian"}

	// outputs contains the results of the corresponding input strings.
	outputs := []string{"Not Found!", "Found!", "Found!", "Found!", "Found!", "Not Found!",
		"Found!", "Found!", "Not Found!", "Not Found!", "Not Found!"}

	var result string
	for i, each := range inputs {
		result = findian(strings.ToLower(each))
		if outputs[i] != result {
			fmt.Println("tests not passed.")
			return
		}
	}
	fmt.Println("tests passed.")
}
