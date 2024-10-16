package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"debtrecyclingcalc.com/internal/calc"
)

func getFormParams(r *http.Request) (*calc.Parameters, error) {
	parseFloat := func(key string) (float64, error) {
		return strconv.ParseFloat(r.Form.Get(key), 64)
	}

	parseInt := func(key string) (int, error) {
		return strconv.Atoi(r.Form.Get(key))
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

	country := r.Form.Get("country")
	reinvestDividends := r.Form.Get("reinvest_dividends") == "on"
	reinvestTaxRefunds := r.Form.Get("reinvest_tax_refunds") == "on"

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

func getQueryParams(query url.Values) (*calc.Parameters, error) {
	parseFloat := func(key string) (float64, error) {
		return strconv.ParseFloat(query.Get(key), 64)
	}

	parseInt := func(key string) (int, error) {
		return strconv.Atoi(query.Get(key))
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

	country := query.Get("country")
	reinvestDividends := query.Get("reinvest_dividends") == "true"
	reinvestTaxRefunds := query.Get("reinvest_tax_refunds") == "true"

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
