package handlers

import (
	"fmt"
	"net/http"
	"net/url"
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

	params, err := getParamsFromForm(r.Form)
	if err != nil {
		http.Error(w, "error parsing form params", http.StatusBadRequest)
		return
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

func getParamsFromForm(form url.Values) (*calc.Parameters, error) {
	parseFloat := func(key string) (float64, error) {
		return strconv.ParseFloat(form.Get(key), 64)
	}

	parseInt := func(key string) (int, error) {
		return strconv.Atoi(form.Get(key))
	}

	salary, err := parseFloat("salary")
	if err != nil {
		return nil, fmt.Errorf("error parsing salary: %w", err)
	}

	initialInvestmentAmount, err := parseFloat("initial_investment")
	if err != nil {
		return nil, fmt.Errorf("error parsing initial investment amount: %w", err)
	}

	annualInvestmentAmount, err := parseFloat("annual_investment")
	if err != nil {
		return nil, fmt.Errorf("error parsing annual investment amount: %w", err)
	}

	mortgageSize, err := parseFloat("mortgage_size")
	if err != nil {
		return nil, fmt.Errorf("error parsing mortgage size: %w", err)
	}

	mortgageInterestRate, err := parseFloat("mortgage_interest_rate")
	if err != nil {
		return nil, fmt.Errorf("error parsing mortgage interest rate: %w", err)
	}

	dividendReturnRate, err := parseFloat("dividend_return_rate")
	if err != nil {
		return nil, fmt.Errorf("error parsing dividend return rate: %w", err)
	}

	capitalGrowthRate, err := parseFloat("capital_growth_rate")
	if err != nil {
		return nil, fmt.Errorf("error parsing capital growth rate: %w", err)
	}

	years, err := parseInt("years")
	if err != nil {
		return nil, fmt.Errorf("error parsing years: %w", err)
	}

	country := form.Get("country")
	reinvestDividends := form.Get("reinvest_dividends") == "true"
	reinvestTaxRefunds := form.Get("reinvest_tax_refunds") == "true"

	return &calc.Parameters{
		Salary:               salary,
		InitialInvestment:    initialInvestmentAmount,
		AnnualInvestment:     annualInvestmentAmount,
		MortgageSize:         mortgageSize,
		MortgageInterestRate: mortgageInterestRate / 100,
		DividendReturnRate:   dividendReturnRate / 100,
		CapitalGrowthRate:    capitalGrowthRate / 100,
		NumYears:             years,
		Country:              country,
		ReinvestDividends:    reinvestDividends,
		ReinvestTaxRefunds:   reinvestTaxRefunds,
	}, nil
}
