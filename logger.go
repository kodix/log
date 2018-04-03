package log

import (
	"log"
	"io"
	"sync"
)

type Verbosity uint8

const (
	Error   Verbosity = iota
	Warning
	Debug
	Info
)

type Logger struct {
	l *log.Logger

	mu  sync.RWMutex
	lvl Verbosity
}

func New(out io.Writer, prefix string, v Verbosity, flag int) *Logger {
	return &Logger{
		l:   log.New(out, prefix, flag),
		mu: sync.RWMutex{},
		lvl: v,
	}
}

func (l *Logger) SetOutput(w io.Writer) {
	l.l.SetOutput(w)
}

func (l *Logger) SetFlags(flag int) {
	l.l.SetFlags(flag)
}

func (l *Logger) SetPrefix(prefix string) {
	l.l.SetPrefix(prefix)
}

func (l *Logger) SetVerbosity(lvl Verbosity) {
	l.mu.Lock()
	l.lvl = lvl
	l.mu.Unlock()
}

func (l *Logger) Verbosity() Verbosity {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.lvl
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.l.Fatalf(format, args...)
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.l.Fatalln(args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.l.Panicf(format, args...)
}

func (l *Logger) Panicln(args ...interface{}) {
	l.l.Panicln(args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.l.Printf(format, args...)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.l.Println(args...)
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	if l.Verbosity() >= Warning {
		l.l.Printf(format, args...)
	}
}

func (l *Logger) Warningln(args ...interface{}) {
	if l.Verbosity() >= Warning {
		l.l.Println(args...)
	}
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.Verbosity() >= Debug {
		l.l.Printf(format, args...)
	}
}

func (l *Logger) Debugln(args ...interface{}) {
	if l.Verbosity() >= Debug {
		l.l.Println(args...)
	}
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if l.Verbosity() >= Info {
		l.l.Printf(format, args...)
	}
}

func (l *Logger) Infoln(args ...interface{}) {
	if l.Verbosity() >= Info {
		l.l.Println(args...)
	}
}
