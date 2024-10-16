package calc

import (
	"errors"
	"math"
	"strings"
)

type Parameters struct {
	Salary               float64
	InitialInvestment    float64
	AnnualInvestment     float64
	MortgageSize         float64
	MortgageInterestRate float64
	DividendReturnRate   float64
	CapitalGrowthRate    float64
	NumYears             int
	Country              string
	ReinvestDividends    bool
	ReinvestTaxRefunds   bool
}

type Data struct {
	DebtRecycled          []float64
	NonDeductibleInterest []float64
	TaxDeductibleInterest []float64
	TaxRefunds            []float64
	CumulativeTaxRefunds  []float64
	DividendReturns       []float64
	CumulativeDividends   []float64
	PortfolioValue        []float64
	NetPosition           []float64
	TotalInvested         float64
	TotalValue            float64
}

type TaxBracket struct {
	LowerBound   float64
	UpperBound   float64
	MarginalRate float64
}

type CountryBrackets []TaxBracket

var countryBrackets = map[string]CountryBrackets{
	"au": {
		TaxBracket{
			LowerBound:   0,
			UpperBound:   18200,
			MarginalRate: 0,
		},
		TaxBracket{
			LowerBound:   18201,
			UpperBound:   45000,
			MarginalRate: 0.16,
		},
		TaxBracket{
			LowerBound:   45001,
			UpperBound:   135000,
			MarginalRate: 0.30,
		},
		TaxBracket{
			LowerBound:   135001,
			UpperBound:   190000,
			MarginalRate: 0.37,
		},
		TaxBracket{
			LowerBound:   190001,
			UpperBound:   math.MaxFloat64,
			MarginalRate: .45,
		},
	},
	"nz": {
		TaxBracket{
			LowerBound:   0,
			UpperBound:   14000,
			MarginalRate: 0.105,
		},
		TaxBracket{
			LowerBound:   14001,
			UpperBound:   48000,
			MarginalRate: 0.175,
		},
		TaxBracket{
			LowerBound:   48001,
			UpperBound:   70000,
			MarginalRate: 0.30,
		},
		TaxBracket{
			LowerBound:   70001,
			UpperBound:   180000,
			MarginalRate: 0.33,
		},
		TaxBracket{
			LowerBound:   180001,
			UpperBound:   math.MaxFloat64,
			MarginalRate: 0.39,
		},
	},
}

func taxRate(salary float64, country string) (float64, error) {
	taxBrackets := countryBrackets[strings.ToLower(country)]
	if taxBrackets == nil {
		return 0, errors.New("unknown country")
	}

	for i := len(taxBrackets) - 1; i >= 0; i-- {
		bracket := taxBrackets[i]
		if salary > bracket.LowerBound {
			return bracket.MarginalRate, nil
		}
	}

	return 0, nil
}

func CAGR(initialValue, finalValue float64, numYears int) float64 {
	if initialValue <= 0 || numYears <= 0 {
		return 0
	}
	return math.Pow(finalValue/initialValue, 1/float64(numYears)) - 1
}

func DebtRecycling(params Parameters) (*Data, error) {
	data := &Data{}

	// Pre-allocate slices with the correct size
	data.DebtRecycled = make([]float64, params.NumYears)
	data.NonDeductibleInterest = make([]float64, params.NumYears)
	data.TaxDeductibleInterest = make([]float64, params.NumYears)
	data.TaxRefunds = make([]float64, params.NumYears)
	data.CumulativeTaxRefunds = make([]float64, params.NumYears)
	data.DividendReturns = make([]float64, params.NumYears)
	data.CumulativeDividends = make([]float64, params.NumYears)
	data.PortfolioValue = make([]float64, params.NumYears+1)
	data.NetPosition = make([]float64, params.NumYears)

	data.PortfolioValue[0] = params.InitialInvestment
	data.NetPosition[0] = params.MortgageSize - params.InitialInvestment

	// Calculate for each year
	for year := 0; year < params.NumYears; year++ {
		// Calculate dividends for the year
		data.DividendReturns[year] = data.PortfolioValue[year] * params.DividendReturnRate

		// Accumulate dividends
		if year > 0 {
			data.CumulativeDividends[year] = data.CumulativeDividends[year-1] + data.DividendReturns[year]
		} else {
			data.CumulativeDividends[year] = data.DividendReturns[year]
		}

		// Calculate total invested amount up to the current year (not gt mortgage size)
		data.DebtRecycled[year] = math.Min(
			params.InitialInvestment+params.AnnualInvestment*float64(year),
			params.MortgageSize,
		)

		// Calculate Non-Deductible and Tax-Deductible Interests
		data.NonDeductibleInterest[year] = math.Max(
			(params.MortgageSize-data.DebtRecycled[year])*params.MortgageInterestRate,
			0,
		)
		data.TaxDeductibleInterest[year] = math.Min(
			data.DebtRecycled[year]*params.MortgageInterestRate,
			params.MortgageSize*params.MortgageInterestRate,
		)

		// Calculate tax savings (adjusting for tax liability)
		taxRate, err := taxRate((params.Salary + data.DividendReturns[year]), params.Country)
		if err != nil {
			return nil, err
		}
		data.TaxRefunds[year] = data.TaxDeductibleInterest[year] * (1 - taxRate)

		// Accumulate tax savings
		if year > 0 {
			data.CumulativeTaxRefunds[year] = data.CumulativeTaxRefunds[year-1] + data.TaxRefunds[year]
		} else {
			data.CumulativeTaxRefunds[year] = data.TaxRefunds[year]
		}

		// Reinvest dividends and tax refunds if applicable
		reinvestment := 0.0
		if params.ReinvestDividends {
			reinvestment += data.DividendReturns[year]
		}
		if params.ReinvestTaxRefunds {
			reinvestment += data.TaxRefunds[year]
		}

		// Apply annual growth and investments to the following year
		data.PortfolioValue[year+1] = (data.PortfolioValue[year] + params.AnnualInvestment + reinvestment) * (1 + params.CapitalGrowthRate)

		// Update the cumulative debt recycled, ensuring it does not exceed mortgage size
		data.DebtRecycled[year] = math.Min(
			params.InitialInvestment+params.AnnualInvestment*float64(year)+reinvestment,
			params.MortgageSize,
		)

		// Recalculate Non-Deductible and Tax-Deductible Interests after reinvestments
		data.NonDeductibleInterest[year] = math.Max(
			(params.MortgageSize-data.DebtRecycled[year])*params.MortgageInterestRate,
			0,
		)
		data.TaxDeductibleInterest[year] = math.Min(
			data.DebtRecycled[year]*params.MortgageInterestRate,
			params.MortgageSize*params.MortgageInterestRate,
		)

		// Calculate net position
		data.NetPosition[year] = data.PortfolioValue[year] - params.MortgageSize
	}

	// Final values
	data.TotalValue = data.PortfolioValue[params.NumYears] + data.CumulativeTaxRefunds[params.NumYears-1] + data.CumulativeDividends[params.NumYears-1]
	data.TotalInvested = params.InitialInvestment + (params.AnnualInvestment * float64(params.NumYears))

	return data, nil
}
