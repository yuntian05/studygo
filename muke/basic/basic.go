package main

import (
	"fmt"
	"math"
)

func triangle() {
	a, b := 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}
