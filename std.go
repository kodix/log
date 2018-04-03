package log

import (
	"os"
	"log"
	"io"
)

var std = New(os.Stdout, "", Error, log.LstdFlags)

func SetOutput(w io.Writer) {
	std.SetOutput(w)
}

func SetFlags(flag int) {
	std.SetFlags(flag)
}

func SetPrefix(prefix string) {
	std.SetPrefix(prefix)
}

func SetVerbosity(lvl Verbosity) {
	std.SetVerbosity(lvl)
}

func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

func Fatalln(args ...interface{}) {
	std.Fatalln(args...)
}

func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

func Panicln(args ...interface{}) {
	std.Panicln(args...)
}

func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

func Errorln(args ...interface{}) {
	std.Errorln(args...)
}

func Warningf(format string, args ...interface{}) {
	std.Warningf(format, args...)
}

func Warningln(args ...interface{}) {
	std.Warningln(args...)
}

func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

func Infoln(args ...interface{}) {
	std.Infoln(args...)
}

func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

func Debugln(args ...interface{}) {
	std.Debugln(args...)
}