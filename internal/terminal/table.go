package terminal

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const (
	DefaultAlignment Align = iota
	LeftAlignment
	CenterAlignment
	RightAlignment
)

type Align int

type Header struct {
	value string
	width int
	align Align
}

func NewHeader(s string) *Header {
	return NewHeaderWithAlignment(s, DefaultAlignment)
}

func NewHeaderWithAlignment(s string, a Align) *Header {
	return &Header{
		value: s,
		width: len(s) + 2,
		align: a,
	}
}

type TableOption func(t *Table)

func SetBorderColor(c Color) TableOption {
	return func(t *Table) {
		t.borderColor = GetColor(c)
	}
}

func SetHeaderColor(c Color) TableOption {
	return func(t *Table) {
		t.headerColor = GetColor(c)
	}
}

func SetHeaders(headers ...*Header) TableOption {
	return func(t *Table) {
		t.headers = headers
	}
}

func SetRows(rows ...[]string) TableOption {
	return func(t *Table) {
		t.rows = rows
	}
}

func ZebraStripe(c, cAlt Color) TableOption {
	return func(t *Table) {
		t.zebraStripe = true
		t.primary = GetColor(c)
		t.alt = GetAltColor(cAlt)
	}
}

type Table struct {
	headers     []*Header
	rows        [][]string
	headerColor *color.Color
	borderColor *color.Color
	zebraStripe bool
	primary     *color.Color
	alt         *color.Color
	divider     string
}

func NewTable(options ...TableOption) *Table {
	t := &Table{
		headers:     make([]*Header, 0),
		rows:        make([][]string, 0),
		headerColor: colorDefaultInverse,
		borderColor: colorDefault,
		zebraStripe: true,
		primary:     colorStandard,
		alt:         colorStandardAlt,
	}

	for _, fn := range options {
		fn(t)
	}

	t.getWidths()

	return t
}

func (t *Table) Render() {
	t.setDivider()

	builder := strings.Builder{}
	//builder.WriteString(strings.ReplaceAll(t.divider, "|", "-"))
	builder.WriteString(t.renderHeaders())

	for r := 0; r < len(t.rows); r++ {
		builder.WriteString(t.divider)

		if r%2 == 0 {
			builder.WriteString(t.renderRow(t.primary, t.rows[r]))
		} else {
			builder.WriteString(t.renderRow(t.alt, t.rows[r]))
		}
	}

	//builder.WriteString(strings.ReplaceAll(t.divider, "|", "-"))

	fmt.Println(builder.String())
}

func (t *Table) setDivider() {
	divider := strings.Builder{}

	for h := 0; h < len(t.headers); h++ {
		divider.WriteString(t.borderColor.Sprintf("|"))
		divider.WriteString(t.borderColor.Sprint(strings.Repeat("-", t.headers[h].width+2)))
	}

	divider.WriteString(t.borderColor.Sprintf("|\n"))

	t.divider = divider.String()
}

func (t *Table) renderHeaders() string {
	builder := strings.Builder{}

	for h := 0; h < len(t.headers); h++ {
		builder.WriteString(t.borderColor.Sprintf("|"))
		builder.WriteString(fmt.Sprintf(" %-*s ", t.headers[h].width, t.headers[h].value))
	}

	builder.WriteString(t.borderColor.Sprintf("|\n"))

	return builder.String()
}

func (t *Table) renderRow(clr *color.Color, cells []string) string {
	builder := strings.Builder{}

	for c := 0; c < len(cells); c++ {
		builder.WriteString(t.borderColor.Sprintf("|"))

		switch t.headers[c].align {
		case CenterAlignment:
			left := (len(cells[c]) - t.headers[c].width) / 2

			if left < 0 {
				left = left * -1
			}

			right := t.headers[c].width - (len(cells[c]) + left)

			builder.WriteString(clr.Sprintf("%*s%s%-*s", left, "", cells[c], right, ""))
		case RightAlignment:
			builder.WriteString(clr.Sprintf(" %*s ", t.headers[c].width, cells[c]))
		default:
			builder.WriteString(clr.Sprintf(" %-*s ", t.headers[c].width, cells[c]))
		}
	}

	builder.WriteString(t.borderColor.Sprintf("|\n"))

	return builder.String()
}

func (t *Table) getWidths() {
	for h := 0; h < len(t.headers); h++ {
		for r := 0; r < len(t.rows); r++ {
			t.headers[h].width = t.max(t.headers[h].width, len(t.rows[r][h])+2)
		}
	}
}

func (t *Table) max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
