package main

import (
	"fmt"
	"os"
)

const resultFileName = "result.txt"

func writeToFile(result string) {
	resultText := fmt.Sprint(result)
	os.WriteFile(resultFileName, []byte(resultText), 0644)
}

func main() {
	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	resultText := fmt.Sprint("EBT: ", ebt, "\nProfit: ", profit, "\nRatio: ", ratio)
	writeToFile(resultText)

}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		panic("Value must be greater than 0.")
	}

	return userInput
}
