package handlers

import (
	"net/http"
	"strconv"

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

	salary, err := strconv.ParseFloat(r.Form.Get("salary"), 64)
	if err != nil {
		http.Error(w, "error parsing salary", http.StatusBadRequest)
		return
	}

	inititalInvestmentAmount, err := strconv.ParseFloat(r.Form.Get("initial_investment"), 64)
	if err != nil {
		http.Error(w, "error parsing initial investment amount", http.StatusBadRequest)
		return
	}

	annualInvestmentAmount, err := strconv.ParseFloat(r.Form.Get("annual_investment"), 64)
	if err != nil {
		http.Error(w, "error parsing annual investment amount", http.StatusBadRequest)
		return
	}

	mortgageSize, err := strconv.ParseFloat(r.Form.Get("mortgage_size"), 64)
	if err != nil {
		http.Error(w, "error parsing mortgage size", http.StatusBadRequest)
		return
	}

	mortgageInterestRate, err := strconv.ParseFloat(r.Form.Get("mortgage_interest_rate"), 64)
	if err != nil {
		http.Error(w, "error parsing mortgage interest rate", http.StatusBadRequest)
		return
	}

	dividendReturnRate, err := strconv.ParseFloat(r.Form.Get("dividend_return_rate"), 64)
	if err != nil {
		http.Error(w, "error parsing dividend return rate", http.StatusBadRequest)
		return
	}

	capitalGrowthRate, err := strconv.ParseFloat(r.Form.Get("capital_growth_rate"), 64)
	if err != nil {
		http.Error(w, "error parsing capital growth rate", http.StatusBadRequest)
		return
	}

	years, err := strconv.Atoi(r.Form.Get("years"))
	if err != nil {
		http.Error(w, "error parsing years", http.StatusBadRequest)
		return
	}

	country := r.Form.Get("country")

	reinvestDividends := r.Form.Get("reinvest_dividends") == "on"
	reinvestTaxRefunds := r.Form.Get("reinvest_tax_refunds") == "on"

	params := &calc.DebtRecyclingParameters{
		Salary:               salary,
		InitialInvestment:    inititalInvestmentAmount,
		AnnualInvestment:     annualInvestmentAmount,
		MortgageSize:         mortgageSize,
		MortgageInterestRate: mortgageInterestRate / 100,
		DividendReturnRate:   dividendReturnRate / 100,
		CapitalGrowthRate:    capitalGrowthRate / 100,
		NumYears:             years,
		Country:              country,
		ReinvestDividends:    reinvestDividends,
		ReinvestTaxRefunds:   reinvestTaxRefunds,
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

	err = results.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
