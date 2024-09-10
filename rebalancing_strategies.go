package main

type RebalancingStrategy interface {
	Rebalance(portfolio Portfolio, year int) Portfolio
}

type RebalanceToZero struct{}

func (r RebalanceToZero) Rebalance(portfolio Portfolio, year int) Portfolio {
	portfolioValue, positiveValAssetTypes, negativeValAssetTypes := getNetPortfolioValue(portfolio)

	// Only rebalance if we're not in the negative
	if portfolioValue < 0.0 {
		return portfolio
	}

	rebalancedPortfolio := Portfolio{}
	for key, value := range portfolio {
		rebalancedPortfolio[key] = value
	}

	for _, assetType := range negativeValAssetTypes {
		assetVal := portfolio[assetType]
		amtToSubtractFromEachPosAsset := assetVal / float64(len(positiveValAssetTypes))
		rebalancedPortfolio[assetType] = 0
		for _, assetType := range positiveValAssetTypes {
			rebalancedPortfolio[assetType] = rebalancedPortfolio[assetType] - amtToSubtractFromEachPosAsset
		}
	}
	return rebalancedPortfolio
}

type RebalanceEveryNYears struct {
	n int
}

func (r RebalanceEveryNYears) Rebalance(portfolio Portfolio, year int) {

}

func getNetPortfolioValue(
	portfolio Portfolio) (
	portfolioValue float64,
	positiveValAssetTypes []assetType,
	negativeValAssetTypes []assetType) {

	portfolioValue = 0.0
	positiveValAssetTypes = []assetType{}
	negativeValAssetTypes = []assetType{}
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
