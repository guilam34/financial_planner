package routes

import (
	"net/http"

	"github.com/guilam34/financial_planner/handlers"
)

func AddRoutes(mux *http.ServeMux) {
	mux.HandleFunc(
		"/forecastPortfolio", handlers.ForecastPortfolioHandler,
	)
}
