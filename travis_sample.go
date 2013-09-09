package main

import (
	"fmt"
)

func main() {
	fmt.Printf("1 + 1 = %v\n", Add(1,1))
}

func Add(n int, m int) int {
	return n + m
}
