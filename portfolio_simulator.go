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

type Portfolio map[assetType]float64

func ForecastFuturePortfolioValueByYear(
	initPortfolio Portfolio,
	annualPortfolioBalanceChanges []AnnualPortfolioBalanceChange,
	portfolioAllocationWitNominalRates PortfolioAllocation,
	annualInflationRate float64,
	years int,
	rebalancingStrategy RebalancingStrategy) ([]Portfolio, error) {

	for _, balanceChange := range annualPortfolioBalanceChanges {
		if balanceChange.EndYear > years {
			return nil, errors.New("annual balance change end year must be less than or equal to last year")
		}
	}

	allocatedPortfolioPct := 0.0
	for _, allocation := range portfolioAllocationWitNominalRates {
		allocatedPortfolioPct = allocatedPortfolioPct + allocation.Allocation
	}
	if allocatedPortfolioPct != 1.0 {
		return nil, errors.New("portfolio allocation percent must sum up to 1")
	}

	var result []Portfolio = []Portfolio{initPortfolio}
	var prevPortfolio Portfolio = initPortfolio
	for year := 1; year <= years; year++ {
		curPortfolio := forecastNextYearPortfolio(
			prevPortfolio,
			annualPortfolioBalanceChanges,
			convertToRealRates(portfolioAllocationWitNominalRates, annualInflationRate),
			year,
			rebalancingStrategy)
		result = append(result, curPortfolio)
		prevPortfolio = curPortfolio
	}
	return result, nil
}

func convertToRealRates(portfolioAllocation PortfolioAllocation, inflationRate float64) PortfolioAllocation {
	portfolioAllocationWithRealRates := PortfolioAllocation{}
	for assetType, allocation := range portfolioAllocation {
		portfolioAllocationWithRealRates[assetType] = assetAllocation{
			ReturnRate: allocation.ReturnRate - inflationRate,
			Allocation: allocation.Allocation,
		}
	}
	return portfolioAllocationWithRealRates
}

func forecastNextYearPortfolio(
	prevPortfolio Portfolio,
	annualPortfolioBalanceChanges []AnnualPortfolioBalanceChange,
	portfolioAllocationWitRealRates PortfolioAllocation,
	year int,
	rebalancingStrategy RebalancingStrategy) Portfolio {

	forecastedPortfolio := Portfolio{}

	for assetType, prevAssetVal := range prevPortfolio {
		forecastedPortfolio[assetType] = prevAssetVal * (1 + portfolioAllocationWitRealRates[assetType].ReturnRate)
	}

	for _, contrib := range annualPortfolioBalanceChanges {
		if year >= contrib.StartYear && year <= contrib.EndYear {
			amtAdjustedForChangePct := contrib.Amount
			// Adjust for change in contribution after the first year
			if year > contrib.StartYear {
				amtAdjustedForChangePct = contrib.Amount * math.Pow(1+contrib.AnnualPctChange, float64(year-contrib.StartYear-1))
			}
			for assetType, assetAllocation := range portfolioAllocationWitRealRates {
				forecastedPortfolio[assetType] = forecastedPortfolio[assetType] + float64(amtAdjustedForChangePct)*assetAllocation.Allocation
			}
		}
	}
	return rebalancingStrategy.Rebalance(forecastedPortfolio, portfolioAllocationWitRealRates, year)
}
