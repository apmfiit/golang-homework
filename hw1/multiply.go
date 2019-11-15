package main

import (
	"fmt"
)

func multiply(a, b float64){
	var c float64
	var i float64
	
	for {
		if b == 0 || a == 0 {
		c = 0
		fmt.Println(c)
		break
		}

		if b > 0 {
		for i = 1; i <= b; i++ {
			c += a
		}
		fmt.Println(c)
		break
		}
		if b < 0 {
		for i = -1; i >= b; i-- {
			c += a
		}
		fmt.Println(-c)
		break
		}
	}
}
	

func main(){
	
	multiply(0, -12)
}


