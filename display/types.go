package display

import (
	"github.com/kdisneur/tableformat/table"
)

type Displayer interface {
	Display(*table.Table) string
}
