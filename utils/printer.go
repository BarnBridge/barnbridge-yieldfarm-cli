package utils

import (
	"os"

	"github.com/kataras/tablewriter"
	"github.com/landoop/tableprinter"
)

func GetPrinter() *tableprinter.Printer {
	printer := tableprinter.New(os.Stdout)
	printer.BorderTop, printer.BorderBottom, printer.BorderLeft, printer.BorderRight = true, true, true, true
	printer.CenterSeparator = "│"
	printer.ColumnSeparator = "│"
	printer.RowSeparator = "─"
	printer.HeaderBgColor = tablewriter.BgBlackColor // set header background color for all headers.
	printer.HeaderFgColor = tablewriter.FgGreenColor // set header foreground color for all headers.

	return printer
}
