package handlers

import (
	"net/http"
	"net/url"
	"strconv"

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
		err := templates.Layout(c, "Not Found", buildinfo.GitTag, buildinfo.BuildDate).
			Render(r.Context(), w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	params := &calc.Parameters{
		Salary:               150000,
		InitialInvestment:    100000,
		AnnualInvestment:     50000,
		MortgageSize:         600000,
		MortgageInterestRate: 0.0500,
		DividendReturnRate:   0.0200,
		CapitalGrowthRate:    0.0800,
		NumYears:             10,
		Country:              "au",
		ReinvestDividends:    true,
		ReinvestTaxRefunds:   true,
	}

	query := r.URL.Query()
	if len(query) != 0 {
		setParamsFromQuery(params, query)
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
		templates.Form(params),
		templates.Results(data, params, positionsChart, incomeChart, interestChart),
	)

	err = templates.Layout(index, "Debt Recycling Calculator", buildinfo.GitTag, buildinfo.BuildDate).
		Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func setParamsFromQuery(params *calc.Parameters, query url.Values) {
	parseFloat := func(key string) (float64, error) {
		if val := query.Get(key); val != "" {
			return strconv.ParseFloat(val, 64)
		}
		return 0, nil // Return zero but no error if key is absent
	}

	parseInt := func(key string) (int, error) {
		if val := query.Get(key); val != "" {
			return strconv.Atoi(val)
		}
		return 0, nil // Return zero but no error if key is absent
	}

	// Update fields if the corresponding query parameter is present
	if salary, err := parseFloat("salary"); err == nil && salary != 0 {
		params.Salary = salary
	}

	if initialInvestmentAmount, err := parseFloat("initial_investment"); err == nil &&
		initialInvestmentAmount != 0 {
		params.InitialInvestment = initialInvestmentAmount
	}

	if annualInvestmentAmount, err := parseFloat("annual_investment"); err == nil &&
		annualInvestmentAmount != 0 {
		params.AnnualInvestment = annualInvestmentAmount
	}

	if mortgageSize, err := parseFloat("mortgage_size"); err == nil && mortgageSize != 0 {
		params.MortgageSize = mortgageSize
	}

	if mortgageInterestRate, err := parseFloat("mortgage_interest_rate"); err == nil &&
		mortgageInterestRate != 0 {
		params.MortgageInterestRate = mortgageInterestRate / 100
	}

	if dividendReturnRate, err := parseFloat("dividend_return_rate"); err == nil &&
		dividendReturnRate != 0 {
		params.DividendReturnRate = dividendReturnRate / 100
	}

	if capitalGrowthRate, err := parseFloat("capital_growth_rate"); err == nil &&
		capitalGrowthRate != 0 {
		params.CapitalGrowthRate = capitalGrowthRate / 100
	}

	if years, err := parseInt("years"); err == nil && years != 0 {
		params.NumYears = years
	}

	// Simple string params
	if country := query.Get("country"); country != "" {
		params.Country = country
	}

	// Checkboxes for reinvestment toggles: maintain defaults if not present
	if query["reinvest_dividends"] != nil {
		params.ReinvestDividends = query.Get("reinvest_dividends") == "true"
	}
	if query["reinvest_tax_refunds"] != nil {
		params.ReinvestTaxRefunds = query.Get("reinvest_tax_refunds") == "true"
	}
}
