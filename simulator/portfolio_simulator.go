package simulator

import (
	"errors"
	"math"

	"github.com/guilam34/financial_planner/models"
)

func ForecastFuturePortfolioValueByYear(forecastRequest models.ForecastPortfolioRequest) (models.ForecastPortfolioResponse, error) {

	for _, balanceChange := range forecastRequest.AnnualPortfolioBalanceChanges {
		if balanceChange.EndYear > forecastRequest.EndYear {
			return models.ForecastPortfolioResponse{}, errors.New("annual balance change end year must be less than or equal to last year")
		}
	}

	allocatedPortfolioPct := 0.0
	for _, allocation := range forecastRequest.PortfolioAllocation {
		allocatedPortfolioPct = allocatedPortfolioPct + allocation.Allocation
	}
	if allocatedPortfolioPct != 1.0 {
		return models.ForecastPortfolioResponse{}, errors.New("portfolio allocation percent must sum up to 1")
	}

	var rebalancingStrategy RebalancingStrategy
	switch forecastRequest.RebalancingStrategy {
	case models.YearlyToZero:
		rebalancingStrategy = RebalanceToZero{}
		break
	case models.EveryNYearsByAlloc:
		rebalancingStrategy = RebalanceEveryNYears{rebalanceCadence: forecastRequest.RebalanceCadence}
		break
	default:
		rebalancingStrategy = RebalanceToZero{}
		break
	}

	result := []models.Portfolio{forecastRequest.InitPortfolio}
	prevPortfolio := forecastRequest.InitPortfolio
	for year := 1; year <= forecastRequest.EndYear; year++ {
		curPortfolio := forecastNextYearPortfolio(
			prevPortfolio,
			forecastRequest.AnnualPortfolioBalanceChanges,
			convertToRealRates(forecastRequest.PortfolioAllocation, forecastRequest.AnnualInflationRate),
			year,
			rebalancingStrategy)
		result = append(result, curPortfolio)
		prevPortfolio = curPortfolio
	}
	return models.ForecastPortfolioResponse{Portfolios: result}, nil
}

func convertToRealRates(portfolioAllocation models.PortfolioAllocation, inflationRate float64) models.PortfolioAllocation {
	portfolioAllocationWithRealRates := models.PortfolioAllocation{}
	for assetType, allocation := range portfolioAllocation {
		portfolioAllocationWithRealRates[assetType] = models.AssetAllocation{
			ReturnRate: allocation.ReturnRate - inflationRate,
			Allocation: allocation.Allocation,
		}
	}
	return portfolioAllocationWithRealRates
}

func forecastNextYearPortfolio(
	prevPortfolio models.Portfolio,
	annualPortfolioBalanceChanges []models.AnnualPortfolioBalanceChange,
	portfolioAllocationWitRealRates models.PortfolioAllocation,
	year int,
	rebalancingStrategy RebalancingStrategy) models.Portfolio {

	forecastedPortfolio := models.Portfolio{}

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
