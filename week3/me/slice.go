package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main () {
	PrintSlice()
}

// sliceisFull checks if data slice is full, Returns true
// if full else false.
func sliceisFull(data []int) bool {
	for _, each := range data {
		if each == 0 {
			return false
		}
	}
	return true
}

// addDataToSlice adds the data to the input slice and sorts
// it.
func addDataToSlice(input *[]int, data string) (result *[]int) {
	intData, _ := strconv.Atoi(data)
	*input = append(*input, intData)
	sort.Ints(*input)
	return input
}

// locateIndexAndAdd will find an empty position in the slice and
// replace it with the data value.
func locateIndexAndAdd(input []int, data string) (result []int) {
	intData, _ := strconv.Atoi(data)
	for j, each := range input {
		if each == 0 {
			input[j] = intData
			break
		}
	}
	sort.Ints(input)
	return input
}

// PrintSlice takes input from user until X is entered.
// this function appends data to the input slice if no of
// elements entered are more than three, else find an empty
// position in the slice.
// TODO: Show the illustration what happens if I don't do the pass by
//  reference in the addDataToSlice function.
func PrintSlice () {
	scanner := bufio.NewScanner(os.Stdin)
	data := make([]int, 3)
	var result *[]int
	for scanner.Scan() {
		s := scanner.Text()
		if s != "X" {
			if sliceisFull(data) {
				result = addDataToSlice(&data, s)
				fmt.Println(*result)
			} else {
				fmt.Println(locateIndexAndAdd(data, s))
			}
		} else {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}