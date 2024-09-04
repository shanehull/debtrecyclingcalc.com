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

	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)

		c := templates.NotFound()
		err := templates.Layout(c, "Not Fountempl.WithStatus(http.StatusNotFound)d", buildinfo.GitTag, buildinfo.BuildDate).
			Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	params := &calc.DebtRecyclingParameters{
		Salary:               150000,
		InitialInvestment:    100000,
		AnnualInvestment:     50000,
		MortgageSize:         600000,
		MortgageInterestRate: 0.05,
		CapitalGrowthRate:    0.08,
		NumYears:             10,
		Country:              "au",
		ReinvestDividends:    true,
		ReinvestTaxRefunds:   true,
	}

	data, err := calc.DebtRecycling(*params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	positionsChart, err := charts.Positions(data, params.NumYears, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	incomeChart, err := charts.Income(data, params.NumYears, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	interestChart, err := charts.Interest(data, params.NumYears, r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	index := templates.Index(
		templates.Hero(),
		templates.Form(),
		templates.Results(data, params, positionsChart, incomeChart, interestChart),
	)

	err = templates.Layout(index, "Debt Recycling Calculator", buildinfo.GitTag, buildinfo.BuildDate).
		Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
