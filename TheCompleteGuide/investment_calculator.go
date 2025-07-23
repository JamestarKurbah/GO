package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount, expectReturn, year float64

	// add user input

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Expected return: ")
	fmt.Scan(&expectReturn)
	fmt.Print("Year: ")
	fmt.Scan(&year)

	// calculation the investment before inflation
	var futureValue = investmentAmount * math.Pow(1+expectReturn/100, year)
	// calculate final amount after inflation
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, year)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
