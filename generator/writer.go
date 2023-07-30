package generator

import "strings"

type Writer struct {
	output string
	indent int
}

func NewWriter() *Writer {
	return &Writer{}
}

func (w *Writer) Indent(by int) {
	w.indent += by
}

func (w *Writer) WriteLine(line string) {
	for i := 0; i < w.indent; i++ {
		w.output += "    "
	}

	w.output += line
	w.output += "\n"
}

func (w *Writer) WriteWriter(child *Writer) {
	output := child.output
	if output == "" {
		return
	}

	output = strings.TrimSuffix(output, "\n")

	lines := strings.Split(output, "\n")

	for _, l := range lines {
		w.WriteLine(l)
	}
}

func (w *Writer) Output() string {
	return w.output
}
