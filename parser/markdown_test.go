package parser

import (
	"bytes"
	"testing"

	"github.com/kdisneur/tableformat/table"
)

func TestWhenStringIsEmpty(t *testing.T) {
	content := ""
	reader := bytes.NewBufferString(content)
	table, err := Markdown{}.ParseFromInput(reader)

	if err == nil {
		t.Errorf("An error should have been raised. Current table: %+v", table)
	}

	if _, valid := err.(*EmptyTable); !valid {
		t.Errorf("Wrong error occured. Expected empty table received: %+v", err)
	}
}

func TestWhenStringHasNoHeaderSeparator(t *testing.T) {
	content := `| col 1 | col 2| col 3 |`

	reader := bytes.NewBufferString(content)
	table, err := Markdown{}.ParseFromInput(reader)

	if err == nil {
		t.Errorf("An error should have been raised. Current table: %+v", table)
	}

	if _, valid := err.(*MissingHeaderSeparator); !valid {
		t.Errorf("Wrong error occured. Expected missing header separator received: %+v", err)
	}
}

func TestWhenEverythingIsFine(t *testing.T) {
	content := `| col 1 | col 2| col 3 | col 4 |
              |--|:-:|-:|:-|
              value 1| value 2|value 3|value 4 |
              |value 5| value 6| value 7|`

	reader := bytes.NewBufferString(content)
	receivedTable, err := Markdown{}.ParseFromInput(reader)

	if err != nil {
		t.Fatalf("An error should not be raised. Current error: %+v", err)
	}

	expectedHeaders := []table.Header{
		{Name: "col 1", Alignment: table.LeftAlign{}, MaxSize: 9},
		{Name: "col 2", Alignment: table.CenterAlign{}, MaxSize: 9},
		{Name: "col 3", Alignment: table.RightAlign{}, MaxSize: 9},
		{Name: "col 4", Alignment: table.LeftAlign{}, MaxSize: 9},
	}

	if len(receivedTable.Headers) != len(expectedHeaders) {
		t.Fatalf("The headers should contain %d entries. Current table: %+v", len(expectedHeaders), receivedTable)
	}

	for index, expectedHeader := range expectedHeaders {
		if expectedHeader.Name != receivedTable.Headers[index].Name {
			t.Errorf("The header %d should have the name '%s' but has '%s'", index, expectedHeader.Name, receivedTable.Headers[index].Name)
		}

		if expectedHeader.Alignment != receivedTable.Headers[index].Alignment {
			t.Errorf("The header %d should have the alignment %T but has %T", index, expectedHeader.Alignment, receivedTable.Headers[index].Alignment)
		}

		if expectedHeader.MaxSize != receivedTable.Headers[index].MaxSize {
			t.Errorf("The header %d should have the max size %d but has %d", index, expectedHeader.MaxSize, receivedTable.Headers[index].MaxSize)
		}
	}

	expectedRows := [][]string{
		[]string{"value 1", "value 2", "value 3", "value 4"},
		[]string{"value 5", "value 6", "value 7", " "},
	}

	for rowIndex, row := range expectedRows {
		for columnIndex, column := range row {
			if column != receivedTable.Rows[rowIndex][columnIndex] {
				t.Errorf("The rows should have be the same but are not. Expected:\n%q\nReceived:\n%q", row, receivedTable.Rows[rowIndex])
			}
		}
	}
}
