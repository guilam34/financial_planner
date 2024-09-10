package simulator

import (
	"testing"

	"github.com/guilam34/financial_planner/models"
	"github.com/guilam34/financial_planner/test_utils"
)

type PortfolioSimulatorTestCase struct {
	CaseName        string
	ForecastRequest models.ForecastPortfolioRequest
	EndPortfolio    models.Portfolio
	ErrorMessage    string
}

var simulationSuccessCases = []PortfolioSimulatorTestCase{
	{
		CaseName: "EndYearIsToday",
		ForecastRequest: models.ForecastPortfolioRequest{
			EndYear:                       0,
			AnnualInflationRate:           0.0,
			AnnualPortfolioBalanceChanges: []models.AnnualPortfolioBalanceChange{},
			PortfolioAllocation: models.PortfolioAllocation{
				models.Equities: {
					ReturnRate: 0.07,
					Allocation: 1.0,
				},
			},
			InitPortfolio: models.Portfolio{
				models.Equities: 200_000,
			},
			RebalanceCadence:    1,
			RebalancingStrategy: models.YearlyToZero,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 200_000,
		},
	},
	{
		CaseName: "EndYearIsOneYearAway_SingleAsset_NoContributions_NoInflation_WithReturns",
		ForecastRequest: models.ForecastPortfolioRequest{
			EndYear:                       1,
			AnnualInflationRate:           0.0,
			AnnualPortfolioBalanceChanges: []models.AnnualPortfolioBalanceChange{},
			PortfolioAllocation: models.PortfolioAllocation{
				models.Equities: {
					ReturnRate: 0.07,
					Allocation: 1.0,
				},
			},
			InitPortfolio: models.Portfolio{
				models.Equities: 200_000,
			},
			RebalanceCadence:    1,
			RebalancingStrategy: models.YearlyToZero,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 214_000,
		},
	},
	{
		CaseName: "EndYearIsOneYearAway_SingleAsset_NoContributions_WithInflation_NoReturns",
		ForecastRequest: models.ForecastPortfolioRequest{
			EndYear:                       1,
			AnnualInflationRate:           0.05,
			AnnualPortfolioBalanceChanges: []models.AnnualPortfolioBalanceChange{},
			PortfolioAllocation: models.PortfolioAllocation{
				models.Equities: {
					ReturnRate: 0.0,
					Allocation: 1.0,
				},
			},
			InitPortfolio: models.Portfolio{
				models.Equities: 200_000,
			},
			RebalanceCadence:    1,
			RebalancingStrategy: models.YearlyToZero,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 190_000,
		},
	},
	{
		CaseName: "EndYearIsOneYearAway_SingleAsset_WithConstantContributions_NoInflation_WithReturns",
		ForecastRequest: models.ForecastPortfolioRequest{
			EndYear:             1,
			AnnualInflationRate: 0.0,
			AnnualPortfolioBalanceChanges: []models.AnnualPortfolioBalanceChange{
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
			PortfolioAllocation: models.PortfolioAllocation{
				models.Equities: {
					ReturnRate: 0.07,
					Allocation: 1.0,
				},
			},
			InitPortfolio: models.Portfolio{
				models.Equities: 200_000,
			},
			RebalanceCadence:    1,
			RebalancingStrategy: models.YearlyToZero,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 264_000,
		},
	},
	{
		CaseName: "EndYearIsOneYearAway_SingleAsset_WithConstantContributions_WithInflation_WithReturns",
		ForecastRequest: models.ForecastPortfolioRequest{
			EndYear:             1,
			AnnualInflationRate: 0.07,
			AnnualPortfolioBalanceChanges: []models.AnnualPortfolioBalanceChange{
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
			PortfolioAllocation: models.PortfolioAllocation{
				models.Equities: {
					ReturnRate: 0.07,
					Allocation: 1.0,
				},
			},
			InitPortfolio: models.Portfolio{
				models.Equities: 200_000,
			},
			RebalanceCadence:    1,
			RebalancingStrategy: models.YearlyToZero,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 250_000,
		},
	},
	{
		CaseName: "EndYearIsOneYearAway_MultipleAssets_WithConstantContributions_NoInflation_NoReturns",
		ForecastRequest: models.ForecastPortfolioRequest{
			EndYear:             1,
			AnnualInflationRate: 0.0,
			AnnualPortfolioBalanceChanges: []models.AnnualPortfolioBalanceChange{
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
			PortfolioAllocation: models.PortfolioAllocation{
				models.Equities: {
					ReturnRate: 0.0,
					Allocation: 0.9,
				},
				models.Cash: {
					ReturnRate: 0.0,
					Allocation: 0.1,
				},
			},
			InitPortfolio: models.Portfolio{
				models.Equities: 200_000,
			},
			RebalanceCadence:    1,
			RebalancingStrategy: models.YearlyToZero,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 245_000,
			models.Cash:     5_000,
		},
	},
	{
		CaseName: "EndYearIsOneYearAway_MultipleAssets_WithConstantContributions_WithInflation_WithReturns",
		ForecastRequest: models.ForecastPortfolioRequest{
			EndYear:             1,
			AnnualInflationRate: 0.05,
			AnnualPortfolioBalanceChanges: []models.AnnualPortfolioBalanceChange{
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
			PortfolioAllocation: models.PortfolioAllocation{
				models.Equities: {
					ReturnRate: 0.07,
					Allocation: 0.9,
				},
				models.Cash: {
					ReturnRate: 0.05,
					Allocation: 0.1,
				},
			},
			InitPortfolio: models.Portfolio{
				models.Equities: 200_000,
			},
			RebalanceCadence:    1,
			RebalancingStrategy: models.YearlyToZero,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 249_000,
			models.Cash:     5_000,
		},
	},
	{
		CaseName: "EndYearIsMultipleYearsAway_MultipleAssets_WithConstantContributions_WithInflation_WithReturns",
		ForecastRequest: models.ForecastPortfolioRequest{
			EndYear:             5,
			AnnualInflationRate: 0.05,
			AnnualPortfolioBalanceChanges: []models.AnnualPortfolioBalanceChange{
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
			PortfolioAllocation: models.PortfolioAllocation{
				models.Equities: {
					ReturnRate: 0.07,
					Allocation: 0.9,
				},
				models.Cash: {
					ReturnRate: 0.05,
					Allocation: 0.1,
				},
			},
			InitPortfolio: models.Portfolio{
				models.Equities: 200_000,
			},
			RebalanceCadence:    1,
			RebalancingStrategy: models.YearlyToZero,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 454_997,
			models.Cash:     25_000,
		},
	},
	{
		CaseName: "EndYearIsMultipleYearsAway_MultipleAssets_WithDecliningContributions_WithInflation_WithReturns",
		ForecastRequest: models.ForecastPortfolioRequest{
			EndYear:             5,
			AnnualInflationRate: 0.05,
			AnnualPortfolioBalanceChanges: []models.AnnualPortfolioBalanceChange{
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
			PortfolioAllocation: models.PortfolioAllocation{
				models.Equities: {
					ReturnRate: 0.07,
					Allocation: 0.9,
				},
				models.Cash: {
					ReturnRate: 0.05,
					Allocation: 0.1,
				},
			},
			InitPortfolio: models.Portfolio{
				models.Equities: 200_000,
			},
			RebalanceCadence:    1,
			RebalancingStrategy: models.YearlyToZero,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 413_412,
			models.Cash:     20_475,
		},
	},
	{
		CaseName: "EndYearIsMultipleYearsAway_MultipleAssets_WithIncreasingContributions_WithInflation_WithReturns",
		ForecastRequest: models.ForecastPortfolioRequest{
			EndYear:             5,
			AnnualInflationRate: 0.05,
			AnnualPortfolioBalanceChanges: []models.AnnualPortfolioBalanceChange{
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
			PortfolioAllocation: models.PortfolioAllocation{
				models.Equities: {
					ReturnRate: 0.07,
					Allocation: 0.9,
				},
				models.Cash: {
					ReturnRate: 0.05,
					Allocation: 0.1,
				},
			},
			InitPortfolio: models.Portfolio{
				models.Equities: 200_000,
			},
			RebalanceCadence:    1,
			RebalancingStrategy: models.YearlyToZero,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 505_682,
			models.Cash:     30_525,
		},
	},
	{
		CaseName: "EndYearIsMultipleYearsAway_MultipleAssets_WithConstantContributionsStoppingPartway_WithInflation_WithReturns",
		ForecastRequest: models.ForecastPortfolioRequest{
			EndYear:             5,
			AnnualInflationRate: 0.05,
			AnnualPortfolioBalanceChanges: []models.AnnualPortfolioBalanceChange{
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
			PortfolioAllocation: models.PortfolioAllocation{
				models.Equities: {
					ReturnRate: 0.07,
					Allocation: 0.9,
				},
				models.Cash: {
					ReturnRate: 0.05,
					Allocation: 0.1,
				},
			},
			InitPortfolio: models.Portfolio{
				models.Equities: 200_000,
			},
			RebalanceCadence:    1,
			RebalancingStrategy: models.YearlyToZero,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 317_279,
			models.Cash:     10_000,
		},
	},
	{
		CaseName: "EndYearIsMultipleYearsAway_MultipleAssets_WithDecliningContributionsStoppingPartway_WithInflation_WithReturns",
		ForecastRequest: models.ForecastPortfolioRequest{
			EndYear:             5,
			AnnualInflationRate: 0.05,
			AnnualPortfolioBalanceChanges: []models.AnnualPortfolioBalanceChange{
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
			PortfolioAllocation: models.PortfolioAllocation{
				models.Equities: {
					ReturnRate: 0.07,
					Allocation: 0.9,
				},
				models.Cash: {
					ReturnRate: 0.05,
					Allocation: 0.1,
				},
			},
			InitPortfolio: models.Portfolio{
				models.Equities: 200_000,
			},
			RebalanceCadence:    1,
			RebalancingStrategy: models.YearlyToZero,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 312_504,
			models.Cash:     9500,
		},
	},
	{
		CaseName: "EndYearIsMultipleYearsAway_MultipleAssets_WithDecliningWithdrawalsStoppingPartway_WithInflation",
		ForecastRequest: models.ForecastPortfolioRequest{
			EndYear:             5,
			AnnualInflationRate: 0.05,
			AnnualPortfolioBalanceChanges: []models.AnnualPortfolioBalanceChange{
				{
					Amount:          -10_000,
					StartYear:       0,
					EndYear:         2,
					AnnualPctChange: -0.1,
				},
				{
					Amount:          -40_000,
					StartYear:       0,
					EndYear:         2,
					AnnualPctChange: -0.1,
				},
			},
			PortfolioAllocation: models.PortfolioAllocation{
				models.Equities: {
					ReturnRate: 0.07,
					Allocation: 0.9,
				},
				models.Cash: {
					ReturnRate: 0.05,
					Allocation: 0.1,
				},
			},
			InitPortfolio: models.Portfolio{
				models.Equities: 200_000,
			},
			RebalanceCadence:    1,
			RebalancingStrategy: models.YearlyToZero,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 118_940,
			models.Cash:     0,
		},
	},

	{
		CaseName: "EndYearIsMultipleYearsAway_MultipleAssets_WithIncreasingContributionsAneDecliningWithdrawalsStoppingPartway_WithInflation_WithReturns",
		ForecastRequest: models.ForecastPortfolioRequest{
			EndYear:             5,
			AnnualInflationRate: 0.05,
			AnnualPortfolioBalanceChanges: []models.AnnualPortfolioBalanceChange{
				{
					Amount:          30_000,
					StartYear:       1,
					EndYear:         2,
					AnnualPctChange: 0.1,
				},
				{
					Amount:          40_000,
					StartYear:       1,
					EndYear:         2,
					AnnualPctChange: 0.1,
				},
				{
					Amount:          -10_000,
					StartYear:       0,
					EndYear:         2,
					AnnualPctChange: -0.1,
				},
				{
					Amount:          -35_000,
					StartYear:       0,
					EndYear:         2,
					AnnualPctChange: -0.1,
				},
			},
			PortfolioAllocation: models.PortfolioAllocation{
				models.Equities: {
					ReturnRate: 0.07,
					Allocation: 0.9,
				},
				models.Cash: {
					ReturnRate: 0.05,
					Allocation: 0.1,
				},
			},
			InitPortfolio: models.Portfolio{
				models.Equities: 200_000,
			},
			RebalanceCadence:    1,
			RebalancingStrategy: models.YearlyToZero,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 273_345,
			models.Cash:     5450,
		},
	},
	{
		CaseName: "LongTimeFrameWithEverythingSet",
		ForecastRequest: models.ForecastPortfolioRequest{
			EndYear:             60,
			AnnualInflationRate: 0.02,
			AnnualPortfolioBalanceChanges: []models.AnnualPortfolioBalanceChange{
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
			PortfolioAllocation: models.PortfolioAllocation{
				models.Equities: {
					ReturnRate: 0.07,
					Allocation: 0.89,
				},
				models.Bonds: {
					ReturnRate: 0.04,
					Allocation: 0.08,
				},
				models.Cash: {
					ReturnRate: 0.03,
					Allocation: 0.03,
				},
			},
			InitPortfolio: models.Portfolio{
				models.Equities: 200_000,
			},
			RebalanceCadence:    1,
			RebalancingStrategy: models.YearlyToZero,
		},
		EndPortfolio: models.Portfolio{
			models.Equities: 19_470_312,
			models.Bonds:    456_206,
			models.Cash:     122_504,
		},
	},
}

func TestForecastFuturePortfolioValueByYearSuccessCases(t *testing.T) {
	for _, test := range simulationSuccessCases {
		t.Run(test.CaseName, func(t *testing.T) {
			forecastedPortfoliosByYear, _ := ForecastFuturePortfolioValueByYear(test.ForecastRequest)
			forecastedPortfolioForEndYear := forecastedPortfoliosByYear.Portfolios[len(forecastedPortfoliosByYear.Portfolios)-1]
			for assetType, expectedVal := range test.EndPortfolio {
				actualVal, ok := forecastedPortfolioForEndYear[assetType]
				if !ok || !test_utils.AlmostEqual(actualVal, expectedVal) {
					t.Errorf("Expected %v but got %v", test.EndPortfolio, forecastedPortfolioForEndYear)
					t.FailNow()
				}
			}
		})
	}
}

var simulationErrorCases = []PortfolioSimulatorTestCase{
	{
		CaseName: "AllocationsDoNotSumUpToOne",
		ForecastRequest: models.ForecastPortfolioRequest{
			EndYear:             1,
			AnnualInflationRate: 0.0,
			AnnualPortfolioBalanceChanges: []models.AnnualPortfolioBalanceChange{
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
			PortfolioAllocation: models.PortfolioAllocation{
				models.Equities: {
					ReturnRate: 0.0,
					Allocation: 0.9,
				},
				models.Cash: {
					ReturnRate: 0.0,
					Allocation: 0.15,
				},
			},
			InitPortfolio:       models.Portfolio{},
			RebalanceCadence:    1,
			RebalancingStrategy: models.YearlyToZero,
		},
		ErrorMessage: "portfolio allocation percent must sum up to 1",
	},
	{
		CaseName: "AnnualContributionStopsAfterEndYear",
		ForecastRequest: models.ForecastPortfolioRequest{
			EndYear:             1,
			AnnualInflationRate: 0.0,
			AnnualPortfolioBalanceChanges: []models.AnnualPortfolioBalanceChange{
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
			InitPortfolio:       models.Portfolio{},
			RebalanceCadence:    1,
			RebalancingStrategy: models.YearlyToZero,
		},
		ErrorMessage: "annual balance change end year must be less than or equal to last year",
	},
}

func TestForecastFuturePortfolioValueByYearErrorCases(t *testing.T) {
	for _, test := range simulationErrorCases {
		t.Run(test.CaseName, func(t *testing.T) {
			_, err := ForecastFuturePortfolioValueByYear(test.ForecastRequest)
			if err.Error() != test.ErrorMessage {
				t.Errorf("expected %v but got %v", test.ErrorMessage, err.Error())
			}
		})
	}
}
