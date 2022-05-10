package main

import (
	"fmt"
)

/*
	Task 6. Implement a function called uniqueFruits, that passing two slices of
	fruits, it will return a slice containing the fruits that appear in
	either or both slices. The returned slice should have no duplicates.
*/

func main() {
	fruits1 := []string{"apple", "melon", "banana", "melon", "grape"}
	fruits2 := []string{"apple", "grape", "grape", "banana"}

	fmt.Println(uniqueFruits(fruits1, fruits2))
}

func uniqueFruits(fruits1, fruits2 []string) []string {
	result := []string{}
	for _, f := range fruits1 {
		if !contains(f, fruits2) && !contains(f, result) {
			result = append(result, f)
		}
	}
	for _, f := range fruits2 {
		if !contains(f, fruits1) && !contains(f, result) {
			result = append(result, f)
		}
	}

	return result
}

func contains(fruit string, otherFruits []string) bool {
	for _, f := range otherFruits {
		if fruit == f {
			return true
		}
	}
	return false
}
