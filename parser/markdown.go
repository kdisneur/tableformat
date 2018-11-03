package parser

import (
	"bufio"
	"io"
	"strings"

	"github.com/kdisneur/tableformat/table"
)

type Markdown struct{}

type EmptyTable struct{}

func (e *EmptyTable) Error() string { return "The table is empty" }

type MissingHeaderSeparator struct{}

func (e *MissingHeaderSeparator) Error() string { return "The table doesn't contain header separator" }

func (p Markdown) ParseFromInput(reader io.Reader) (*table.Table, error) {
	content := bufio.NewScanner(reader)

	if content.Scan() {
		rawHeaders := content.Text()
		if content.Scan() {
			rawAlignments := content.Text()
			headers := toHeaders(rawHeaders, rawAlignments)

			var rows [][]string

			for content.Scan() {
				rawRow := cleanupMarkup(content.Text())

				rows = append(rows, toColumns(headers, rawRow))
			}

			table := table.Table{Headers: headers, Rows: rows}

			updateMaxSizes(&table)

			return &table, nil
		} else {
			return nil, &MissingHeaderSeparator{}
		}
	} else {
		return nil, &EmptyTable{}
	}
}

func cleanupMarkup(markup string) string {
	return strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(markup), "|"), "|")
}

func toAlignment(rawAlignment string) table.Alignment {
	if strings.HasPrefix(rawAlignment, ":") && strings.HasSuffix(rawAlignment, ":") {
		return table.CenterAlign{}
	} else if strings.HasSuffix(rawAlignment, ":") {
		return table.RightAlign{}
	}
	return table.LeftAlign{}
}

func toColumns(headers []table.Header, rawRow string) []string {
	var columns []string

	rows := strings.Split(rawRow, "|")

	for index, _ := range headers {
		if len(rows) > index {
			columns = append(columns, strings.TrimSpace(rows[index]))
		} else {
			columns = append(columns, " ")
		}
	}

	return columns
}

func toHeaders(rawHeaders string, rawAlignments string) []table.Header {
	rawHeaders = cleanupMarkup(rawHeaders)
	rawAlignments = cleanupMarkup(rawAlignments)

	alignments := strings.Split(rawAlignments, "|")

	var headers []table.Header

	for index, rawHeaderName := range strings.Split(rawHeaders, "|") {
		headerName := strings.TrimSpace(rawHeaderName)
		var alignment table.Alignment
		if len(alignments) > index {
			alignment = toAlignment(alignments[index])
		} else {
			alignment = table.LeftAlign{}
		}

		headers = append(headers, table.Header{
			Name:      headerName,
			Alignment: alignment,
			MaxSize:   len(headerName) + 2,
		})
	}
	return headers
}

func updateMaxSizes(table *table.Table) {
	for _, row := range table.Rows {
		for index, column := range row {
			newSize := len(column) + 2
			if table.Headers[index].MaxSize < newSize {
				table.Headers[index].MaxSize = newSize
			}
		}
	}
}
