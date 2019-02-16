package main

import "fmt"

func main() {
	fmt.Printf("Sum: %d\n", sum(42, 42, 42, 96, 102, 3))
}

func sum(vals ...int) int {
	total := 0
	for _, val :=  range vals {
		total += val
	}
	return total
}