package display

import (
	"fmt"

	"github.com/kdisneur/tableformat/table"
)

type Markdown struct{}

func (d Markdown) Display(table *table.Table) string {
	var display string

	display = d.displayHeaders(table.Headers)
	for _, row := range table.Rows {
		display += d.displayRow(table.Headers, row)
	}

	return display
}

func (d Markdown) displayHeaders(headers []table.Header) string {
	var headerNames []string

	for _, header := range headers {
		headerNames = append(headerNames, header.Name)
	}

	display := d.displayRow(headers, headerNames)
	return display + d.displayHeaderSeparator(headers)
}

func (d Markdown) displayRow(headers []table.Header, columns []string) string {
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

		display += fmt.Sprintf("|%*s%s%*s", beforeSpaces, " ", columns[index], afterSpaces, " ")
	}
	display += fmt.Sprintf("%s\n", "|")

	return display
}

func (d Markdown) displayHeaderSeparator(headers []table.Header) string {
	var display string

	for _, header := range headers {
		var startingMarker bool
		var endingMarker bool

		switch header.Alignment.(type) {
		case table.LeftAlign:
			startingMarker = true
			endingMarker = false
		case table.RightAlign:
			startingMarker = false
			endingMarker = true
		case table.CenterAlign:
			startingMarker = true
			endingMarker = true
		}

		display += fmt.Sprintf("|")
		maxSize := header.MaxSize

		if startingMarker {
			maxSize--
		}
		if endingMarker {
			maxSize--
		}

		if startingMarker {
			display += fmt.Sprintf(":")
		}

		for i := 0; i < maxSize; i++ {
			display += fmt.Sprintf("%s", "-")
		}

		if endingMarker {
			display += fmt.Sprintf(":")
		}
	}
	display += fmt.Sprintf("|\n")

	return display
}
