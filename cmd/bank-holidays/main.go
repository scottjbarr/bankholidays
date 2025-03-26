package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/scottjbarr/bankholidays"
)

func main() {
	srv, err := bankholidays.Build()
	if err != nil {
		panic(err)
	}

	// table output
	data := [][]string{}

	for _, it := range srv.All() {
		values := []string{
			it.Date,
			it.Country,
			it.Name,
		}

		data = append(data, values)
	}

	w := buildTableWriter()
	w.AppendBulk(data)
	w.Render()
}

func buildTableWriter() *tablewriter.Table {
	w := tablewriter.NewWriter(os.Stdout)
	w.SetHeader([]string{"Date", "Country", "Name"})
	w.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	w.SetAutoFormatHeaders(false)
	w.SetCenterSeparator("|")
	w.SetAutoWrapText(false)

	// column alignment
	w.SetColumnAlignment([]int{
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_LEFT,
	})

	return w
}
