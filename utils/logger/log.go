package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/logrusorgru/aurora"
)

//severity: Severidade do log
type Severity string

//Níveis de severidade
const (
	debug   Severity = "DEBUG"
	info    Severity = "INFO"
	warn    Severity = "WARN"
	failure Severity = "ERROR"
)

var (
	Level Severity
)

func DefineLoggingLevel() {
	loggingLevel := os.Getenv("LOGGING_LEVEL")

	if loggingLevel == "" {
		Level = info
	} else {
		Level = Enum(loggingLevel)
	}
}

//Debug : Debug log message
func Debug(message ...interface{}) {
	if Level == debug {
		log(debug, message)
	}
}

//Info : Info log message
func Info(message ...interface{}) {
	if Level == debug || Level == info {
		log(info, message)
	}
}

//Warn : Warn log message
func Warn(message ...interface{}) {
	log(warn, message)
}

//Error : Error log message
func Error(message ...interface{}) {
	isNull := false

	if len(message) > 0 {
		for _, msg := range message {
			if msg == nil {
				isNull = true
			}
		}
	}

	if !isNull {
		log(failure, message)
	}
}

func Enum(value string) Severity {
	return Severity(value)
}

func log(tag Severity, message ...interface{}) {
	pc, _, _, _ := runtime.Caller(2)
	function := fmt.Sprintf("[%s]", runtime.FuncForPC(pc).Name())

	date := time.Now().Format("02/01/2006 15:04:05")

	var stacktrace string
	var content string

	for _, msg := range message {
		if tag == "ERROR" {
			switch datatype := msg.(type) {
			case []interface{}:
				for index, element := range datatype {
					_, error := json.Marshal(element)

					if index != 0 {
						stacktrace += fmt.Sprintf("[%s] ", fmt.Sprint(element))
						stacktrace = strings.ReplaceAll(stacktrace, "\"", "")
					}

					if error != nil {
						fail(date, function, error)
					}
				}

				stacktrace = fmt.Sprintf("{%s}", stacktrace)
				stacktrace = strings.ReplaceAll(stacktrace, "] }", "]}")
			}
		}

		content += fmt.Sprintf("%s ", fmt.Sprint(msg))
	}

	switch tag {
	case debug:
		fmt.Println(
			aurora.Bold(
				aurora.Cyan(
					fmt.Sprintf("[%s]", tag),
				),
			),
			aurora.Cyan(date),
			aurora.Cyan(function),
			aurora.Cyan(content),
		)
	case info:
		fmt.Println(
			aurora.Bold(
				aurora.Green(
					fmt.Sprintf("[%s]", tag),
				),
			),
			aurora.Green(date),
			aurora.Green(function),
			aurora.Green(content),
		)
	case warn:
		fmt.Println(
			aurora.Bold(
				aurora.Yellow(
					fmt.Sprintf("[%s]", tag),
				),
			),
			aurora.Yellow(date),
			aurora.Yellow(function),
			aurora.Yellow(content),
		)
	case failure:
		fail(date, function, content)
	}
}

func fail(date string, function string, message interface{}) {
	fmt.Println(
		aurora.Bold(
			aurora.Red("[ERROR]"),
		),
		aurora.Red(date),
		aurora.Red(function),
		aurora.Red(message),
	)
}
