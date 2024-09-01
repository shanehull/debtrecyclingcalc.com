package calc

import (
	"math"
	"testing"
)

func Test_CAGR(t *testing.T) {
	cagr := CAGR(100, 200, 10)
	if math.Abs(cagr-0.071773) > 0.000001 {
		t.Fatalf(
			"CAGR(%d, %d, %d) = %f, want %f",
			100, 200, 10, cagr, 0.071773,
		)
	}
}

func Test_DebtRecycling(t *testing.T) {
	type drTests struct {
		params   DebtRecyclingParameters
		expected *DebtRecyclingData
	}

	cases := []drTests{
		{
			params: DebtRecyclingParameters{
				Salary:               150000,
				InitialInvestment:    100000,
				AnnualInvestment:     50000,
				MortgageSize:         600000,
				MortgageInterestRate: 0.05,
				DividendReturnRate:   0.02,
				CapitalGrowthRate:    0.08,
				NumYears:             10,
				Country:              "au",
				ReinvestDividends:    false,
				ReinvestTaxRefunds:   false,
			},
			expected: &DebtRecyclingData{
				DebtRecycled: []float64{
					100000,
					150000,
					200000,
					250000,
					300000,
					350000,
					400000,
					450000,
					500000,
					550000,
				},
				DividendReturns: []float64{
					2000,
					3240,
					4579.200000000001,
					6025.536000000001,
					7587.578880000002,
					9274.585190400003,
					11096.552005632004,
					13064.276166082565,
					15189.41825936917,
					17484.571720118704,
				},
				CumulativeDividends: []float64{
					2000,
					5240,
					9819.2,
					15844.736,
					23432.31488,
					32706.900070400006,
					43803.45207603201,
					56867.72824211457,
					72057.14650148374,
					89541.71822160245,
				},
				NonDeductibleInterest: []float64{
					25000,
					22500,
					20000,
					17500,
					15000,
					12500,
					10000,
					7500,
					5000,
					2500,
				},
				TaxDeductibleInterest: []float64{
					5000,
					7500,
					10000,
					12500,
					15000,
					17500,
					20000,
					22500,
					25000,
					27500,
				},
				TaxRefunds: []float64{
					4247.11052631579,
					6361.59749412686,
					8469.289529251026,
					10569.58676302833,
					12661.860740683269,
					14745.456506917692,
					16819.69535874436,
					18883.878320960488,
					20937.290390528968,
					22979.20558344762,
				},
				CumulativeTaxRefunds: []float64{
					4247.11052631579,
					10608.70802044265,
					19077.997549693675,
					29647.584312722007,
					42309.445053405274,
					57054.90156032296,
					73874.59691906732,
					92758.47524002781,
					113695.76563055678,
					136674.97121400438,
				},
				PortfolioValue: []float64{
					100000,
					162000,
					228960.00000000003,
					301276.80000000005,
					379378.9440000001,
					463729.25952000014,
					554827.6002816001,
					653213.8083041282,
					759470.9129684585,
					874228.5860059352,
					998166.8728864101,
				},
				NetPosition: []float64{
					-498000,
					-432760,
					-361220.79999999993,
					-282878.464,
					-197188.7411199999,
					-103563.84040959983,
					-1368.9476423677988,
					110081.53654624277,
					231528.05946994224,
					363770.3042275377,
				},
				TotalInvested: 600000,
				TotalValue:    1.9294296249046014e+06,
			},
		},
		{
			params: DebtRecyclingParameters{
				Salary:               150000,
				InitialInvestment:    100000,
				AnnualInvestment:     50000,
				MortgageSize:         600000,
				MortgageInterestRate: 0.05,
				DividendReturnRate:   0.02,
				CapitalGrowthRate:    0.08,
				NumYears:             10,
				Country:              "au",
				ReinvestDividends:    true,
				ReinvestTaxRefunds:   false,
			},
			expected: &DebtRecyclingData{
				PortfolioValue: []float64{
					100000,
					164000,
					234400.00000000003,
					311840,
					397024,
					490726.4,
					593799.0400000002,
					707178.9440000003,
					831896.8384000004,
					969086.5222400004,
					1.1199951744640006e+06,
				},
				CumulativeDividends: []float64{
					2000,
					5280,
					9968,
					16204.8,
					24145.28,
					33959.808,
					45835.7888,
					59979.36768000001,
					76617.30444800001,
					95999.03489280002,
				},
				DebtRecycled: []float64{
					102000,
					153280,
					204688,
					256236.8,
					307940.48,
					359814.528,
					411875.9808,
					464143.57888,
					516637.936768,
					569381.7304448,
				},
				CumulativeTaxRefunds: []float64{
					4247.11052631579,
					10608.417937589276,
					19076.674040195816,
					29643.800351026555,
					42300.8306449702,
					57037.85610488333,
					73843.97470107718,
					92707.24665168812,
					113614.65801918795,
					136552.09466515813,
				},
				NetPosition: []float64{
					-500000,
					-434000,
					-360320,
					-278192,
					-186771.2,
					-85128.31999999995,
					27758.848000000115,
					153014.73280000023,
					291876.2060800004,
					445703.82668800047,
				},
				TotalInvested: 600000,
				TotalValue:    1.9736181870340928e+06,
			},
		},
		{
			params: DebtRecyclingParameters{
				Salary:               150000,
				InitialInvestment:    100000,
				AnnualInvestment:     50000,
				MortgageSize:         600000,
				MortgageInterestRate: 0.05,
				DividendReturnRate:   0.02,
				CapitalGrowthRate:    0.08,
				NumYears:             10,
				Country:              "au",
				ReinvestDividends:    false,
				ReinvestTaxRefunds:   true,
			},
			expected: &DebtRecyclingData{
				PortfolioValue: []float64{
					100000,
					166247.11052631578,
					239907.86103614027,
					321567.7011835439,
					411857.9836250151,
					511459.60838165897,
					621106.9662296256,
					741592.2063071347,
					873769.8551086588,
					1.0185618161322233e+06,
					1.176962781658853e+06,
				},
				DebtRecycled: []float64{
					104247.1105263158,
					156360.9816677192,
					208467.21126451236,
					260564.86634678775,
					312652.9860666426,
					364730.58917743387,
					416796.68277913897,
					468850.27229695325,
					520890.3726148719,
					572916.0202360519,
				},
				CumulativeDividends: []float64{
					2000,
					5324.942210526316,
					10123.09943124912,
					16554.453454919996,
					24791.6131274203,
					35020.80529505348,
					47442.94461964599,
					62274.78874578868,
					79750.18584796187,
					100121.42217060634,
				},
				CumulativeTaxRefunds: []float64{
					4247.11052631579,
					10608.092194034987,
					19075.30345854735,
					29640.169805335096,
					42293.15587197774,
					57023.74504941159,
					73820.42782855057,
					92670.70012550385,
					113561.07274037572,
					136477.09297642764,
				},
				NetPosition: []float64{
					-498000,
					-428427.9472631579,
					-349969.0395326106,
					-261877.84536153614,
					-163350.40324756457,
					-53519.58632328757,
					68549.91084927157,
					203866.99505292345,
					353520.0409566206,
					518683.2383028297,
				},
				TotalInvested: 600000,
				TotalValue:    2.0033068141620778e+06,
			},
		},
		// TODO: continue test cases
		// {
		// 	params: DebtRecyclingParameters{
		// 		Salary:               150000,
		// 		InitialInvestment:    100000,
		// 		AnnualInvestment:     50000,
		// 		MortgageSize:         600000,
		// 		MortgageInterestRate: 0.05,
		// 		DividendReturnRate:   0.02,
		// 		CapitalGrowthRate:    0.08,
		// 		NumYears:             10,
		// 		Country:              "au",
		// 		ReinvestDividends:    true,
		// 		ReinvestTaxRefunds:   true,
		// 	},
		// 	expected: &DebtRecyclingData{

		// blah
		//},
		// },
		// {
		// 	params: DebtRecyclingParameters{
		// 		Salary:               150000,
		// 		InitialInvestment:    100000,
		// 		AnnualInvestment:     50000,
		// 		MortgageSize:         600000,
		// 		MortgageInterestRate: 0.05,
		// 		DividendReturnRate:   0.02,
		// 		CapitalGrowthRate:    0.08,
		// 		NumYears:             10,
		// 		Country:              "nz",
		// 		ReinvestDividends:    false,
		// 		ReinvestTaxRefunds:   false,
		// 	},
		// 	expected: &DebtRecyclingData{},
		// },
	}
	for i, c := range cases {
		dr := DebtRecycling(c.params)

		// Compare each float64 field
		compareAllFloat64Values(
			t,
			dr.PortfolioValue,
			c.expected.PortfolioValue,
			"PortfolioValue",
			i,
			c.params,
		)
		compareAllFloat64Values(
			t,
			dr.DebtRecycled,
			c.expected.DebtRecycled,
			"DebtRecycled",
			i,
			c.params,
		)
		compareAllFloat64Values(
			t,
			dr.CumulativeDividends,
			c.expected.CumulativeDividends,
			"CumulativeDividends",
			i,
			c.params,
		)
		compareAllFloat64Values(
			t,
			dr.CumulativeTaxRefunds,
			c.expected.CumulativeTaxRefunds,
			"CumulativeTaxRefunds",
			i,
			c.params,
		)
		compareAllFloat64Values(
			t,
			dr.NetPosition,
			c.expected.NetPosition,
			"NetPosition",
			i,
			c.params,
		)

		if dr.TotalInvested != c.expected.TotalInvested {
			t.Errorf(
				"Test %d: Parameters %v - TotalInvested is incorrect. Expected %v, got %v",
				i,
				c.params,
				c.expected.TotalInvested,
				dr.TotalInvested,
			)
		}

		if dr.TotalValue != c.expected.TotalValue {
			t.Errorf(
				"Test %d: Parameters %v - TotalValue is incorrect. Expected %v, got %v",
				i,
				c.params,
				c.expected.TotalValue,
				dr.TotalValue,
			)
		}
	}
}

// Helper function to compare all values in slices of float64
func compareAllFloat64Values(
	t *testing.T,
	got, want []float64,
	fieldName string,
	testIndex int,
	params DebtRecyclingParameters,
) {
	if len(got) != len(want) {
		t.Errorf(
			"Test %d: Parameters %v - %s length is incorrect. Expected length %d, got length %d. Expected: %v, Got: %v",
			testIndex,
			params,
			fieldName,
			len(want),
			len(got),
			want,
			got,
		)
		return
	}

	for i := range got {
		if math.Abs(got[i]-want[i]) > 0.000001 {
			t.Errorf(
				"Test %d: Parameters %v - %s[%d] is incorrect. Expected %v, got %v. Full expected: %v, Full got: %v",
				testIndex,
				params,
				fieldName,
				i,
				want[i],
				got[i],
				want,
				got,
			)
		}
	}
}
