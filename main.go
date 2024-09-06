package main

import (
	"fmt"
	"math"
)

type AnnualIncomeOrWithdrawal struct {
	Amount          float64
	StartYear       int
	EndYear         int
	AnnualPctChange float64
}

func main() {
	retirementYear := 30
	lastYear := 60
	expectedTaxRateInRetirement := 0.24
	annualIncomeOrWithdrawals := []AnnualIncomeOrWithdrawal{
		{
			Amount:          10_000,
			StartYear:       0,
			EndYear:         retirementYear,
			AnnualPctChange: 0.01,
		},
		{
			Amount:          40_000,
			StartYear:       0,
			EndYear:         retirementYear,
			AnnualPctChange: 0.02,
		},
		{
			Amount:          -150_000,
			StartYear:       retirementYear,
			EndYear:         lastYear,
			AnnualPctChange: 0.02,
		},
	}

	fmt.Println(forecastFuturePortfolioValueByYear(200_000, annualIncomeOrWithdrawals, 0.07, 0.02, lastYear, expectedTaxRateInRetirement))
}

func forecastFuturePortfolioValueByYear(initVal float64, annualIncomeOrWithdrawals []AnnualIncomeOrWithdrawal, nominalAnnualReturnRate float64, annualInflationRate float64, years int, expectedTaxRateInRetirement float64) []float64 {
	var result []float64
	realAnnualReturnRate := convertToRealRate(nominalAnnualReturnRate, annualInflationRate)
	for year := 0; year <= years; year++ {
		result = append(result, forecastFuturePortfolioValue(initVal, annualIncomeOrWithdrawals, realAnnualReturnRate, year, expectedTaxRateInRetirement))
	}
	return result
}

func convertToRealRate(nominalRate float64, inflationRate float64) float64 {
	return ((1 + nominalRate) / (1 + inflationRate)) - 1
}

func forecastFuturePortfolioValue(initVal float64, annualIncomeOrWithdrawals []AnnualIncomeOrWithdrawal, annualReturnRate float64, year int, expectedTaxRateInRetirement float64) float64 {
	result := initVal * math.Pow(1+annualReturnRate, float64(year))
	for curYear := 0; curYear <= year; curYear++ {
		for _, contrib := range annualIncomeOrWithdrawals {
			if curYear >= contrib.StartYear && curYear <= contrib.EndYear {
				contribAmtAdjustedForChangePct := contrib.Amount
				// Adjust for change in contribution after the first year
				if curYear >= contrib.StartYear {
					contribAmtAdjustedForChangePct = contrib.Amount * math.Pow(1+contrib.AnnualPctChange, float64(curYear-contrib.StartYear-1))
				}
				// Adjust for taxes if this is a withdrawal
				if contrib.Amount <= 0 {
					contribAmtAdjustedForChangePct = contribAmtAdjustedForChangePct / (1 - expectedTaxRateInRetirement)
				}
				result = result + float64(contribAmtAdjustedForChangePct)*math.Pow(1+annualReturnRate, float64(year-curYear))
			}
		}
	}
	return result
}
