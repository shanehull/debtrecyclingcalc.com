package handlers

import (
	"fmt"
	"net/http"

	"debtrecyclingcalc.com/internal/calc"
	"debtrecyclingcalc.com/internal/charts"
	"debtrecyclingcalc.com/internal/templates"
)

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}

	// Parse the form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "error parsing form data", http.StatusBadRequest)
		return
	}

	params, err := getFormParams(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// if params is empty respond with error
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

	results := templates.Results(data, params, positionsChart, incomeChart, interestChart)

	w.Header().
		Set("HX-Push-Url", fmt.Sprintf("/?salary=%.2f"+
			"&initial_investment=%.2f"+
			"&annual_investment=%.2f"+
			"&mortgage_size=%.2f"+
			"&mortgage_interest_rate=%.2f"+
			"&dividend_return_rate=%.2f"+
			"&capital_growth_rate=%.2f"+
			"&years=%d"+
			"&country=%s"+
			"&reinvest_dividends=%t"+
			"&reinvest_tax_refunds=%t",
			params.Salary,
			params.InitialInvestment,
			params.AnnualInvestment,
			params.MortgageSize,
			params.MortgageInterestRate*100,
			params.DividendReturnRate*100,
			params.CapitalGrowthRate*100,
			params.NumYears,
			params.Country,
			params.ReinvestDividends,
			params.ReinvestTaxRefunds,
		))

	err = results.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
