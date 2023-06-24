package main

import "fmt"

func main() {
	x := 5
	increment(x)

	fmt.Println(x)
}

func increment(x int) {
    x++
}