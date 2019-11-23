package main

import (
	"fmt"
	//"sort"
)

func bubbleSort(array *[]float64) {
	for i := 0; i < len(*array); i++ {
		for j := 0; j < len(*array)-1-i; j++ {
			if (*array)[j] < (*array)[j+1] {
				(*array)[j], (*array)[j+1] = (*array)[j+1], (*array)[j]
			}
		}
	}
}

func main() {
	array:= []float64{1,2,-3,55,23,12}
	bubbleSort(&array)
	fmt.Println(array)
}