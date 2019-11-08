package main

import (
	"fmt"
)

func fibonacci1(n int) int {

	a := []int{0, 1}

	for i := 2; i <= n; i++ {

		a = append(a, a[i - 1] + a[i - 2])

	}

	return a[n]

}

func main() {
	
	fmt.Println(fibonacci1(7))
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89 ..
}