package main

import (
	"errors"
	"math"
)

type AnnualPortfolioBalanceChange struct {
	Amount          float64
	StartYear       int
	EndYear         int
	AnnualPctChange float64
}

type AnnualContribution AnnualPortfolioBalanceChange
type AnnualWithdrawal AnnualPortfolioBalanceChange

type assetType int

const (
	Equities assetType = iota
	Bonds
	Cash
)

type assetAllocation struct {
	ReturnRate float64
	Allocation float64
}

type PortfolioAllocation map[assetType]assetAllocation

type assetValue struct {
	AssetType assetType
	Amount    float64
}

type Portfolio map[assetType]float64

func ForecastFuturePortfolioValueByYear(
	initPortfolio Portfolio,
	annualContributions []AnnualContribution,
	annualWithdrawals []AnnualWithdrawal,
	portfolioAllocationWitNominalRates PortfolioAllocation,
	annualInflationRate float64,
	years int,
	rebalancingStrategy RebalancingStrategy) ([]Portfolio, error) {

	for _, annualContribution := range annualContributions {
		if annualContribution.EndYear > years {
			return nil, errors.New("Annual contribution end year must be less than or equal to last year")
		}
	}

	for _, annualWithdrawal := range annualWithdrawals {
		if annualWithdrawal.EndYear > years {
			return nil, errors.New("Annual withdrawal end year must be less than or equal to last year")
		}
	}

	allocatedPortfolioPct := 0.0
	for _, allocation := range portfolioAllocationWitNominalRates {
		allocatedPortfolioPct = allocatedPortfolioPct + allocation.Allocation
	}
	if allocatedPortfolioPct != 1.0 {
		return nil, errors.New("Portfolio allocation percent must sum up to 1")
	}

	portfolioAllocationWithRealRates := PortfolioAllocation{}
	for assetType, allocation := range portfolioAllocationWitNominalRates {
		portfolioAllocationWithRealRates[assetType] = assetAllocation{
			ReturnRate: convertToRealRate(allocation.ReturnRate, annualInflationRate),
			Allocation: allocation.Allocation,
		}
	}

	var result []Portfolio = []Portfolio{initPortfolio}
	var prevPortfolio Portfolio = initPortfolio
	for year := 1; year <= years; year++ {
		curPortfolio := forecastNextYearPortfolio(prevPortfolio, annualContributions, annualWithdrawals, portfolioAllocationWithRealRates, year, rebalancingStrategy)
		result = append(result, curPortfolio)
		prevPortfolio = curPortfolio
	}
	return result, nil
}

func convertToRealRate(nominalRate float64, inflationRate float64) float64 {
	return nominalRate - inflationRate
}

func forecastNextYearPortfolio(
	prevPortfolio Portfolio,
	annualContributions []AnnualContribution,
	annualWithdrawals []AnnualWithdrawal,
	portfolioAllocationWitRealRates PortfolioAllocation,
	year int,
	rebalancingStrategy RebalancingStrategy) Portfolio {

	forecastedPortfolio := Portfolio{}

	for assetType, prevAssetVal := range prevPortfolio {
		forecastedPortfolio[assetType] = prevAssetVal * (1 + portfolioAllocationWitRealRates[assetType].ReturnRate)
	}

	for _, contrib := range annualContributions {
		if year >= contrib.StartYear && year <= contrib.EndYear {
			contribAmtAdjustedForChangePct := contrib.Amount
			// Adjust for change in contribution after the first year
			if year > contrib.StartYear {
				contribAmtAdjustedForChangePct = contrib.Amount * math.Pow(1+contrib.AnnualPctChange, float64(year-contrib.StartYear-1))
			}
			for assetType, assetAllocation := range portfolioAllocationWitRealRates {
				forecastedPortfolio[assetType] = forecastedPortfolio[assetType] + float64(contribAmtAdjustedForChangePct)*assetAllocation.Allocation
			}
		}
	}

	for _, withdrawal := range annualWithdrawals {
		if year >= withdrawal.StartYear && year <= withdrawal.EndYear {
			withdrawalAmtAdjustedForChangePct := withdrawal.Amount
			// Adjust for change in withdrawal after the first year
			if year > withdrawal.StartYear {
				withdrawalAmtAdjustedForChangePct = withdrawal.Amount * math.Pow(1+withdrawal.AnnualPctChange, float64(year-withdrawal.StartYear-1))
			}
			for assetType, assetAllocation := range portfolioAllocationWitRealRates {
				forecastedPortfolio[assetType] = forecastedPortfolio[assetType] - float64(withdrawalAmtAdjustedForChangePct)*assetAllocation.Allocation
			}
		}
	}
	return rebalancingStrategy.Rebalance(forecastedPortfolio)
}
