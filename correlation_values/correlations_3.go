package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
)

func correlation(x, y []float64) float64 {
	return stat.Correlation(x, y, nil)
}

func main() {
	// correlation between x3 and y3
	x3 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y3 := []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}

	// calculate correlation
	corr := correlation(x3, y3)

	// print correlation coefficient
	fmt.Printf("Correlation coefficient between x3 and y3: %v\n", corr)
}
