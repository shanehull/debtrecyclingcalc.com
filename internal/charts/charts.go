package charts

import (
	"context"
	"fmt"
	"html/template"
	"io"

	"debtrecyclingcalc.com/internal/calc"
	"debtrecyclingcalc.com/internal/middleware"
	"github.com/a-h/templ"

	echarts "github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/render"
)

type Renderable interface {
	Render(w io.Writer) error
}

type Renderer interface {
	Render(w io.Writer) error
	RenderContent() []byte
	RenderSnippet() render.ChartSnippet
}

var baseTpl = `
<div id="{{ .ChartID }}" class="echarts-chart-wrapper"></div>
<script type="text/javascript" nonce="{{ getScriptNonce }}">
    "use strict";
    let goecharts_{{ .ChartID | safeJS }} = echarts.init(
        document.getElementById('{{ .ChartID | safeJS }}'),
            "{{ .Theme }}",
            { 
                renderer: "{{ .Initialization.Renderer }}", 
            },
    );
    let option_{{ .ChartID | safeJS }} = {{ .JSON }};
    goecharts_{{ .ChartID | safeJS }}.setOption(option_{{ .ChartID | safeJS }});
    {{- range .JSFunctions.Fns }}
    {{ . | safeJS }}
    {{- end }}
</script>
`

type tailwindRenderer struct {
	render.BaseRender
	styleNonce  string
	scriptNonce string
	c           interface{}
	before      []func()
}

func TailwindRender(
	styleNonce string,
	scriptNonce string,
	c interface{},
	before ...func(),
) render.Renderer {
	return &tailwindRenderer{styleNonce: styleNonce, scriptNonce: scriptNonce, c: c, before: before}
}

func (r *tailwindRenderer) Render(w io.Writer) error {
	const tplName = "chart"
	for _, fn := range r.before {
		fn()
	}

	tpl := template.
		Must(template.New(tplName).
			Funcs(template.FuncMap{
				"safeJS": func(s interface{}) template.JS {
					return template.JS(fmt.Sprint(s))
				},
				"getStyleNonce": func() template.HTMLAttr {
					return template.HTMLAttr(r.styleNonce)
				},
				"getScriptNonce": func() template.HTMLAttr {
					return template.HTMLAttr(r.scriptNonce)
				},
			}).
			Parse(baseTpl),
		)

	err := tpl.ExecuteTemplate(w, tplName, r.c)
	return err
}

func ChartToTemplComponent(chart Renderable) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		return chart.Render(w)
	})
}

func Positions(
	data *calc.Data,
	years int,
	ctx context.Context,
) (*echarts.Line, error) {
	line := echarts.NewLine()

	styleNonce := middleware.GetInlineStyleNonce(ctx)
	scriptNonce := middleware.GetInlineScriptNonce(ctx)
	line.Renderer = TailwindRender(styleNonce, scriptNonce, line, line.Validate)

	line.SetGlobalOptions(
		echarts.WithTitleOpts(
			opts.Title{
				Subtitle: "Chart in 000's",
				Bottom:   "3%",
				Left:     "center",
			},
		),
		echarts.WithLegendOpts(opts.Legend{Left: "center"}),
		echarts.WithTooltipOpts(
			opts.Tooltip{
				// TODO: tooltip injects style attributes and triggers csp violation
				// https://github.com/apache/echarts/issues/19938
				Show:        opts.Bool(false),
				Trigger:     "axis",
				AxisPointer: &opts.AxisPointer{Type: "cross"},
			},
		),
		echarts.WithColorsOpts(opts.Colors{"green", "blue", "orange"}),
		echarts.WithInitializationOpts(opts.Initialization{Renderer: "svg"}),
	)

	portfolioValueLineData := make([]opts.LineData, years)
	debtRecycledLineData := make([]opts.LineData, years)
	netPositionLineData := make([]opts.LineData, years)
	xAxisData := make([]string, years)

	for year := 0; year < years; year++ {
		portfolioValueLineData[year] = opts.LineData{
			Value: int(data.PortfolioValue[year] / 1000),
		}

		debtRecycledLineData[year] = opts.LineData{
			Value: int(data.DebtRecycled[year] / 1000),
		}

		netPositionLineData[year] = opts.LineData{
			Value: int(data.NetPosition[year] / 1000),
		}
		xAxisData[year] = fmt.Sprintf("%d", year+1)
	}

	line.SetXAxis(xAxisData).
		AddSeries("Portfolio Value",
			portfolioValueLineData,
			echarts.WithLineChartOpts(
				opts.LineChart{
					Stack: "portfolio_value",
				},
			),
		).
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
		SetSeriesOptions(
			echarts.WithAreaStyleOpts(
				opts.AreaStyle{
					Opacity: 0.5,
				},
			),
		)

	return line, nil
}

func Income(data *calc.Data, years int, ctx context.Context) (*echarts.Bar, error) {
	bar := echarts.NewBar()

	styleNonce := middleware.GetInlineStyleNonce(ctx)
	scriptNonce := middleware.GetInlineScriptNonce(ctx)
	bar.Renderer = TailwindRender(styleNonce, scriptNonce, bar, bar.Validate)

	bar.SetGlobalOptions(
		echarts.WithTitleOpts(
			opts.Title{
				Subtitle: "Chart in 000's",
				Bottom:   "3%",
				Left:     "center",
			},
		),
		echarts.WithLegendOpts(opts.Legend{Left: "center"}),
		echarts.WithTooltipOpts(
			opts.Tooltip{
				// TODO: tooltip injects style attributes and triggers csp violation
				// https://github.com/apache/echarts/issues/19938
				Show:        opts.Bool(false),
				Trigger:     "axis",
				AxisPointer: &opts.AxisPointer{Type: "cross"},
			},
		),
		echarts.WithColorsOpts(opts.Colors{"grey", "pink"}),
		echarts.WithInitializationOpts(opts.Initialization{Renderer: "svg"}),
	)

	dividendsLineData := make([]opts.BarData, years)
	taxRefundsLineData := make([]opts.BarData, years)
	xAxisData := make([]string, years)

	for year := 0; year < years; year++ {
		dividendsLineData[year] = opts.BarData{
			Value: int(data.DividendReturns[year] / 1000),
		}

		taxRefundsLineData[year] = opts.BarData{
			Value: int(data.TaxRefunds[year] / 1000),
		}

		xAxisData[year] = fmt.Sprintf("%d", year+1)
	}

	bar.SetXAxis(xAxisData).
		AddSeries(
			"Tax Refunds",
			taxRefundsLineData,
			echarts.WithLineChartOpts(
				opts.LineChart{
					Stack: "income",
				},
			),
		).
		AddSeries(
			"Dividends",
			dividendsLineData,
			echarts.WithLineChartOpts(
				opts.LineChart{
					Stack: "income",
				},
			),
		).
		SetSeriesOptions(
			echarts.WithAreaStyleOpts(
				opts.AreaStyle{
					Opacity: 0.6,
				},
			),
		)

	return bar, nil
}

func Interest(data *calc.Data, years int, ctx context.Context) (*echarts.Bar, error) {
	bar := echarts.NewBar()

	styleNonce := middleware.GetInlineStyleNonce(ctx)
	scriptNonce := middleware.GetInlineScriptNonce(ctx)
	bar.Renderer = TailwindRender(styleNonce, scriptNonce, bar, bar.Validate)

	bar.SetGlobalOptions(
		echarts.WithTitleOpts(
			opts.Title{
				Subtitle: "Chart in 000's",
				Bottom:   "3%",
				Left:     "center",
			},
		),
		echarts.WithLegendOpts(opts.Legend{Left: "center"}),
		echarts.WithTooltipOpts(
			opts.Tooltip{
				// TODO: tooltip injects style attributes and triggers csp violation
				// https://github.com/apache/echarts/issues/19938
				Show:        opts.Bool(false),
				Trigger:     "axis",
				AxisPointer: &opts.AxisPointer{Type: "cross"},
			},
		),
		echarts.WithColorsOpts(opts.Colors{"blue", "grey"}),
		echarts.WithInitializationOpts(opts.Initialization{Renderer: "svg"}),
	)

	nonDeductibleLineData := make([]opts.BarData, years)
	taxdeductibleLineData := make([]opts.BarData, years)
	xAxisData := make([]string, years)

	for year := 0; year < years; year++ {
		nonDeductibleLineData[year] = opts.BarData{
			Value: int(data.NonDeductibleInterest[year] / 1000),
		}

		taxdeductibleLineData[year] = opts.BarData{
			Value: int(data.TaxDeductibleInterest[year] / 1000),
		}

		xAxisData[year] = fmt.Sprintf("%d", year+1)
	}

	bar.SetXAxis(xAxisData).
		AddSeries(
			"Tax-Deductible",
			taxdeductibleLineData,
			echarts.WithLineChartOpts(
				opts.LineChart{
					Stack: "interest",
				},
			),
		).
		AddSeries(
			"Non-Deductible",
			nonDeductibleLineData,
			echarts.WithLineChartOpts(
				opts.LineChart{
					Stack: "interest",
				},
			),
		).
		SetSeriesOptions(
			echarts.WithAreaStyleOpts(
				opts.AreaStyle{
					Opacity: 0.6,
				},
			),
		)

	return bar, nil
}
