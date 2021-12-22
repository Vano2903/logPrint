package logPrint

import (
	"fmt"
)

var (
	colorReset  string = "\033[0m"
	colorRed    string = "\033[31m"
	colorGreen  string = "\033[32m"
	colorYellow string = "\033[33m"
	colorBlue   string = "\033[34m"
	colorCyan   string = "\033[36m"

	IDLE     string = string(colorYellow) + "-" + string(colorReset)
	OK       string = string(colorGreen) + "V" + string(colorReset)
	ERR      string = string(colorRed) + "X" + string(colorReset)
	CRITICAL string = string(colorRed) + "!" + string(colorReset)
	INFO     string = string(colorBlue) + "i" + string(colorReset)
	WARN     string = string(colorYellow) + "!" + string(colorReset)
	DEBUG    string = string(colorCyan) + "d" + string(colorReset)
)

func Idle(msg string) {
	fmt.Printf("[%s] %s\n", IDLE, msg)
}

func Idlef(msg string, args ...interface{}) {
	fmt.Printf("[%s] %s\n", IDLE, fmt.Sprintf(msg, args...))
}

func Info(msg string) {
	fmt.Printf("[%s] %s\n", INFO, msg)
}

func Infof(msg string, args ...interface{}) {
	fmt.Printf("[%s] %s\n", INFO, fmt.Sprintf(msg, args...))
}

func Debug(msg string) {
	fmt.Printf("[%s] %s\n", DEBUG, msg)
}

func Debugf(msg string, args ...interface{}) {
	fmt.Printf("[%s] %s\n", DEBUG, fmt.Sprintf(msg, args...))
}

func Warn(msg string) {
	fmt.Printf("[%s] %s\n", WARN, string(colorYellow)+msg+string(colorReset))
}

func Warnf(msg string, args ...interface{}) {
	fmt.Printf("[%s] %s\n", WARN, string(colorYellow)+fmt.Sprintf(msg, args...)+string(colorReset))
}

func Critical(msg string) {
	fmt.Printf("[%s] %s\n", CRITICAL, string(colorRed)+msg+string(colorReset))
}

func Criticalf(msg string, args ...interface{}) {
	fmt.Printf("[%s] %s\n", CRITICAL, string(colorRed)+fmt.Sprintf(msg, args...)+string(colorReset))
}

func Error(msg string) {
	fmt.Printf("[%s] %s\n", ERR, msg)
}

func Errorf(msg string, args ...interface{}) {
	fmt.Printf("[%s] %s\n", ERR, fmt.Sprintf(msg, args...))
}

func Ok(msg string) {
	fmt.Printf("[%s] %s\n", OK, msg)
}

func Okf(msg string, args ...interface{}) {
	fmt.Printf("[%s] %s\n", OK, fmt.Sprintf(msg, args...))
}
