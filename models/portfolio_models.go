package models

type AnnualPortfolioBalanceChange struct {
	Amount          float64
	StartYear       int
	EndYear         int
	AnnualPctChange float64
}

type AssetType int

const (
	Equities AssetType = iota
	Bonds
	Cash
)

type AssetAllocation struct {
	ReturnRate float64
	Allocation float64
}

type PortfolioAllocation map[AssetType]AssetAllocation

type Portfolio map[AssetType]float64

type RebalancingStrategyEnum int

const (
	YearlyToZero RebalancingStrategyEnum = iota
	EveryNYearsByAlloc
)
