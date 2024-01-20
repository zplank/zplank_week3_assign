package main

import (
	"fmt"

	"gonum.org/v1/gonum/stat"
)

func correlation(x, y []float64) float64 {
	return stat.Correlation(x, y, nil)
}

func main() {
	// correlation between x4 and y4
	x4 := []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}
	y4 := []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89}

	// calculate correlation
	corr := correlation(x4, y4)

	// print correlation coefficient
	fmt.Printf("Correlation coefficient between x4 and y4: %v\n", corr)
}
