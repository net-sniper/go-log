/*
Package log provides support for logging to stdout and stderr.

Log entries will be logged in the following format:

    timestamp hostname tag[pid]: SEVERITY Message
*/
package log

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"

	log "github.com/Sirupsen/logrus"
)

type Formatter struct {
	once *sync.Once
}

// tag represents the application name generating the log message. The tag
// string will appear in all log entires.
var (
	formatter = &Formatter{&sync.Once{}}
	tag       string
	file      string
	line      int
)

func (c *Formatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Format(time.RFC3339)
	hostname, _ := os.Hostname()
	return []byte(fmt.Sprintf("%s %s : %s\t%s:%d[%d] %s\n", timestamp, hostname, strings.ToUpper(entry.Level.String()), file, line, os.Getpid(), entry.Message)), nil
}

func Init(logFile, logLevel string) {
	init := func() {
		if logLevel == "" {
			logLevel = "debug"
		}

		tag = os.Args[0]
		log.SetFormatter(formatter)
		SetLevel(logLevel)

		if err := os.MkdirAll(path.Dir(logFile), os.ModeDir); err != nil {
			Fatal(fmt.Sprintf(`create log file dir error: "%s".`, path.Dir(logFile)))
		}

		f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			Fatal(fmt.Sprintf(`can not open log file: "%s".`, logFile))
		}
		log.SetOutput(f)

	}

	formatter.once.Do(init)
}

// SetTag sets the tag.
func SetTag(t string) {
	tag = t
}

// SetLevel sets the log level. Valid levels are panic, fatal, error, warn, info and debug.
func SetLevel(level string) {
	lvl, err := log.ParseLevel(level)
	if err != nil {
		Fatal(fmt.Sprintf(`not a valid level: "%s"`, level))
	}
	log.SetLevel(lvl)
}

// Debug logs a message with severity DEBUG.
func Debug(v ...interface{}) {
	_, file, line, _ = runtime.Caller(1)
	log.Debug(fmt.Sprint(v...))
}

// Error logs a message with severity ERROR.
func Error(v ...interface{}) {
	_, file, line, _ = runtime.Caller(1)
	log.Error(fmt.Sprint(v...))
}

// Fatal logs a message with severity ERROR followed by a call to os.Exit().
func Fatal(v ...interface{}) {
	_, file, line, _ = runtime.Caller(1)
	log.Fatal(fmt.Sprint(v...))
}

// Info logs a message with severity INFO.
func Info(v ...interface{}) {
	_, file, line, _ = runtime.Caller(1)
	log.Info(fmt.Sprint(v...))
}

// Warning logs a message with severity WARNING.
func Warning(v ...interface{}) {
	_, file, line, _ = runtime.Caller(1)
	log.Warning(fmt.Sprint(v...))
}

func Debugf(format string, v ...interface{}) {
	_, file, line, _ = runtime.Caller(1)
	log.Debug(fmt.Sprintf(format, v...))
}

// Error logs a message with severity ERROR.
func Errorf(format string, v ...interface{}) {
	_, file, line, _ = runtime.Caller(1)
	log.Error(fmt.Sprintf(format, v...))
}

// Fatal logs a message with severity ERROR followed by a call to os.Exit().
func Fatalf(format string, v ...interface{}) {
	_, file, line, _ = runtime.Caller(1)
	log.Fatal(fmt.Sprintf(format, v...))
}

// Info logs a message with severity INFO.
func Infof(format string, v ...interface{}) {
	_, file, line, _ = runtime.Caller(1)
	log.Info(fmt.Sprintf(format, v...))
}

// Warning logs a message with severity WARNING.
func Warningf(format string, v ...interface{}) {
	_, file, line, _ = runtime.Caller(1)
	log.Warning(fmt.Sprintf(format, v...))
}
