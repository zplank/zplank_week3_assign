package correlations

import (
	"gonum.org/v1/gonum/stat"
)

// calcualte correlation between an x and y value
func CalculateCorrelation(x, y []float64) float64 {
	return stat.Correlation(x, y, nil)
}
