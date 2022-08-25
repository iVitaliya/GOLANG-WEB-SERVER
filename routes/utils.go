package routes

import (
	"fmt"

	"github.com/iVitaliya/colors-go"
)

const (
	InfoSeverity    string = "INFO"
	DebugSeverity   string = "DEBUG"
	WarningSeverity string = "WARNING"
	ErrorSeverity   string = "ERROR"
	FatalSeverity   string = "FATAL"
)

type ILogger struct {
	Log func(message string)
}

func Log(Severity string, message ...interface{}) {
	mer := func(m ...interface{}) string {
		var mm string

		for i := 0; i < len(message); i++ {
			mm = mm + fmt.Sprint(message[i])
		}

		return mm
	}

	Stater := func() string {
		var st string

		switch Severity {
		case InfoSeverity:
			st = colors.Dim(colors.BrightGreen(InfoSeverity))
			break
		case DebugSeverity:
			st = colors.Green(DebugSeverity)
			break
		case WarningSeverity:
			st = colors.BrightYellow(WarningSeverity)
			break
		case ErrorSeverity:
			st = colors.BrightRed(ErrorSeverity)
			break
		case FatalSeverity:
			st = colors.Red(FatalSeverity)
			break
		}

		return st
	}

	var (
		msg   string = mer(message)
		state string = Stater()
	)

	fmt.Println(FormatString("{0}{1}{2} {3}{4}{5}: {6}", []string{
		colors.Gray("["),
		colors.Dim(colors.BrightYellow(Date())),
		colors.Gray("]"),
		colors.Gray("("),
		state,
		colors.Gray(")"),
		msg,
	}))
}
