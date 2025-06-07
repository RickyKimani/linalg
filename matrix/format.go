package matrix

import (
	"fmt"
	"io"
)

// Format implements the fmt.Formatter interface
func (m Matrix[T]) Format(f fmt.State, verb rune) {
	if len(m) == 0 {
		io.WriteString(f, "[]")
		return
	}

	switch verb {
	case 'v', 's', 'f':

	default:
		fmt.Fprintf(f, "%%!%c(Matrix)", verb)
		return
	}

	maxCols := 0
	for _, row := range m {
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}

	colWidths := make([]int, maxCols)
	for i := range m {
		for j, val := range m[i] {
			if j < len(colWidths) {
				width := len(fmt.Sprintf("%v", val))
				if width > colWidths[j] {
					colWidths[j] = width
				}
			}
		}
	}

	precision := 4
	if p, ok := f.Precision(); ok {
		precision = p
	}

	io.WriteString(f, "{\n")
	for i := range m {
		io.WriteString(f, "  [")
		for j, val := range m[i] {
			if j >= len(colWidths) {
				continue
			}

			if j > 0 {
				io.WriteString(f, ", ")
			}

			switch verb {
			case 'v', 's':
				format := fmt.Sprintf("%%%dv", colWidths[j])
				fmt.Fprintf(f, format, val)
			case 'f':
				if floatVal, ok := any(val).(float64); ok {
					format := fmt.Sprintf("%%%d.%df", colWidths[j], precision)
					fmt.Fprintf(f, format, floatVal)
				} else {
					format := fmt.Sprintf("%%%dv", colWidths[j])
					fmt.Fprintf(f, format, val)
				}
			}
		}
		io.WriteString(f, "]")
		if i < len(m)-1 {
			io.WriteString(f, ",")
		}
		io.WriteString(f, "\n")
	}
	io.WriteString(f, "}")
}
