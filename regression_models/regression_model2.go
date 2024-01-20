package main

import (
	"fmt"

	"github.com/sajari/regression"
)

func main() {
	// x2/y2 data
	x := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y := []float64{9.14, 8.14, 8.74, 8.77, 9.26, 8.1, 6.13, 3.1, 9.13, 7.26, 4.74}

	// create linear regressio model
	model2 := new(regression.Regression)

	// add data to model
	for i, xi := range x {
		model2.Train(regression.DataPoint(y[i], []float64{xi}))
	}

	// train model
	model2.Run()

	// get coeefficients
	intercept, slope := model2.Coeff(0), model2.Coeff(1)

	fmt.Printf("Intercept: %.4f\n", intercept)
	fmt.Printf("Slope: %.4f\n", slope)
}
