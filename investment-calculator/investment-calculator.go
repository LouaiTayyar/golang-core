package main

import (
	"fmt" 
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount, years, expectedReturnRate float64

	fmt.Println("Investment Amoount:")
	fmt.Scan(&investmentAmount)
	fmt.Println("Years:")
	fmt.Scan(&years)
	fmt.Println("Expected Return Rate:")
	fmt.Scan(&expectedReturnRate)

	futureValue,futureRealValue := calculateFutureValues(investmentAmount,expectedReturnRate,inflationRate,years)
	formattedFV := fmt.Sprintf("Future Value: %.0f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value (adjusted for inflation): %.0f\n", futureRealValue)
	fmt.Print(formattedFV,formattedFRV)
	outputText("Function\n")
}


func outputText(text string){
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount,expectedReturnRate,inflationRate,years float64) (float64,float64) {
	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100,years)
	return futureValue,futureRealValue
}