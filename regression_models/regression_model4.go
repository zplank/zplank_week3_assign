package main

import (
	"fmt"

	"github.com/sajari/regression"
)

func main() {
	// x4/y4 data
	x := []float64{8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8}
	y := []float64{6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89}

	// create linear regressio model
	model4 := new(regression.Regression)

	// add data to model
	for i, xi := range x {
		model4.Train(regression.DataPoint(y[i], []float64{xi}))
	}

	// train model
	model4.Run()

	// get coeefficients
	intercept, slope := model4.Coeff(0), model4.Coeff(1)

	fmt.Printf("Intercept: %.4f\n", intercept)
	fmt.Printf("Slope: %.4f\n", slope)
}
