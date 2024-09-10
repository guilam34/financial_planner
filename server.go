package main

import (
	"log"
	"net/http"

	"github.com/guilam34/financial_planner/routes"
)

func run() {
	mux := http.NewServeMux()
	routes.AddRoutes(mux)
	log.Println("Listening....")
	http.ListenAndServe(":3000", mux)
}

func main() {
	run()
}
