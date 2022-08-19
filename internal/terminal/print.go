package terminal

import (
	"io"
	"os"

	"github.com/fatih/color"
)

var (
	exitFunc func(int)
	writer   io.Writer
)

func init() {
	SetExitFunc(os.Exit)
	SetWriter(os.Stdout)
}

func SetExitFunc(f func(int)) {
	exitFunc = f
}

func SetWriter(w io.Writer) {
	writer = w
}

func Error(args ...interface{}) {
	_, _ = colorError.Fprintln(writer, args...)
}

func Errorf(format string, args ...interface{}) {
	_, _ = colorError.Fprintf(writer, format, args...)
}

func ErrorZebra(args ...interface{}) {
	zebraStripe(colorError, colorErrorAlt, args...)
}

func Info(args ...interface{}) {
	_, _ = colorInfo.Fprintln(writer, args...)
}

func Infof(format string, args ...interface{}) {
	_, _ = colorInfo.Fprintf(writer, format, args...)
}

func InfoAlt(args ...interface{}) {
	_, _ = colorInfoAlt.Fprintln(writer, args...)
}

func InfoAltf(format string, args ...interface{}) {
	_, _ = colorInfoAlt.Fprintf(writer, format, args...)
}

func InfoZebra(args ...interface{}) {
	zebraStripe(colorInfo, colorInfoAlt, args...)
}

func Panic(args ...interface{}) {
	_, _ = colorPanic.Fprintln(writer, args...)
	exitFunc(1)
}

func Panicf(format string, args ...interface{}) {
	_, _ = colorPanic.Fprintf(writer, format, args...)
	exitFunc(1)
}

func Standard(args ...interface{}) {
	_, _ = colorStandard.Fprintln(writer, args...)
}

func Standardf(format string, args ...interface{}) {
	_, _ = colorStandard.Fprintf(writer, format, args...)
}

func StandardZebra(args ...interface{}) {
	zebraStripe(colorStandard, colorStandardAlt, args...)
}

func Success(args ...interface{}) {
	_, _ = colorSuccess.Fprintln(writer, args...)
}

func Successf(format string, args ...interface{}) {
	_, _ = colorSuccess.Fprintf(writer, format, args...)
}

func SuccessZebra(args ...interface{}) {
	zebraStripe(colorSuccess, colorSuccessAlt, args...)
}

func Warn(args ...interface{}) {
	_, _ = colorWarn.Fprintln(writer, args...)
}

func Warnf(format string, args ...interface{}) {
	_, _ = colorWarn.Fprintf(writer, format, args...)
}

func WarnZebra(args ...interface{}) {
	zebraStripe(colorWarn, colorWarnAlt, args...)
}

func zebraStripe(primary, secondary *color.Color, args ...interface{}) {
	for i := 0; i < len(args); i++ {
		if i%2 == 0 {
			_, _ = primary.Fprintln(writer, args[i])
		} else {
			_, _ = secondary.Fprintln(writer, args[i])
		}
	}
}
