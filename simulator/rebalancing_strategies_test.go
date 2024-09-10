package simulator

import (
	"testing"

	"github.com/guilam34/financial_planner/models"
	"github.com/guilam34/financial_planner/test_utils"
)

type RebalanceToZeroTestCase struct {
	CaseName      string
	Year          int
	InitPortfolio models.Portfolio
	EndPortfolio  models.Portfolio
}

var rebalanceToZeroCases = []RebalanceToZeroTestCase{
	{
		CaseName: "One positive",
		InitPortfolio: models.Portfolio{
			models.Equities: 200_000,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 200_000,
		},
	},
	{
		CaseName: "One negative",
		InitPortfolio: models.Portfolio{
			models.Equities: -200_000,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: -200_000,
		},
	},
	{
		CaseName: "One positive, one negative",
		InitPortfolio: models.Portfolio{
			models.Equities: 200_000,
			models.Bonds:    -10_000,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 190_000,
			models.Bonds:    0,
		},
	},
	{
		CaseName: "Two positive, one negative",
		InitPortfolio: models.Portfolio{
			models.Equities: 200_000,
			models.Cash:     50_000,
			models.Bonds:    -10_000,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 195_000,
			models.Cash:     45_000,
			models.Bonds:    0,
		},
	},
}

func TestRebalancetoZeroCases(t *testing.T) {
	for _, test := range rebalanceToZeroCases {
		t.Run(test.CaseName, func(t *testing.T) {
			actualPortfolio := RebalanceToZero{}.Rebalance(
				test.InitPortfolio,
				models.PortfolioAllocation{},
				test.Year)
			for assetType, expectedVal := range test.EndPortfolio {
				actualVal, ok := actualPortfolio[assetType]
				if !ok || !test_utils.AlmostEqual(actualVal, expectedVal) {
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
	PortfolioAllocation models.PortfolioAllocation
	InitPortfolio       models.Portfolio
	EndPortfolio        models.Portfolio
}

var rebalanceEveryNYearsCases = []RebalanceEveryNYearsTestCase{
	{
		CaseName:           "One negative",
		Year:               1,
		RebalancingCadence: 2,
		PortfolioAllocation: models.PortfolioAllocation{
			models.Equities: models.AssetAllocation{Allocation: 1.0},
		},
		InitPortfolio: models.Portfolio{
			models.Equities: -200_000,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: -200_000,
		},
	},
	{
		CaseName:           "Not rebalancing year w/ all assets positive",
		Year:               1,
		RebalancingCadence: 2,
		PortfolioAllocation: models.PortfolioAllocation{
			models.Equities: models.AssetAllocation{Allocation: 1.0},
		},
		InitPortfolio: models.Portfolio{
			models.Equities: 200_000,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 200_000,
		},
	},
	{
		CaseName:           "Not rebalancing year w/ some assets negative",
		Year:               1,
		RebalancingCadence: 2,
		PortfolioAllocation: models.PortfolioAllocation{
			models.Equities: models.AssetAllocation{Allocation: 0.7},
			models.Bonds:    models.AssetAllocation{Allocation: 0.3},
		},
		InitPortfolio: models.Portfolio{
			models.Equities: 200_000,
			models.Bonds:    -10_000,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 190_000,
			models.Bonds:    0,
		},
	},
	{
		CaseName:           "Rebalancing year",
		Year:               2,
		RebalancingCadence: 2,
		PortfolioAllocation: models.PortfolioAllocation{
			models.Equities: models.AssetAllocation{Allocation: 0.7},
			models.Bonds:    models.AssetAllocation{Allocation: 0.3},
		},
		InitPortfolio: models.Portfolio{
			models.Equities: 200_000,
			models.Bonds:    -10_000,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 133_000,
			models.Bonds:    57_000,
		},
	},
}

func TestRebalanceEveryNYearsCases(t *testing.T) {
	for _, test := range rebalanceEveryNYearsCases {
		t.Run(test.CaseName, func(t *testing.T) {
			actualPortfolio := RebalanceEveryNYears{rebalanceCadence: test.RebalancingCadence}.Rebalance(
				test.InitPortfolio,
				test.PortfolioAllocation,
				test.Year)
			for assetType, expectedVal := range test.EndPortfolio {
				actualVal, ok := actualPortfolio[assetType]
				if !ok || !test_utils.AlmostEqual(actualVal, expectedVal) {
					t.Errorf("Expected %v but got %v", test.EndPortfolio, actualPortfolio)
					t.FailNow()
				}
			}
		})
	}
}
