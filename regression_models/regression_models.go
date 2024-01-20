package main

import (
	"fmt"

	"github.com/sajari/regression"
)

func main() {
	// define data
	data := [][]float64{
		{10, 8.04, 10, 9.14, 10, 7.46, 8, 6.58},
		{8, 6.95, 8, 8.14, 8, 6.77, 8, 5.76},
		{13, 7.58, 13, 8.74, 13, 12.74, 8, 7.71},
		{9, 8.81, 9, 8.77, 9, 7.11, 8, 8.84},
		{11, 8.33, 11, 9.26, 11, 7.81, 8, 8.47},
		{14, 9.96, 14, 8.1, 14, 8.84, 8, 7.04},
		{6, 7.24, 6, 6.13, 6, 6.08, 8, 5.25},
		{4, 4.26, 4, 3.1, 4, 5.39, 19, 12.5},
		{12, 10.84, 12, 9.13, 12, 8.15, 8, 5.56},
		{7, 4.82, 7, 7.26, 7, 6.42, 8, 7.91},
		{5, 5.68, 5, 4.74, 5, 5.73, 8, 6.89},
	}

	// regression for each xi yi
	for i := 0; i < len(data[0]); i += 2 {
		x := make([]float64, len(data))
		y := make([]float64, len(data))

		// Extract x and y values
		for j := 0; j < len(data); j++ {
			x[j] = data[j][i]
			y[j] = data[j][i+1]
		}

		// create regression model
		model := new(regression.Regression)

		// Add data points to the model
		for j := 0; j < len(x); j++ {
			model.Train(regression.DataPoint(y[j], []float64{x[j]}))
		}

		// Run the regression
		model.Run()

		// Print the results
		fmt.Printf("Regression Analysis for y%d ~ x%d:\n", i/2+1, i/2+1)
		fmt.Println("Coefficients:", model.Coeff)
		fmt.Println("R-squared:", model.R2)
		fmt.Println()
	}
}
