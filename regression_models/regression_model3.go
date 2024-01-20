package main

import (
	"fmt"

	"github.com/sajari/regression"
)

func main() {
	// x3/y3 data
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73}

	// create linear regressio model
	model3 := new(regression.Regression)

	// add data to model
	for i, xi := range x {
		model3.Train(regression.DataPoint(y[i], []float64{xi}))
	}

	// train model
	model3.Run()

	// get coeefficients
	intercept, slope := model3.Coeff(0), model3.Coeff(1)

	fmt.Printf("Intercept: %.4f\n", intercept)
	fmt.Printf("Slope: %.4f\n", slope)
}
