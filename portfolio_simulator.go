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

func ForecastFuturePortfolioValueByYear(initPortfolio Portfolio, annualContributions []AnnualContribution, portfolioAllocationWitNominalRates PortfolioAllocation, annualInflationRate float64, years int) ([]Portfolio, error) {
	for _, annualContribution := range annualContributions {
		if annualContribution.EndYear > years {
			return nil, errors.New("Annual contribution end year must be less than or equal to last year")
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

	var result []Portfolio
	for year := 0; year <= years; year++ {
		result = append(result, forecastFuturePortfolioValue(initPortfolio, annualContributions, portfolioAllocationWithRealRates, year))
	}
	return result, nil
}

func convertToRealRate(nominalRate float64, inflationRate float64) float64 {
	return nominalRate - inflationRate
}

func forecastFuturePortfolioValue(initPortfolio Portfolio, annualContributions []AnnualContribution, portfolioAllocationWitRealRates PortfolioAllocation, year int) Portfolio {
	forecastedPortfolio := Portfolio{}

	for assetType, initAssetVal := range initPortfolio {
		forecastedPortfolio[assetType] = initAssetVal * math.Pow(1+portfolioAllocationWitRealRates[assetType].ReturnRate, float64(year))
	}

	for curYear := 1; curYear <= year; curYear++ {
		for _, contrib := range annualContributions {
			if curYear >= contrib.StartYear && curYear <= contrib.EndYear {
				contribAmtAdjustedForChangePct := contrib.Amount
				// Adjust for change in contribution after the first year
				if curYear > contrib.StartYear {
					contribAmtAdjustedForChangePct = contrib.Amount * math.Pow(1+contrib.AnnualPctChange, float64(curYear-contrib.StartYear-1))
				}
				for assetType, assetAllocation := range portfolioAllocationWitRealRates {
					forecastedPortfolio[assetType] = forecastedPortfolio[assetType] + float64(contribAmtAdjustedForChangePct)*assetAllocation.Allocation*math.Pow(1+assetAllocation.ReturnRate, float64(year-curYear))
				}
			}
		}
	}
	return forecastedPortfolio
}
