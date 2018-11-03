package parser

import (
	"io"

	"github.com/kdisneur/tableformat/table"
)

type TableParser interface {
	ParseFromInput(io.Reader) (*table.Table, error)
}
