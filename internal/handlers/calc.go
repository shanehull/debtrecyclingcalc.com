package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"debtrecyclingcalculator.com.au/internal/calc"
	"debtrecyclingcalculator.com.au/internal/charts"
	"debtrecyclingcalculator.com.au/internal/templates"
)

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
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
		ReinvestDividends:    reinvestDividends,
		ReinvestTaxRefunds:   reinvestTaxRefunds,
		NumYears:             years,
	}

	data := calc.DebtRecycling(*params)

	lineChart, err := charts.StackedLineChart(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	results := templates.Results(data, params, lineChart)

	err = results.Render(r.Context(), w)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
