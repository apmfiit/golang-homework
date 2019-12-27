package main

import (
	"fmt"
)

func unique_count(a *[]int) int {
	var m = make(map[int]bool)
	n := 0
	for i := 0; i < len((*a)); i++ {
		if !m[(*a)[i]] {
			m[(*a)[i]] = true
			n = n + 1
		}
	}
	return n
}

func main() {
	a := []int{3, 3, 3, 4, 4, 1}
	fmt.Println(unique_count(&a))
} 