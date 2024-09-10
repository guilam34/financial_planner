package models

type ForecastPortfolioRequest struct {
	InitPortfolio                 Portfolio
	AnnualPortfolioBalanceChanges []AnnualPortfolioBalanceChange
	PortfolioAllocation           PortfolioAllocation
	AnnualInflationRate           float64
	EndYear                       int
	RebalanceCadence              int
	RebalancingStrategy           RebalancingStrategyEnum
}

type ForecastPortfolioResponse struct {
	Portfolios []Portfolio
}
