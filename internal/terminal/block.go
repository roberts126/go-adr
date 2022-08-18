package terminal

import (
	"fmt"
	"strings"
)

const indentFmt = "%s%s"

type Block struct {
	header string
	indent string
	lines  []string
}

type Blocks []*Block

func NewBlock(header string) *Block {
	return &Block{
		header: header,
		indent: "    ",
		lines:  make([]string, 0),
	}
}

func (b *Block) SetIndent(i string) *Block {
	b.indent = i

	return b
}

func (b *Block) Error(args ...interface{}) *Block {
	line := fmt.Sprint(args...)
	line = colorError.Sprintf(indentFmt, b.indent, line)

	b.lines = append(b.lines, line)

	return b
}

func (b *Block) Errorf(format string, args ...interface{}) *Block {
	b.lines = append(b.lines, colorError.Sprintf(b.indent+format, args...))

	return b
}

func (b *Block) Info(args ...interface{}) *Block {
	line := fmt.Sprint(args...)
	line = colorInfo.Sprintf(indentFmt, b.indent, line)

	b.lines = append(b.lines, line)

	return b
}

func (b *Block) Infof(format string, args ...interface{}) *Block {
	b.lines = append(b.lines, colorInfo.Sprintf(b.indent+format, args...))

	return b
}

func (b *Block) Standard(args ...interface{}) *Block {
	line := fmt.Sprint(args...)
	line = colorStandard.Sprintf(indentFmt, b.indent, line)

	b.lines = append(b.lines, line)

	return b
}

func (b *Block) Standardf(format string, args ...interface{}) *Block {
	b.lines = append(b.lines, colorStandard.Sprintf(b.indent+format, args...))

	return b
}

func (b *Block) Success(args ...interface{}) *Block {
	line := fmt.Sprint(args...)
	line = colorSuccess.Sprintf(indentFmt, b.indent, line)

	b.lines = append(b.lines, line)

	return b
}

func (b *Block) Successf(format string, args ...interface{}) *Block {
	b.lines = append(b.lines, colorSuccess.Sprintf(b.indent+format, args...))

	return b
}

func (b *Block) Warn(args ...interface{}) *Block {
	line := fmt.Sprint(args...)
	line = colorWarn.Sprintf(indentFmt, b.indent, line)

	b.lines = append(b.lines, line)

	return b
}

func (b *Block) Warnf(format string, args ...interface{}) *Block {
	b.lines = append(b.lines, colorWarn.Sprintf(b.indent+format, args...))

	return b
}

func (b *Block) Render() {
	_, _ = colorStandard.Fprintln(writer, b.header)
	_, _ = fmt.Fprintf(writer, "%s\n", strings.Join(b.lines, "\n"))
}

func (b *Block) Reset() {
	b.lines = make([]string, 0)
}
