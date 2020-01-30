package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	start := time.Now()
	for i := 0; i < 10; i++ {
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
	fmt.Println(time.Since(start))
}

func main() {
	for i := 0; i < 10; i++ {
		//go f(i)
		f(i)
	}
	var input string
	fmt.Scanln(&input)
}
