package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

)

var (
	input  = flag.String("input", "empty", ".go run -input ian")
	test = flag.Bool("test", false, "go run -test")
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
	if strings.HasPrefix(str, "i") && strings.Contains(str, "a") && strings.HasSuffix(str, "n") {
		fmt.Printf("%s Found!\n", str)
	} else {
		fmt.Printf("%s Not Found!\n", str)
	}

}

func testInputs() {
	inputs := []string{"ianAIN", "ian", "Ian", "iuiygaygn", "I d skd a efju N", `d
skd
a
efju
N`, `i
a
N`,"iaaaan", "ihhhhhn", "ina", "xian"}

	for _, each := range inputs {
		str := strings.ToLower(each)
		if strings.HasPrefix(str, "i") && strings.Contains(str, "a") && strings.HasSuffix(str, "n") {
			fmt.Printf("%s Found!\n", each)
		} else {
			fmt.Printf("%s Not Found!\n", each)
		}
		fmt.Println("------------------------------")
	}
}