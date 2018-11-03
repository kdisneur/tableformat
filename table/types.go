package table

type Table struct {
	Headers []Header
	Rows    [][]string
}

type Header struct {
	Name string
	Alignment
	MaxSize int
}

type Alignment interface {
	IsAlignment() bool
}

type CenterAlign struct{}
type LeftAlign struct{}
type RightAlign struct{}

func (a CenterAlign) IsAlignment() bool { return true }
func (a LeftAlign) IsAlignment() bool   { return true }
func (a RightAlign) IsAlignment() bool  { return true }
