package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	totalMiles := 0
	totalGallons := 0

	for {
		fmt.Print("Enter trip miles: ")
		scanner.Scan()
		miles, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		fmt.Print("Enter gallons used: ")
		scanner.Scan()
		gallon, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		milesPerGallon := float64(miles) / float64(gallon)
		fmt.Printf("The trip mile per gallon is %.2f\n", milesPerGallon)

		totalMiles += miles
		totalGallons += gallon

		fmt.Println("Do you want to add another trip?: ")
		scanner.Scan()
		answer := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if answer == "yes" || answer == "y" || answer == "1" {
			continue
		} else {
			break
		}
	}

	averageMilesPerGallon := float64(totalMiles) / float64(totalGallons)
	fmt.Printf("The combined miles per gallon is: %.2f miles/gallon\n", averageMilesPerGallon)
}
