package routes

import (
	"fmt"
	"os"

	"github.com/iVitaliya/colors-go"
)

const (
	InfoSeverity    string = "INFO"
	DebugSeverity   string = "DEBUG"
	WarningSeverity string = "WARNING"
	ErrorSeverity   string = "ERROR"
	FatalSeverity   string = "FATAL"
)

type ILoggerInstance struct {
	Configure func(config ILoggerConfig) *ILogger
	Info      func(message string)
	Debug     func(message string)
	Warning   func(message string)
	Error     func(message string)
	Fatal     func(message string)
}

type ILoggerConfig struct {
	Severity string
}

type ILogger struct {
	Log     func(message string)
	Info    func(message string)
	Debug   func(message string)
	Warning func(message string)
	Error   func(message string)
	Fatal   func(message string)
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
		case DebugSeverity:
			st = colors.Green(DebugSeverity)
		case WarningSeverity:
			st = colors.BrightYellow(WarningSeverity)
		case ErrorSeverity:
			st = colors.BrightRed(ErrorSeverity)
		case FatalSeverity:
			st = colors.Red(FatalSeverity)
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

// Initiates a new Logger instance.
var Logger *ILoggerInstance = &ILoggerInstance{
	Configure: func(config ILoggerConfig) *ILogger {
		return &ILogger{
			Log: func(message string) {
				Log(config.Severity, message)
				if config.Severity == FatalSeverity {
					os.Exit(1)
				}
			},
			Info: func(message string) {
				Log(InfoSeverity, message)
			},
			Debug: func(message string) {
				Log(DebugSeverity, message)
			},
			Warning: func(message string) {
				Log(WarningSeverity, message)
			},
			Error: func(message string) {
				Log(ErrorSeverity, message)
			},
			Fatal: func(message string) {
				Log(FatalSeverity, message)
				os.Exit(1)
			},
		}
	},
	Info: func(message string) {
		Log(InfoSeverity, message)
	},
	Debug: func(message string) {
		Log(DebugSeverity, message)
	},
	Warning: func(message string) {
		Log(WarningSeverity, message)
	},
	Error: func(message string) {
		Log(ErrorSeverity, message)
	},
	Fatal: func(message string) {
		Log(FatalSeverity, message)
		os.Exit(1)
	},
}
