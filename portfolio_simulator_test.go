package main

import (
	"math"
	"strconv"
	"testing"
)

type TestCase struct {
	CaseName                    string
	LastYear                    int
	ExpectedTaxRateInRetirement float64
	AnnualInflationRate         float64
	AnnualContributions         []AnnualContribution
	PortfolioAllocation         PortfolioAllocation
	InitPortfolio               Portfolio
	EndPortfolio                Portfolio
	ErrorMessage                string
}

const float64EqualityThreshold = 1

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) <= float64EqualityThreshold
}

func fromScientificNotation(a string) float64 {
	val, _ := strconv.ParseFloat(a, 64)
	return val
}

var successCases = []TestCase{
	{
		CaseName:            "LastYearIsToday",
		LastYear:            0,
		AnnualInflationRate: 0.0,
		AnnualContributions: []AnnualContribution{},
		PortfolioAllocation: PortfolioAllocation{
			Equities: {
				ReturnRate: 0.07,
				Allocation: 1.0,
			},
		},
		InitPortfolio: Portfolio{
			Equities: 200_000,
		},
		EndPortfolio: Portfolio{
			Equities: 200_000,
		},
	},
	{
		CaseName:            "LastYearIsOneYearAway_SingleAsset_NoContributions_NoInflation_WithReturns",
		LastYear:            1,
		AnnualInflationRate: 0.0,
		AnnualContributions: []AnnualContribution{},
		PortfolioAllocation: PortfolioAllocation{
			Equities: {
				ReturnRate: 0.07,
				Allocation: 1.0,
			},
		},
		InitPortfolio: Portfolio{
			Equities: 200_000,
		},
		EndPortfolio: Portfolio{
			Equities: 214_000,
		},
	},
	{
		CaseName:            "LastYearIsOneYearAway_SingleAsset_NoContributions_WithInflation_NoReturns",
		LastYear:            1,
		AnnualInflationRate: 0.05,
		AnnualContributions: []AnnualContribution{},
		PortfolioAllocation: PortfolioAllocation{
			Equities: {
				ReturnRate: 0.0,
				Allocation: 1.0,
			},
		},
		InitPortfolio: Portfolio{
			Equities: 200_000,
		},
		EndPortfolio: Portfolio{
			Equities: 190_000,
		},
	},
	{
		CaseName:            "LastYearIsOneYearAway_SingleAsset_WithConstantContributions_NoInflation_WithReturns",
		LastYear:            1,
		AnnualInflationRate: 0.0,
		AnnualContributions: []AnnualContribution{
			{
				Amount:          10_000,
				StartYear:       0,
				EndYear:         1,
				AnnualPctChange: 0.0,
			},
			{
				Amount:          40_000,
				StartYear:       0,
				EndYear:         1,
				AnnualPctChange: 0.0,
			},
		},
		PortfolioAllocation: PortfolioAllocation{
			Equities: {
				ReturnRate: 0.07,
				Allocation: 1.0,
			},
		},
		InitPortfolio: Portfolio{
			Equities: 200_000,
		},
		EndPortfolio: Portfolio{
			Equities: 264_000,
		},
	},
	{
		CaseName:            "LastYearIsOneYearAway_SingleAsset_WithConstantContributions_WithInflation_WithReturns",
		LastYear:            1,
		AnnualInflationRate: 0.07,
		AnnualContributions: []AnnualContribution{
			{
				Amount:          10_000,
				StartYear:       0,
				EndYear:         1,
				AnnualPctChange: 0.0,
			},
			{
				Amount:          40_000,
				StartYear:       0,
				EndYear:         1,
				AnnualPctChange: 0.0,
			},
		},
		PortfolioAllocation: PortfolioAllocation{
			Equities: {
				ReturnRate: 0.07,
				Allocation: 1.0,
			},
		},
		InitPortfolio: Portfolio{
			Equities: 200_000,
		},
		EndPortfolio: Portfolio{
			Equities: 250_000,
		},
	},
	{
		CaseName:            "LastYearIsOneYearAway_MultipleAssets_WithConstantContributions_NoInflation_NoReturns",
		LastYear:            1,
		AnnualInflationRate: 0.0,
		AnnualContributions: []AnnualContribution{
			{
				Amount:          10_000,
				StartYear:       0,
				EndYear:         1,
				AnnualPctChange: 0.0,
			},
			{
				Amount:          40_000,
				StartYear:       0,
				EndYear:         1,
				AnnualPctChange: 0.0,
			},
		},
		PortfolioAllocation: PortfolioAllocation{
			Equities: {
				ReturnRate: 0.0,
				Allocation: 0.9,
			},
			Cash: {
				ReturnRate: 0.0,
				Allocation: 0.1,
			},
		},
		InitPortfolio: Portfolio{
			Equities: 200_000,
		},
		EndPortfolio: Portfolio{
			Equities: 245_000,
			Cash:     5_000,
		},
	},
	{
		CaseName:            "LastYearIsOneYearAway_MultipleAssets_WithConstantContributions_WithInflation_WithReturns",
		LastYear:            1,
		AnnualInflationRate: 0.05,
		AnnualContributions: []AnnualContribution{
			{
				Amount:          10_000,
				StartYear:       0,
				EndYear:         1,
				AnnualPctChange: 0.0,
			},
			{
				Amount:          40_000,
				StartYear:       0,
				EndYear:         1,
				AnnualPctChange: 0.0,
			},
		},
		PortfolioAllocation: PortfolioAllocation{
			Equities: {
				ReturnRate: 0.07,
				Allocation: 0.9,
			},
			Cash: {
				ReturnRate: 0.05,
				Allocation: 0.1,
			},
		},
		InitPortfolio: Portfolio{
			Equities: 200_000,
		},
		EndPortfolio: Portfolio{
			Equities: 249_000,
			Cash:     5_000,
		},
	},
	{
		CaseName:            "LastYearIsMultipleYearsAway_MultipleAssets_WithConstantContributions_WithInflation_WithReturns",
		LastYear:            5,
		AnnualInflationRate: 0.05,
		AnnualContributions: []AnnualContribution{
			{
				Amount:          10_000,
				StartYear:       0,
				EndYear:         5,
				AnnualPctChange: 0.0,
			},
			{
				Amount:          40_000,
				StartYear:       0,
				EndYear:         5,
				AnnualPctChange: 0.0,
			},
		},
		PortfolioAllocation: PortfolioAllocation{
			Equities: {
				ReturnRate: 0.07,
				Allocation: 0.9,
			},
			Cash: {
				ReturnRate: 0.05,
				Allocation: 0.1,
			},
		},
		InitPortfolio: Portfolio{
			Equities: 200_000,
		},
		EndPortfolio: Portfolio{
			Equities: 454_997,
			Cash:     25_000,
		},
	},
	{
		CaseName:            "LastYearIsMultipleYearsAway_MultipleAssets_WithDecliningContributions_WithInflation_WithReturns",
		LastYear:            5,
		AnnualInflationRate: 0.05,
		AnnualContributions: []AnnualContribution{
			{
				Amount:          10_000,
				StartYear:       0,
				EndYear:         5,
				AnnualPctChange: -0.1,
			},
			{
				Amount:          40_000,
				StartYear:       0,
				EndYear:         5,
				AnnualPctChange: -0.1,
			},
		},
		PortfolioAllocation: PortfolioAllocation{
			Equities: {
				ReturnRate: 0.07,
				Allocation: 0.9,
			},
			Cash: {
				ReturnRate: 0.05,
				Allocation: 0.1,
			},
		},
		InitPortfolio: Portfolio{
			Equities: 200_000,
		},
		EndPortfolio: Portfolio{
			Equities: 413_412,
			Cash:     20_475,
		},
	},
	{
		CaseName:            "LastYearIsMultipleYearsAway_MultipleAssets_WithIncreasingContributions_WithInflation_WithReturns",
		LastYear:            5,
		AnnualInflationRate: 0.05,
		AnnualContributions: []AnnualContribution{
			{
				Amount:          10_000,
				StartYear:       0,
				EndYear:         5,
				AnnualPctChange: 0.1,
			},
			{
				Amount:          40_000,
				StartYear:       0,
				EndYear:         5,
				AnnualPctChange: 0.1,
			},
		},
		PortfolioAllocation: PortfolioAllocation{
			Equities: {
				ReturnRate: 0.07,
				Allocation: 0.9,
			},
			Cash: {
				ReturnRate: 0.05,
				Allocation: 0.1,
			},
		},
		InitPortfolio: Portfolio{
			Equities: 200_000,
		},
		EndPortfolio: Portfolio{
			Equities: 505_682,
			Cash:     30_525,
		},
	},
	{
		CaseName:            "LastYearIsMultipleYearsAway_MultipleAssets_WithConstantContributionsStoppingPartway_WithInflation_WithReturns",
		LastYear:            5,
		AnnualInflationRate: 0.05,
		AnnualContributions: []AnnualContribution{
			{
				Amount:          10_000,
				StartYear:       0,
				EndYear:         2,
				AnnualPctChange: 0.0,
			},
			{
				Amount:          40_000,
				StartYear:       0,
				EndYear:         2,
				AnnualPctChange: 0.0,
			},
		},
		PortfolioAllocation: PortfolioAllocation{
			Equities: {
				ReturnRate: 0.07,
				Allocation: 0.9,
			},
			Cash: {
				ReturnRate: 0.05,
				Allocation: 0.1,
			},
		},
		InitPortfolio: Portfolio{
			Equities: 200_000,
		},
		EndPortfolio: Portfolio{
			Equities: 317_279,
			Cash:     10_000,
		},
	},
	{
		CaseName:            "LastYearIsMultipleYearsAway_MultipleAssets_WithDecliningConstantContributionsStoppingPartway_WithInflation_WithReturns",
		LastYear:            5,
		AnnualInflationRate: 0.05,
		AnnualContributions: []AnnualContribution{
			{
				Amount:          10_000,
				StartYear:       0,
				EndYear:         2,
				AnnualPctChange: -0.1,
			},
			{
				Amount:          40_000,
				StartYear:       0,
				EndYear:         2,
				AnnualPctChange: -0.1,
			},
		},
		PortfolioAllocation: PortfolioAllocation{
			Equities: {
				ReturnRate: 0.07,
				Allocation: 0.9,
			},
			Cash: {
				ReturnRate: 0.05,
				Allocation: 0.1,
			},
		},
		InitPortfolio: Portfolio{
			Equities: 200_000,
		},
		EndPortfolio: Portfolio{
			Equities: 312_504,
			Cash:     9500,
		},
	},
	{
		CaseName:                    "LongTimeFrameWithEverythingSet",
		LastYear:                    60,
		ExpectedTaxRateInRetirement: 0.24,
		AnnualInflationRate:         0.02,
		AnnualContributions: []AnnualContribution{
			{
				Amount:          10_000,
				StartYear:       0,
				EndYear:         60,
				AnnualPctChange: 0.0,
			},
			{
				Amount:          40_000,
				StartYear:       0,
				EndYear:         60,
				AnnualPctChange: 0.0,
			},
		},
		PortfolioAllocation: PortfolioAllocation{
			Equities: {
				ReturnRate: 0.07,
				Allocation: 0.89,
			},
			Bonds: {
				ReturnRate: 0.04,
				Allocation: 0.08,
			},
			Cash: {
				ReturnRate: 0.03,
				Allocation: 0.03,
			},
		},
		InitPortfolio: Portfolio{
			Equities: 200_000,
		},
		EndPortfolio: Portfolio{
			Equities: 19_470_312,
			Bonds:    456_206,
			Cash:     122_504,
		},
	},
}

func TestForecastFuturePortfolioValueByYearSuccessCases(t *testing.T) {
	for _, test := range successCases {
		t.Run(test.CaseName, func(t *testing.T) {
			forecastedPortfoliosByYear, _ := ForecastFuturePortfolioValueByYear(test.InitPortfolio, test.AnnualContributions, test.PortfolioAllocation, test.AnnualInflationRate, test.LastYear)
			forecastedPortfolioForLastYear := forecastedPortfoliosByYear[len(forecastedPortfoliosByYear)-1]
			for assetType, expectedVal := range test.EndPortfolio {
				actualVal, ok := forecastedPortfolioForLastYear[assetType]
				if !ok || !almostEqual(actualVal, expectedVal) {
					t.Errorf("Expected %v but got %v", test.EndPortfolio, forecastedPortfolioForLastYear)
					t.FailNow()
				}
			}
		})
	}
}

var errorCases = []TestCase{
	{
		CaseName:            "AllocationsDoNotSumUpToOne",
		LastYear:            1,
		AnnualInflationRate: 0.0,
		AnnualContributions: []AnnualContribution{
			{
				Amount:          10_000,
				StartYear:       0,
				EndYear:         1,
				AnnualPctChange: 0.0,
			},
			{
				Amount:          40_000,
				StartYear:       0,
				EndYear:         1,
				AnnualPctChange: 0.0,
			},
		},
		PortfolioAllocation: PortfolioAllocation{
			Equities: {
				ReturnRate: 0.0,
				Allocation: 0.9,
			},
			Cash: {
				ReturnRate: 0.0,
				Allocation: 0.15,
			},
		},
		InitPortfolio: Portfolio{},
		ErrorMessage:  "Portfolio allocation percent must sum up to 1",
	},
	{
		CaseName:            "AnnualContributionStopsAfterLastYear",
		LastYear:            1,
		AnnualInflationRate: 0.0,
		AnnualContributions: []AnnualContribution{
			{
				Amount:          10_000,
				StartYear:       0,
				EndYear:         2,
				AnnualPctChange: 0.0,
			},
			{
				Amount:          40_000,
				StartYear:       0,
				EndYear:         2,
				AnnualPctChange: 0.0,
			},
		},
		InitPortfolio: Portfolio{},
		ErrorMessage:  "Annual contribution end year must be less than or equal to last year",
	},
}

func TestForecastFuturePortfolioValueByYearErrorCases(t *testing.T) {
	for _, test := range errorCases {
		t.Run(test.CaseName, func(t *testing.T) {
			_, err := ForecastFuturePortfolioValueByYear(test.InitPortfolio, test.AnnualContributions, test.PortfolioAllocation, test.AnnualInflationRate, test.LastYear)
			if err.Error() != test.ErrorMessage {
				t.Errorf("Expected %v but got %v", test.ErrorMessage, err.Error())
			}
		})
	}
}
