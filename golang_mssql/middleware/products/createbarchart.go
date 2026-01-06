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
			Subtitle: "Revenue Over Time",
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
