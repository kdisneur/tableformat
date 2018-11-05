package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kdisneur/tableformat/display"
	"github.com/kdisneur/tableformat/parser"
)

const STDIN_EXAMPLE = `No data in STDIN.

You could:
- pipe data to the command:
  cat file | %s

- inject data from a file:
  %[1]s < file

- inject data with a heredoc string:
  %[1]s <<EOF
  |col1|col2|col3|
  |-|:-:|-:|
  |v1|v2|v3|`

func exit(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}

func main() {
	var output string

	flag.StringVar(&output, "output", "markdown", "output format can be one of: ascii, markdown")

	flag.Parse()

	var displayer display.Displayer

	switch output {
	case "ascii":
		displayer = display.ASCII{}
	case "":
	case "markdown":
		displayer = display.Markdown{}
	default:
		exit("output option must be one of ascii or markdown")
	}

	info, _ := os.Stdin.Stat()
	if info.Size() <= 0 {
		exit(fmt.Sprintf(STDIN_EXAMPLE, os.Args[0]))
	}

	parser := parser.Markdown{}

	table, err := parser.ParseFromInput(os.Stdin)

	if err != nil {
		exit(err.Error())
	}

	fmt.Printf(displayer.Display(table))
}
