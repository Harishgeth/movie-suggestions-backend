package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"

	"github.com/mgutz/ansi"
	"github.com/pborman/uuid"
)

type Logger struct {
	l     *log.Logger
	ref   string
	info  func(string) string
	warn  func(string) string
	err   func(string) string
	debug func(string) string
	fatal func(string) string
}

func NewLogger(ref string) *Logger {
	if ref == "" {
		ref = uuid.New()
	}

	return &Logger{
		l:     log.New(os.Stdout, fmt.Sprintf("[ %s ] ", ref), 0),
		ref:   ref,
		info:  ansi.ColorFunc("green:black"),
		warn:  ansi.ColorFunc("yellow:black"),
		err:   ansi.ColorFunc("red+h:black"),
		debug: ansi.ColorFunc("cyan:black"),
		fatal: ansi.ColorFunc("red:black"),
	}
}

func (l *Logger) Init(ref string) {
	l.ref = ref
	if l.ref == "" {
		l.ref = uuid.New()
	}
	l.l = log.New(os.Stdout, fmt.Sprintf("[ %s ] ", l.ref), 0)
}

func (l *Logger) GetRef() string {
	return l.ref
}

func (l *Logger) Warn(a ...interface{}) {
	l.l.Println(l.warn(fmt.Sprint(formatLog("WARN", a...)...)))
}

func (l *Logger) Error(a ...interface{}) {
	l.l.Println(l.err(fmt.Sprint(formatLog("ERROR", a...)...)))
}

func (l *Logger) Info(a ...interface{}) {
	l.l.Println(l.info(fmt.Sprint(formatLog("INFO", a...)...)))
}

func (l *Logger) Debug(a ...interface{}) {
	l.l.Println(l.debug(fmt.Sprint(formatLog("DEBUG", a...)...)))
}

func (l *Logger) Fatal(a ...interface{}) {
	l.l.Println(l.fatal(fmt.Sprint(formatLog("FATAL", a...)...)))
}

func formatLog(logType string, a ...interface{}) []interface{} {
	var n []interface{}
	n = append(n, "["+logType+"] ")
	_, file, line, _ := runtime.Caller(2)
	// Comment the block if you want the entire file path to be displayed
	short := file
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}
	// end
	file = short
	n = append(n, file+":"+strconv.Itoa(line)+" ")
	n = append(n, a...)
	return n
}
