package simulator

import "github.com/guilam34/financial_planner/models"

type RebalancingStrategy interface {
	Rebalance(
		portfolio models.Portfolio,
		portfolioAllocation models.PortfolioAllocation, year int) models.Portfolio
}

type RebalanceToZero struct{}

func (r RebalanceToZero) Rebalance(
	portfolio models.Portfolio,
	portfolioAllocation models.PortfolioAllocation,
	year int) models.Portfolio {

	portfolioValue, positiveValAssetTypes, negativeValAssetTypes := getNetPortfolioValue(portfolio)

	// Only rebalance if we're not in the negative
	if portfolioValue < 0.0 {
		return portfolio
	}

	rebalancedPortfolio := models.Portfolio{}
	for key, value := range portfolio {
		rebalancedPortfolio[key] = value
	}

	for _, assetType := range negativeValAssetTypes {
		assetVal := portfolio[assetType]
		amtToSubtractFromEachPosAsset := assetVal / float64(len(positiveValAssetTypes))
		rebalancedPortfolio[assetType] = 0
		for _, assetType := range positiveValAssetTypes {
			rebalancedPortfolio[assetType] = rebalancedPortfolio[assetType] + amtToSubtractFromEachPosAsset
		}
	}
	return rebalancedPortfolio
}

type RebalanceEveryNYears struct {
	rebalanceCadence int
}

func (r RebalanceEveryNYears) Rebalance(
	portfolio models.Portfolio,
	portfolioAllocation models.PortfolioAllocation,
	year int) models.Portfolio {

	portfolioValue, _, negativeValAssetTypes := getNetPortfolioValue(portfolio)

	// Only rebalance if we're not in the negative
	if portfolioValue < 0.0 {
		return portfolio
	}

	if year%r.rebalanceCadence == 0 {
		rebalancedPortfolio := models.Portfolio{}
		for assetType, allocation := range portfolioAllocation {
			rebalancedPortfolio[assetType] = portfolioValue * allocation.Allocation
		}
		return rebalancedPortfolio
	}

	if len(negativeValAssetTypes) > 0 {
		return RebalanceToZero{}.Rebalance(portfolio, portfolioAllocation, year)
	}

	return portfolio
}

func getNetPortfolioValue(
	portfolio models.Portfolio) (
	portfolioValue float64,
	positiveValAssetTypes []models.AssetType,
	negativeValAssetTypes []models.AssetType) {

	portfolioValue = 0.0
	positiveValAssetTypes = []models.AssetType{}
	negativeValAssetTypes = []models.AssetType{}
	for idx, assetVal := range portfolio {
		portfolioValue = portfolioValue + assetVal
		if assetVal > 0 {
			positiveValAssetTypes = append(positiveValAssetTypes, idx)
		} else if assetVal < 0 {
			negativeValAssetTypes = append(negativeValAssetTypes, idx)
		}
	}
	return
}
