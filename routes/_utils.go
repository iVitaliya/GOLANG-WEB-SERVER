package routes

import (
	"fmt"
	"github.com/iVitaliya/colors-go"
)

const (
	InfoSeverity string = "INFO"
	DebugSeverity string = "DEBUG"
	WarningSeverity string = "WARNING"
	ErrorSeverity string = "ERROR"
	FatalSeverity string = "FATAL"
)

type ILogger struct {
	Log func(message string)
}

func Log(Severity string, message ...interface{}) {
	mer := func (m ...interface{}) {
		var mm string
		
		for i := 0; i < len(message); i++ {
			mm = mm + message[i]
		}

		return mm
	}

	var (
		msg string = mer(message)
		state string
	)

	go func() {
		switch Severity {
		case InfoSeverity:
			state = colors.Dim(colors.BrightGreen(InfoSeverity))
			break
		case DebugSeverity:
			state = colors.Green(DebugSeverity)
			break
		case WarningSeverity:
			state = colors.BrightYellow(WarningSeverity)
			break
		case ErrorSeverity:
			state = colors.BrightRed(ErrorSeverity)
			break
		case FatalSeverity:
			state = colors.Red(FatalSeverity)
			break
		}
	}

	m := 
}