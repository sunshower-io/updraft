package ir

import (
	"bytes"
	"fmt"
	"strings"
)

type CrossReferencer struct {
}

const nameFormat = "%16s"

func (c *CrossReferencer) Print(st SymbolTableStack) string {

	b := new(bytes.Buffer)

	printHeader(b)

	printTableHeaders(b)

	printValues(st, b)

	return b.String()
}

func printValues(stack SymbolTableStack, b *bytes.Buffer) {
	tables := stack.Tables()
	for _, t := range tables {
		syms := t.Entries(true)
		for _, s := range syms {
			name := s.GetName()
			b.WriteString(name)
			b.WriteString(strings.Repeat(" ", 21-len(name)))
			lines := s.GetLines()
			for _, l := range lines {
				b.WriteString(fmt.Sprintf("%03d ", l.Number))
			}
			b.WriteString("\n")
		}
	}
}

func printTableHeaders(buffer *bytes.Buffer) {

	li := len("Identifier")
	ln := len("Line Numbers")
	sp := strings.Repeat(" ", 10)

	buffer.WriteString("Identifier")
	buffer.WriteString(sp)
	buffer.WriteString("Line Numbers\n")
	buffer.WriteString(strings.Repeat("-", li))
	buffer.WriteString(sp)
	buffer.WriteString(strings.Repeat("-", ln))
	buffer.WriteString("\n")
}

func printHeader(buffer *bytes.Buffer) {
	prefix := strings.Repeat("=", 5)
	buffer.WriteString(fmt.Sprintf("%s SYMBOL TABLE %s\n", prefix, prefix))
}
