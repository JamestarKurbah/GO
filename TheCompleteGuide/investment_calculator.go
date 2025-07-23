package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investAmt float64 = 1000
	var expectRurn float64 = 4.5
	var year float64 = 10

	var futureValue = investAmt * math.Pow(1+expectRurn/100, year)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100,year)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
