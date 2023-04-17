package main

import "fmt"

func main() {
	var sales []float64
	var item float64

	// Input the sales of the salesperson
	fmt.Println("Enter the sales of the salesperson:")
	fmt.Println("(Enter -1 to end the input)")

	for {
		fmt.Scanln(&item)
		if item == -1 {
			break
		}
		sales = append(sales, item)
	}

	totalSales := 0.0
	for _, sale := range sales {
		totalSales += sale
	}

	earnings := 200.0 + 0.09*totalSales

	fmt.Printf("The salesperson's earnings for last week: $%.2f\n", earnings)
}
