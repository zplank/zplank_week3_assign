package main

import (
	"fmt"
)

type Row struct {
	X1 float64
	X2 float64
	X3 float64
	X4 float64
	Y1 float64
	Y2 float64
	Y3 float64
	Y4 float64
}

func main() {
	// Create a slice of structs representing the DataFrame
	anscombe := []Row{
		{10, 10, 10, 8, 8.04, 9.14, 7.46, 6.58},
		{8, 8, 8, 8, 6.95, 8.14, 6.77, 5.76},
		{13, 13, 13, 8, 7.58, 8.74, 12.74, 7.71},
		{9, 9, 9, 8, 8.81, 8.77, 7.11, 8.84},
		{11, 11, 11, 8, 8.33, 9.26, 7.81, 8.47},
		{14, 14, 14, 8, 9.96, 8.1, 8.84, 7.04},
		{6, 6, 6, 8, 7.24, 6.13, 6.08, 5.25},
		{4, 4, 4, 19, 4.26, 3.1, 5.39, 12.5},
		{12, 12, 12, 8, 10.84, 9.13, 8.15, 5.56},
		{7, 7, 7, 8, 4.82, 7.26, 6.42, 7.91},
		{5, 5, 5, 8, 5.68, 4.74, 5.73, 6.89},
	}

	// Print the DataFrame
	printDataFrame(anscombe)
}

func printDataFrame(data []Row) {
	// Print the DataFrame header
	fmt.Printf("%-5s%-5s%-5s%-5s%-7s%-7s%-7s%-7s\n", "X1", "X2", "X3", "X4", "Y1", "Y2", "Y3", "Y4")

	// Print the DataFrame rows
	for _, row := range data {
		fmt.Printf("%-5.2f%-5.2f%-5.2f%-5.2f%-7.2f%-7.2f%-7.2f%-7.2f\n", row.X1, row.X2, row.X3, row.X4, row.Y1, row.Y2, row.Y3, row.Y4)
	}
}
