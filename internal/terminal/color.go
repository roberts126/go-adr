package terminal

import "github.com/fatih/color"

const (
	ColorDefault Color = iota
	ColorDefaultInverse
	ColorDefaultAlt
	ColorError
	ColorErrorAlt
	ColorInfo
	ColorInfoInverse
	ColorInfoAlt
	ColorPanic
	ColorSuccess
	ColorSuccessInverse
	ColorSuccessAlt
	ColorStandard
	ColorStandardAlt
	ColorWarn
	ColorWarnInverse
	ColorWarnAlt
)

var (
	colorDefault        = color.New(color.FgHiWhite, color.Bold)
	colorDefaultInverse = color.New(color.BgHiWhite, color.FgBlack, color.Bold)
	colorDefaultAlt     = color.New(color.FgWhite, color.Bold)
	colorError          = color.New(color.FgHiRed, color.Bold)
	colorErrorAlt       = color.New(color.FgRed, color.Bold)
	colorInfo           = color.New(color.FgHiCyan, color.Bold)
	colorInfoInverse    = color.New(color.BgHiCyan, color.FgWhite, color.Bold)
	colorInfoAlt        = color.New(color.FgCyan, color.Bold)
	colorPanic          = color.New(color.BgRed, color.FgHiWhite, color.Bold)
	colorSuccess        = color.New(color.FgHiGreen, color.Bold)
	colorSuccessInverse = color.New(color.BgHiGreen, color.FgHiWhite, color.Bold)
	colorSuccessAlt     = color.New(color.FgGreen, color.Bold)
	colorStandard       = color.New(color.FgHiWhite, color.Bold)
	colorStandardAlt    = color.New(color.FgWhite, color.Bold)
	colorWarn           = color.New(color.FgYellow, color.Bold)
	colorWarnInverse    = color.New(color.BgYellow, color.FgBlack, color.Bold)
	colorWarnAlt        = color.New(color.FgHiYellow, color.Bold)
)

var colors = map[Color]*color.Color{
	ColorDefault:        colorDefault,
	ColorDefaultAlt:     colorDefaultAlt,
	ColorDefaultInverse: colorDefaultInverse,
	ColorError:          colorError,
	ColorErrorAlt:       colorErrorAlt,
	ColorInfo:           colorInfo,
	ColorInfoAlt:        colorInfoAlt,
	ColorInfoInverse:    colorInfoInverse,
	ColorPanic:          colorPanic,
	ColorSuccess:        colorSuccess,
	ColorSuccessAlt:     colorSuccessAlt,
	ColorSuccessInverse: colorSuccessInverse,
	ColorStandard:       colorStandard,
	ColorStandardAlt:    colorStandardAlt,
	ColorWarn:           colorWarn,
	ColorWarnInverse:    colorWarnInverse,
	ColorWarnAlt:        colorWarnAlt,
}

type Color int

func GetColor(c Color) *color.Color {
	return getColor(c, colorDefault)
}

func GetAltColor(c Color) *color.Color {
	return getColor(c, colorDefaultAlt)
}

func getColor(c Color, d *color.Color) *color.Color {
	c2, ok := colors[c]
	if !ok {
		c2 = d
	}

	return c2
}
