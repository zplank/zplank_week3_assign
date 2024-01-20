package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
)

func correlation(x, y []float64) float64 {
	return stat.Correlation(x, y, nil)
}

func main() {
	// correlation between x1 and y1
	x1 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y1 := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	// calculate correlation
	corr := correlation(x1, y1)

	// print correlation coefficient
	fmt.Printf("Correlation coefficient between x1 and y1: %v\n", corr)
}
