package correlations

import (
	"testing"
)

func TestCalculateCorrelation(t *testing.T) {
	// use Anscombe's quartet data to test correlation function
	x1 := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}
	y1 := []float64{8.04, 6.95, 7.58, 8.81, 8.33, 9.96, 7.24, 4.26, 10.84, 4.82, 5.68}

	//
	corr := CalculateCorrelation(x1, y1)

	// Expected correlation (based on correlation calculations in correlation_value files)
	expectedCorr := 0.81642051634484

	// Check if the calculated correlation is close to the expected value
	if corr != expectedCorr {
		t.Errorf("Correlation calculation failed. Got: %v, Expected: %v", corr, expectedCorr)
	}
}
