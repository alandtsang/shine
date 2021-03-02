package logs

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

var debug bool

func init() {
	debug = os.Getenv("DEBUG") != ""
}

// Info
func Info(v ...interface{}) {
	log.Print(v...)
}

// Infof
func Infof(format string, v ...interface{}) {
	log.Printf(strings.Join([]string{"[Info]", format}, " "), v...)
}

// Debug
func Debug(format string, v ...interface{}) {
	if debug {
		log.Print(v...)
	}
}

// Debugf
func Debugf(format string, v ...interface{}) {
	if debug {
		log.Printf(strings.Join([]string{"[Info]", format}, " "), v...)
	}
}

// Error
func Error(v ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	log.Printf(fmt.Sprintf("[Error] [%s] %s:%d", runtime.FuncForPC(pc).Name(), fn, line), v...)
}

// Errorf
func Errorf(format string, v ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	log.Printf(fmt.Sprintf("[Error] [%s] %s:%d %s", runtime.FuncForPC(pc).Name(), fn, line, format), v...)
}

// Fatal
func Fatal(v ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	log.Printf(fmt.Sprintf("[Fatal] [%s] %s:%d", runtime.FuncForPC(pc).Name(), fn, line), v...)
}

// Fatalf
func Fatalf(format string, v ...interface{}) {
	pc, fn, line, _ := runtime.Caller(1)
	log.Printf(fmt.Sprintf("[Fatal] [%s] %s:%d %s", runtime.FuncForPC(pc).Name(), fn, line, format), v...)
}
