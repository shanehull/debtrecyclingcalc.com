package handlers

import (
	"net/http"

	"debtrecyclingcalc.com/internal/buildinfo"
	"debtrecyclingcalc.com/internal/calc"
	"debtrecyclingcalc.com/internal/charts"
	"debtrecyclingcalc.com/internal/templates"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	data := &calc.DebtRecyclingData{
		PortfolioValue:  []float64{0},
		DividendReturns: []float64{0},
		TaxRefunds:      []float64{0},
	}
	params := &calc.DebtRecyclingParameters{}

	lineChart, err := charts.StackedLineChart(data, params.NumYears)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	index := templates.Index(templates.Form(), templates.Results(data, params, lineChart))

	err = templates.Layout(index, "Debt Recycling Calculator", buildinfo.GitTag, buildinfo.BuildDate).
		Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
