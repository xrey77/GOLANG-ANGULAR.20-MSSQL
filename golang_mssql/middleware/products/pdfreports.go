package middleware

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"log"
	"os"

	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"net/http"
	"src/golang_mssql/dbconfig"
	"src/golang_mssql/models"
	"time"
)

func ProductPDFReport(c *gin.Context) {

	cfg := config.NewBuilder().
		WithPageNumber(props.PageNumber{
			Pattern: "Page {current} of {total}",
			Place:   props.Bottom,
		}).
		WithLeftMargin(20).
		WithRightMargin(20).
		// WithTopMargin(15).
		// WithBottomMargin(15).
		Build()

	m := maroto.New(cfg)

	err := m.RegisterFooter(
		text.NewRow(10, "Prepared by: Rey Gragasin", props.Text{
			Size:  6,
			Align: align.Left,
			Top:   5,
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	db := dbconfig.Connection()

	// 1. Fetch data from MSSQL (Assuming 'db' is your *gorm.DB instance)
	var products []models.Product
	if err := db.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error: " + err.Error()})
		return
	}

	now := time.Now()
	imgBytes, err := os.ReadFile("assets/images/logo.png")
	if err != nil {
		log.Fatalf("failed to read logo: %v", err)
	}

	m.AddRows(
		row.New(20).Add(
			image.NewFromBytesCol(14, imgBytes, extension.Png, props.Rect{
				Center:  true,
				Percent: 80,
			}),
		),
		row.New(15).Add(
			text.NewCol(12, "Product List Report", props.Text{
				Top:   5,
				Size:  20,
				Style: fontstyle.Bold,
				Align: align.Center,
			}),
		),
		row.New(18).Add(
			text.NewCol(12, "As of "+now.Format("January 6, 2006"), props.Text{
				Size:  10,
				Style: fontstyle.Normal,
				Align: align.Center,
			}),
		),
		row.New(5).Add(
			text.NewCol(1, "ID", props.Text{Style: fontstyle.Bold, Align: align.Center}).
				WithStyle(&props.Cell{
					BorderType:      border.Full,
					BorderColor:     &props.Color{Red: 0, Green: 0, Blue: 0}, // Black
					BorderThickness: 0.1,
				}),
			text.NewCol(6, "Product Descriptions", props.Text{Style: fontstyle.Bold, Align: align.Center}).
				WithStyle(&props.Cell{
					BorderType:      border.Full,
					BorderThickness: 0.1,
				}),

			text.NewCol(1, "Stocks", props.Text{Style: fontstyle.Bold, Align: align.Center}).
				WithStyle(&props.Cell{
					BorderType:      border.Full,
					BorderThickness: 0.1,
				}),

			text.NewCol(2, "CostPrice", props.Text{Style: fontstyle.Bold, Align: align.Center}).
				WithStyle(&props.Cell{
					BorderType:      border.Full,
					BorderThickness: 0.1,
				}),

			text.NewCol(2, "SellPrice", props.Text{Style: fontstyle.Bold, Align: align.Center}).
				WithStyle(&props.Cell{
					BorderType:      border.Full,
					BorderThickness: 0.1,
				}),
		),
	)

	// 4. Populate Data Rows
	for _, p := range products {
		costpriceFloat, _ := p.Costprice.Float64()
		formattedCostPrice := humanize.Commaf(costpriceFloat)
		formattedCostPrice = humanize.FormatFloat("#,###.##", costpriceFloat)

		sellpriceFloat, _ := p.Sellprice.Float64()
		formattedSellPrice := humanize.Commaf(sellpriceFloat)
		formattedSellPrice = humanize.FormatFloat("#,###.##", sellpriceFloat)

		m.AddRows(row.New(8).Add(
			// ID (Narrower)
			text.NewCol(1, fmt.Sprintf("%d", p.ID), props.Text{Style: fontstyle.Normal, Align: align.Center}),
			// Descriptions (Wider)
			text.NewCol(6, p.Descriptions),
			// Qty (Narrower)
			text.NewCol(1, fmt.Sprintf("%d", p.Qty), props.Text{Style: fontstyle.Normal, Align: align.Center}),
			// // CostPrice & SellPrice (Narrower)
			text.NewCol(2, formattedCostPrice, props.Text{Align: align.Center}),
			text.NewCol(2, formattedSellPrice, props.Text{Align: align.Center}),
		))
	}

	// 5. Generate and Send PDF
	doc, err := m.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate report"})
		return
	}

	pdfBytes := doc.GetBytes()

	// c.Header("Content-Disposition", "attachment; filename=products_report_2026.pdf")
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Length", fmt.Sprintf("%d", len(pdfBytes)))
	c.Header("Content-Disposition", "attachment; filename=products.pdf")
	c.Data(http.StatusOK, "application/pdf", pdfBytes)
}
