package calc

import "math"

type DebtRecyclingParameters struct {
	Salary               float64
	InitialInvestment    float64
	AnnualInvestment     float64
	MortgageSize         float64
	MortgageInterestRate float64
	DividendReturnRate   float64
	CapitalGrowthRate    float64
	ReinvestDividends    bool
	ReinvestTaxRefunds   bool
	NumYears             int
}

type DebtRecyclingData struct {
	CumulativeDebtRecycled []float64
	DividendReturn         []float64
	CumulativeDividends    []float64
	NonDeductibleInterest  []float64
	TaxDeductibleInterest  []float64
	TaxSavings             []float64
	CumulativeTaxSavings   []float64
	PortfolioValue         []float64
	NetPosition            []float64
	TotalInvested          float64
	TotalValue             float64
}

func taxRate(salary float64) float64 {
	taxBrackets := []struct {
		lowerBound   float64
		upperBound   float64
		marginalRate float64
	}{
		{0, 18200, 0},
		{18201, 45000, 0.16},
		{45001, 135000, 0.30},
		{135001, 190000, 0.37},
		{190001, math.MaxFloat64, 0.45},
	}

	taxLiability := 0.0
	taxableIncome := salary
	for _, bracket := range taxBrackets {
		if taxableIncome > bracket.lowerBound {
			incomeInBracket := math.Min(taxableIncome, bracket.upperBound) - bracket.lowerBound
			taxLiability += incomeInBracket * bracket.marginalRate
			taxableIncome -= incomeInBracket
		}
		if taxableIncome <= 18200 {
			break
		}
	}

	return taxLiability / salary
}

func CAGR(initialValue, finalValue float64, numYears int) float64 {
	if initialValue <= 0 || numYears <= 0 {
		return 0
	}
	return math.Pow(finalValue/initialValue, 1/float64(numYears)) - 1
}

func DebtRecycling(params DebtRecyclingParameters) *DebtRecyclingData {
	data := &DebtRecyclingData{}

	// Pre-allocate slices with the correct size
	data.CumulativeDebtRecycled = make([]float64, params.NumYears)
	data.CumulativeDividends = make([]float64, params.NumYears)
	data.DividendReturn = make([]float64, params.NumYears)
	data.NonDeductibleInterest = make([]float64, params.NumYears)
	data.TaxDeductibleInterest = make([]float64, params.NumYears)
	data.TaxSavings = make([]float64, params.NumYears)
	data.CumulativeTaxSavings = make([]float64, params.NumYears)
	data.PortfolioValue = make([]float64, params.NumYears+1)
	data.NetPosition = make([]float64, params.NumYears)

	// Initialize portfolio value
	data.PortfolioValue[0] = params.InitialInvestment
	data.NetPosition[0] = params.InitialInvestment - (params.MortgageSize - params.InitialInvestment)

	// Calculate for each year
	for year := 0; year < params.NumYears; year++ {
		// Calculate portfolio value for the next year
		data.PortfolioValue[year+1] = (data.PortfolioValue[year] + params.AnnualInvestment) * (1 + params.CapitalGrowthRate)

		// Calculate dividends for the year
		data.DividendReturn[year] = data.PortfolioValue[year] * params.DividendReturnRate

		// Accumulate dividends
		if year > 0 {
			data.CumulativeDividends[year] = data.CumulativeDividends[year-1] + data.DividendReturn[year]
		} else {
			data.CumulativeDividends[year] = data.DividendReturn[year]
		}

		// Calculate total invested amount up to the current year (not gt mortgage size)
		data.CumulativeDebtRecycled[year] = math.Min(
			params.InitialInvestment+params.AnnualInvestment*float64(year),
			params.MortgageSize,
		)

		// Non-deductible interest is the interest charged on the non-recycled debt
		data.NonDeductibleInterest[year] = math.Max(
			(params.MortgageSize-data.CumulativeDebtRecycled[year])*params.MortgageInterestRate,
			0,
		)

		// Calculate Non-Deductible and Tax-Deductible Interests
		data.NonDeductibleInterest[year] = (params.MortgageSize - data.CumulativeDebtRecycled[year]) * params.MortgageInterestRate
		data.TaxDeductibleInterest[year] = data.CumulativeDebtRecycled[year] * params.MortgageInterestRate

		// Calculate tax savings (adjusting for tax liability)
		taxRate := taxRate(params.Salary)
		data.TaxSavings[year] = data.TaxDeductibleInterest[year] * (1 - taxRate)

		// Accumulate tax savings
		if year > 0 {
			data.CumulativeTaxSavings[year] = data.CumulativeTaxSavings[year-1] + data.TaxSavings[year]
		} else {
			data.CumulativeTaxSavings[year] = data.TaxSavings[year]
		}

		// Current dividend and tax savings balance
		dividendBalance := data.CumulativeDividends[year]
		taxRefundBalance := data.CumulativeTaxSavings[year]

		// Reinvest dividends and tax refunds if applicable
		if params.ReinvestDividends {
			data.PortfolioValue[year+1] += data.DividendReturn[year]
			dividendBalance -= data.DividendReturn[year]
		}
		if params.ReinvestTaxRefunds {
			data.PortfolioValue[year+1] += data.TaxSavings[year]
			taxRefundBalance -= data.TaxSavings[year]
		}

		// Calculate Net Position considering debt repayment and avoiding double-counting
		data.NetPosition[year] = data.PortfolioValue[year+1] + dividendBalance - params.MortgageSize

		// Calculate the total invested amount
		data.TotalInvested = params.AnnualInvestment * float64(year+1)

		// Calculate Total Returns
		data.TotalValue += data.PortfolioValue[params.NumYears] + dividendBalance + taxRefundBalance

	}

	return data
}
