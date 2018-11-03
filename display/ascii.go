package display

import (
	"fmt"

	"github.com/kdisneur/tableformat/table"
)

type ASCII struct{}

func (d ASCII) Display(table *table.Table) string {
	var display string

	display = d.displayHeaders(table.Headers)
	for _, row := range table.Rows {
		display += d.displayRow(table.Headers, row)
	}
	display += d.displayFooter(table.Headers)

	return display
}

func (d ASCII) displayHeaders(headers []table.Header) string {
	var headerNames []string

	for _, header := range headers {
		headerNames = append(headerNames, header.Name)
	}

	display := d.displayFullLine(headers, "┌", "┬", "┐")

	return display + d.displayContentLine(headers, headerNames)
}

func (d ASCII) displayRow(headers []table.Header, columns []string) string {
	display := d.displayFullLine(headers, "├", "┼", "┤")

	return display + d.displayContentLine(headers, columns)
}

func (d ASCII) displayFooter(headers []table.Header) string {
	return d.displayFullLine(headers, "└", "┴", "┘")
}

func (d ASCII) displayContentLine(headers []table.Header, columns []string) string {
	var display string

	for index, header := range headers {
		var beforeSpaces int

		switch header.Alignment.(type) {
		case table.LeftAlign:
			beforeSpaces = 1
		case table.RightAlign:
			beforeSpaces = header.MaxSize - (len(columns[index]) + 1)
		case table.CenterAlign:
			beforeSpaces = (header.MaxSize - len(columns[index])) / 2
		}

		afterSpaces := header.MaxSize - len(columns[index]) - beforeSpaces

		display += fmt.Sprintf("│%*s%s%*s", beforeSpaces, " ", columns[index], afterSpaces, " ")
	}
	display += fmt.Sprintf("%s\n", "│")

	return display
}

func (d ASCII) displayFullLine(headers []table.Header, start string, middle string, end string) string {
	var display string

	for index, header := range headers {
		if index == 0 {
			display += fmt.Sprintf(start)
		} else {
			display += fmt.Sprintf(middle)
		}

		for i := 0; i < header.MaxSize; i++ {
			display += fmt.Sprintf("%s", "─")
		}
	}
	display += fmt.Sprintf("%s\n", end)

	return display
}
