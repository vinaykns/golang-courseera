package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main () {
	scanner := bufio.NewScanner(os.Stdin)
	data := make([]int, 0)
	for scanner.Scan() {
		s := scanner.Text()
		if s != "X" {
			if len(data) > 3 {
				intData, _ := strconv.Atoi(s)
				data = append(data, intData)
				sort.Ints(data)
				fmt.Println(data)
			} else {
				intData, _ := strconv.Atoi(s)
				data = append(data, intData)
				sort.Ints(data)
				fmt.Println(data)
			}
		} else {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}
}
