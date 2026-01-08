package middleware

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func CreateBarChart(values []float64, labels []string) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Apple Inc.",
			Subtitle: "Revenue Over Time for the year 2025",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Show:      opts.Bool(true),
				Formatter: opts.FuncOpts("function (value) { var date = new Date(value); return date.toLocaleString('default', { month: 'short' }); }"),
			},
		}),
	)
	bar.SetXAxis(labels).AddSeries("Sales", generateBarItems(values))
	return bar
}

func generateBarItems(data []float64) []opts.BarData {
	items := make([]opts.BarData, 0)
	for _, v := range data {
		items = append(items, opts.BarData{Value: v})
	}
	return items
}
