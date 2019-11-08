package main

import (
	"fmt"
)

func multiply(a, b float64){
	var c float64
	var i float64
	

	if b > 0 {

		for i = 1; i <= b; i++ {
			c += a
		}
		fmt.Println(c)
	}

	if b < 0 {
		for i = -1; i >= b; i-- {
			c += a
		}
		fmt.Println(-c)
	}
}

func main(){
	
	multiply(5, -40)
}


