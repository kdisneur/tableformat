package display

import (
	"github.com/kdisneur/tableformat/table"
)

func createTestTable() table.Table {
	return table.Table{
		Headers: []table.Header{
			{
				Name:      "Col 1",
				Alignment: table.LeftAlign{},
				MaxSize:   26,
			},
			{
				Name:      "Col 2 is a long column",
				Alignment: table.CenterAlign{},
				MaxSize:   24,
			},
			{
				Name:      "Col 3",
				Alignment: table.CenterAlign{},
				MaxSize:   22,
			},
			{
				Name:      "Col 4",
				Alignment: table.RightAlign{},
				MaxSize:   7,
			},
		},
		Rows: [][]string{
			[]string{
				"Value 1",
				"Value 2",
				"value 3 can be long",
				"42",
			},
			[]string{
				"Value 1 can also be long",
				"Value 2",
				"value 3",
				"1337",
			},
		},
	}
}
