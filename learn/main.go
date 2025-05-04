package main

import (
	"fmt"
)

func waging(x, y float64) float64 {
	return x * y
}

func report() {
	for {
		var i string
		fmt.Print("\nEnter \n1 for Wage\n2 for Tax \nor # to exit: ")
		fmt.Scanln(&i)

		if i == "#" {
			fmt.Println("Exiting report...")
			break
		} else if i == "1" {
			fmt.Println("Wage in Ontario is $17.20")
		} else if i == "2" {
			fmt.Println("Tax in Ontario is 13%")
		} else {
			fmt.Println("Invalid input. Please enter 1, 2, or #.")
		}
	}
}

func main() {
	var x float64
	const y float64 = 17.20

	fmt.Println("--------Ontario Wage Calculator--------")
	fmt.Println("")
	fmt.Print("Please enter the number of hours worked: ")
	fmt.Scanln(&x)

	calculation := waging(x, y)
	fmt.Println("")
	fmt.Printf("Your bi-weekly pay is: $%.2f\n", calculation)

	fmt.Println("\n---------REPORT---------")
	report()
}
