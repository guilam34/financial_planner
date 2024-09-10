package handlers

import (
	"net/http"

	"github.com/guilam34/financial_planner/models"
	"github.com/guilam34/financial_planner/simulator"
)

func ForecastPortfolioHandler(w http.ResponseWriter, r *http.Request) {
	req, err := decode[models.ForecastPortfolioRequest](r)
	if err != nil {
		encodeError(w, 400, err)
	}
	forecast, forecastErr := simulator.ForecastFuturePortfolioValueByYear(req)
	if forecastErr != nil {
		encodeError(w, 400, forecastErr)
	} else {
		encode(w, 500, forecast)
	}
}
