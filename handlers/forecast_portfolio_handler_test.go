package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/guilam34/financial_planner/models"
	"github.com/guilam34/financial_planner/test_utils"
)

func TestForecastPortfolio(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		portfolioRequest := models.ForecastPortfolioRequest{
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
		}
		portfolioRequestBuf := new(bytes.Buffer)
		json.NewEncoder(portfolioRequestBuf).Encode(portfolioRequest)

		expectedResponse := models.ForecastPortfolioResponse{
			Portfolios: []models.Portfolio{
				models.Portfolio{
					models.Equities: 200_000,
				},
				models.Portfolio{
					models.Equities: 245_000,
					models.Cash:     5_000,
				},
			},
		}

		request, _ := http.NewRequest(http.MethodGet, "/forecastPortfolio", portfolioRequestBuf)
		response := httptest.NewRecorder()

		ForecastPortfolioHandler(response, request)

		var actualResponse models.ForecastPortfolioResponse
		json.NewDecoder(response.Body).Decode(&actualResponse)

		if len(actualResponse.Portfolios) != len(expectedResponse.Portfolios) {
			t.Errorf("expected %v but got %v", expectedResponse, actualResponse)
			t.FailNow()
		}

		for i := 0; i < len(actualResponse.Portfolios); i++ {
			expectedPortfolio := expectedResponse.Portfolios[i]
			actualPortfolio := actualResponse.Portfolios[i]

			for assetType, expectedVal := range expectedPortfolio {
				actualVal := actualPortfolio[assetType]
				if !test_utils.AlmostEqual(actualVal, expectedVal) {
					t.Errorf("expected %v but got %v", expectedResponse, actualResponse)
					t.FailNow()
				}
			}
		}
	})
}

func TestForecastPortfolioWithIncompleteRequest(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		portfolioRequestBuf := new(bytes.Buffer)
		portfolioRequest := models.ForecastPortfolioRequest{
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
		}
		json.NewEncoder(portfolioRequestBuf).Encode(portfolioRequest)

		request, _ := http.NewRequest(http.MethodGet, "/forecastPortfolio", portfolioRequestBuf)
		response := httptest.NewRecorder()

		ForecastPortfolioHandler(response, request)

		if response.Result().StatusCode != 400 {
			t.Errorf("expected 400 but got %d", response.Code)
		}
	})
}

func TestForecastPortfolioWithMalformedRequest(t *testing.T) {
	t.Run("returns Pepper's score", func(t *testing.T) {
		portfolioRequestBuf := new(bytes.Buffer)
		portfolioRequestBuf.WriteString("INVALID_REQUEST")

		request, _ := http.NewRequest(http.MethodGet, "/forecastPortfolio", portfolioRequestBuf)
		response := httptest.NewRecorder()

		ForecastPortfolioHandler(response, request)

		if response.Result().StatusCode != 400 {
			t.Errorf("expected 400 but got %d", response.Code)
		}
	})
}
