package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
)

func correlation(x, y []float64) float64 {
	return stat.Correlation(x, y, nil)
}

func main() {
	// correlation between x2 and y2
	x2 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y2 := []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}

	// calculate correlation
	corr := correlation(x2, y2)

	// print correlation coefficient
	fmt.Printf("Correlation coefficient between x2 and y2: %v\n", corr)
}
