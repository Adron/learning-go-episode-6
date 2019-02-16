package main

import "fmt"

func main() {
	fmt.Printf("Sum: %d\n", sum(42, 42, 42, 96, 102, 3))

	theNumbers := []int{42, 51, 23, 62, 102, 9, 23, 64}

	sum, sentence := sumAndPresent(theNumbers,
		"Hello, ", "cheeseheads", " of the world!")

	fmt.Printf("Sum: %d\nSentence: %s", sum, sentence)
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func sumAndPresent(vals []int, words ...string) (int, string) {
	total := sum(vals...)
	sentences := "To say: "
	for _, word := range words {
		sentences += word
	}
	return total, sentences
}
