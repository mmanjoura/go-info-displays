package main

import (
	"fmt"
)

func vmap(value, low1, high1, low2, high2 float64) float64 {
    return low2 + (high2-low2)*(value-low1)/(high1-low1)
}

func main() {

	datax := 15.0
	datay := 5000.0

	x := vmap(datax, 10, 20, 0, 767)
	y := vmap(datay, 0, 10000, 0, 1023)
	
	fmt.Printf("(%.2f, %.2f) -> (%.2f, %.2f)\n", datax, datay, x, y)
}

