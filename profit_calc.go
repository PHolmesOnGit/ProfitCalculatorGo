package main

import (
	"fmt"
)

func main() {
	revenue, taxRate, expenses := getInput()

	earningsBeforeTax, incomeTax, profit, ratio := calculate(revenue, taxRate, expenses)

	printResults(earningsBeforeTax, incomeTax, profit, ratio)

}

func getInput() (float64, float64, float64) {
	// Gets user input to use for calculate()
	var revenue float64
	var taxRate float64
	var expenses float64

	fmt.Println("What is your gross income? : ")
	fmt.Scan(&revenue)

	fmt.Println("What is your estimated expenses amount? : ")
	fmt.Scan(&expenses)

	fmt.Println("What is the income tax rate? : ")
	fmt.Scan(&taxRate)

	return revenue, taxRate, expenses
}

func calculate(revenue float64, taxRate float64, expenses float64) (float64, float64, float64, float64) {
	// Calculates the users data from getInput()
	earningsBeforeTax := revenue - expenses
	incomeTax := (revenue / 100) * taxRate
	profit := revenue - expenses - incomeTax
	ratio := earningsBeforeTax / profit
	return earningsBeforeTax, incomeTax, profit, ratio
}

func printResults(earningsBeforeTax float64, incomeTax float64, profit float64, ratio float64) {
	fmt.Print("Earnings before tax = ")
	fmt.Println(earningsBeforeTax)

	fmt.Printf("Taxed income = %.2f\n", incomeTax)

	fmt.Printf("Profit = %v\n", profit)

	fmt.Printf("Ratio = %.2f\n", ratio)
}
