package display

import (
	"testing"
)

func TestMarkdownWithLeftCenterRightAlignment(t *testing.T) {
	table := createTestTable()

	expected := `| Col 1                    | Col 2 is a long column |        Col 3         | Col 4 |
|:-------------------------|:----------------------:|:--------------------:|------:|
| Value 1                  |        Value 2         | value 3 can be long  |    42 |
| Value 1 can also be long |        Value 2         |       value 3        |  1337 |
`

	received := Markdown{}.Display(&table)

	if expected != received {
		t.Errorf("expected: \n%s\n received: \n%s\n", expected, received)
	}
}
