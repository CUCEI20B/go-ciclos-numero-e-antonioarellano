package main

import (
	"fmt"
)

func main() {
	var x, e, i float64
	i = 0
	fmt.Scanln(&x)
	for i <= x {
		e = e + (1 / factorial(i))
		i++
	}
	fmt.Println(e)
}

func factorial(n float64) float64 {
	if n < 0 {
		return 0
	} else {
		if n > 1 {
			return n * factorial(n-1)
		}
		return 1
	}
}
