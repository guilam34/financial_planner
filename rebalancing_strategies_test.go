package main

import (
	"testing"
)

type RebalanceToZeroTestCase struct {
	CaseName      string
	Year          int
	InitPortfolio Portfolio
	EndPortfolio  Portfolio
}

var rebalanceToZeroCases = []RebalanceToZeroTestCase{
	{
		CaseName: "One positive",
		InitPortfolio: Portfolio{
			Equities: 200_000,
		},
		EndPortfolio: Portfolio{
			Equities: 200_000,
		},
	},
	{
		CaseName: "One negative",
		InitPortfolio: Portfolio{
			Equities: -200_000,
		},
		EndPortfolio: Portfolio{
			Equities: -200_000,
		},
	},
	{
		CaseName: "One positive, one negative",
		InitPortfolio: Portfolio{
			Equities: 200_000,
			Bonds:    -10_000,
		},
		EndPortfolio: Portfolio{
			Equities: 190_000,
			Bonds:    0,
		},
	},
	{
		CaseName: "Two positive, one negative",
		InitPortfolio: Portfolio{
			Equities: 200_000,
			Cash:     50_000,
			Bonds:    -10_000,
		},
		EndPortfolio: Portfolio{
			Equities: 195_000,
			Cash:     45_000,
			Bonds:    0,
		},
	},
}

func TestRebalancetoZeroCases(t *testing.T) {
	for _, test := range rebalanceToZeroCases {
		t.Run(test.CaseName, func(t *testing.T) {
			actualPortfolio := RebalanceToZero{}.Rebalance(
				test.InitPortfolio,
				PortfolioAllocation{},
				test.Year)
			for assetType, expectedVal := range test.EndPortfolio {
				actualVal, ok := actualPortfolio[assetType]
				if !ok || !almostEqual(actualVal, expectedVal) {
					t.Errorf("Expected %v but got %v", test.EndPortfolio, actualPortfolio)
					t.FailNow()
				}
			}
		})
	}
}

type RebalanceEveryNYearsTestCase struct {
	CaseName            string
	Year                int
	RebalancingCadence  int
	PortfolioAllocation PortfolioAllocation
	InitPortfolio       Portfolio
	EndPortfolio        Portfolio
}

var rebalanceEveryNYearsCases = []RebalanceEveryNYearsTestCase{
	{
		CaseName:           "One negative",
		Year:               1,
		RebalancingCadence: 2,
		PortfolioAllocation: PortfolioAllocation{
			Equities: assetAllocation{Allocation: 1.0},
		},
		InitPortfolio: Portfolio{
			Equities: -200_000,
		},
		EndPortfolio: Portfolio{
			Equities: -200_000,
		},
	},
	{
		CaseName:           "Not rebalancing year w/ all assets positive",
		Year:               1,
		RebalancingCadence: 2,
		PortfolioAllocation: PortfolioAllocation{
			Equities: assetAllocation{Allocation: 1.0},
		},
		InitPortfolio: Portfolio{
			Equities: 200_000,
		},
		EndPortfolio: Portfolio{
			Equities: 200_000,
		},
	},
	{
		CaseName:           "Not rebalancing year w/ some assets negative",
		Year:               1,
		RebalancingCadence: 2,
		PortfolioAllocation: PortfolioAllocation{
			Equities: assetAllocation{Allocation: 0.7},
			Bonds:    assetAllocation{Allocation: 0.3},
		},
		InitPortfolio: Portfolio{
			Equities: 200_000,
			Bonds:    -10_000,
		},
		EndPortfolio: Portfolio{
			Equities: 190_000,
			Bonds:    0,
		},
	},
	{
		CaseName:           "Rebalancing year",
		Year:               2,
		RebalancingCadence: 2,
		PortfolioAllocation: PortfolioAllocation{
			Equities: assetAllocation{Allocation: 0.7},
			Bonds:    assetAllocation{Allocation: 0.3},
		},
		InitPortfolio: Portfolio{
			Equities: 200_000,
			Bonds:    -10_000,
		},
		EndPortfolio: Portfolio{
			Equities: 133_000,
			Bonds:    57_000,
		},
	},
}

func TestRebalanceEveryNYearsCases(t *testing.T) {
	for _, test := range rebalanceEveryNYearsCases {
		t.Run(test.CaseName, func(t *testing.T) {
			actualPortfolio := RebalanceEveryNYears{n: 2}.Rebalance(
				test.InitPortfolio,
				test.PortfolioAllocation,
				test.Year)
			for assetType, expectedVal := range test.EndPortfolio {
				actualVal, ok := actualPortfolio[assetType]
				if !ok || !almostEqual(actualVal, expectedVal) {
					t.Errorf("Expected %v but got %v", test.EndPortfolio, actualPortfolio)
					t.FailNow()
				}
			}
		})
	}
}
