package charts

import (
	"context"
	"fmt"
	"io"

	"debtrecyclingcalculator.com.au/internal/calc"
	"github.com/a-h/templ"

	echarts "github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Renderable interface {
	Render(w io.Writer) error
}

func ChartToTemplComponent(chart Renderable) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return chart.Render(w)
	})
}

func StackedLineChart(data *calc.DebtRecyclingData) (*echarts.Line, error) {
	// Create a new line chart
	line := echarts.NewLine()

	// Set the chart title
	line.SetGlobalOptions(
		echarts.WithLegendOpts(opts.Legend{Left: "center"}),
	)

	years := len(data.CumulativeDebtRecycled)
	debtRecycledLineData := make([]opts.LineData, years)
	netPositionLineData := make([]opts.LineData, years)
	portfolioValueLineData := make([]opts.LineData, years)
	xAxisData := make([]string, len(data.CumulativeDebtRecycled))

	// Create LineData for all series in loop
	for year := 0; year < years; year++ {
		debtRecycledLineData[year] = opts.LineData{
			Value: int(data.CumulativeDebtRecycled[year]),
		}

		netPositionLineData[year] = opts.LineData{
			Value: int(data.NetPosition[year]),
		}

		portfolioValueLineData[year] = opts.LineData{
			Value: int(data.PortfolioValue[year]),
		}
		xAxisData[year] = fmt.Sprintf("%d", year+1)
	}

	line.SetXAxis(xAxisData).
		AddSeries(
			"Debt Recycled",
			debtRecycledLineData,
			echarts.WithLineChartOpts(
				opts.LineChart{
					Stack: "debt_recycled",
				},
			),
		).
		AddSeries(
			"Net Position",
			netPositionLineData,
			echarts.WithLineChartOpts(
				opts.LineChart{
					Stack: "net_position",
				},
			),
		).
		AddSeries("Portfolio Value",
			portfolioValueLineData,
			echarts.WithLineChartOpts(
				opts.LineChart{
					Stack: "portfolio_value",
				},
			),
		).
		SetSeriesOptions(
			echarts.WithAreaStyleOpts(
				opts.AreaStyle{
					Opacity: 0.5,
				},
			),
		)

	// Render the chart to the writer
	return line, nil
}
