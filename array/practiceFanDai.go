package array

import (
	"fmt"
	"math"
	"sort"
)

type InterestRate int

const (
	BANK InterestRate = iota
	YOUNG1
	YOUNG2
	YOUNG3
	GOV
)

func ShiShuanFangDai(interestRateType InterestRate, totalYear int) {
	generateInterestRate(2000000, interestRateType, totalYear)
}

func generateInterestRate(totalAmount int, interestRateType InterestRate, totalYear int) {

	interestRates := make(map[int]float64, totalYear * 12)
	switch interestRateType {
	case 0:
		interestRates = getBankInterestRate(totalYear)
	}

	cal(totalAmount, interestRates)

}

func cal(totalAmount int, interestRates map[int]float64) {
	totalMonths := len(interestRates)
	orderedYears := getSequenceMonths(1, totalMonths)

	temp := float64(totalAmount)
	for _, year := range orderedYears {

		monthInterestRate := interestRates[year]
		avgApportion := (math.Pow(1.0 + monthInterestRate, float64(totalMonths)) * monthInterestRate) / (math.Pow(1.0 + monthInterestRate, float64(totalMonths))  - 1)
		fmt.Println("總月數", totalMonths, "年", year, "利率", monthInterestRate, "試算", avgApportion, "試算", temp * avgApportion)

		temp -= temp * avgApportion
		//fmt.Println("year", year, "apportion", avgApportion, "apportion", float64(totalAmount) * avgApportion)
	}

}

func getBankInterestRate(totalYear int) map[int]float64 {
	interestRates := make(map[int]float64, totalYear * 12)
	result := getSequenceMonths(1, totalYear * 12)

	for _, value := range result {
		interestRates[value] = getFixedBankRate()
	}

	return interestRates
}


func getSequenceMonths(min, max int) []int{
	fmt.Println("月數介於", min, max)

	yearsArray := make([]int, max - min + 1)
	for index, _ := range yearsArray  {
		yearsArray[index] = min + index
	}

	sort.Ints(yearsArray)

	return yearsArray
}

func getFixedBankRate() float64 {
	return 0.0140
}
