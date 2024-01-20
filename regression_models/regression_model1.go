package main

import (
	"fmt"

	"github.com/sajari/regression"
)

func main() {
	// x1/y1 data
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	// create linear regressio model
	model1 := new(regression.Regression)

	// add data to model
	for i, xi := range x {
		model1.Train(regression.DataPoint(y[i], []float64{xi}))
	}

	// train model
	model1.Run()

	// get coeefficients
	intercept, slope := model1.Coeff(0), model1.Coeff(1)

	fmt.Printf("Intercept: %.4f\n", intercept)
	fmt.Printf("Slope: %.4f\n", slope)
}
