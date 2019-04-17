package log

import (
	"io"
	"log"
	"os"
	"sync"
)

type Verbosity uint8

const (
	Error Verbosity = iota
	Warning
	Info
	Debug
)

type Logger struct {
	l *log.Logger

	mu  sync.RWMutex
	lvl Verbosity
}

func New(out io.Writer, prefix string, v Verbosity, flag int) *Logger {
	return &Logger{
		l:   log.New(out, prefix, flag),
		mu:  sync.RWMutex{},
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
	l.l.Fatalf("[fatal]"+format, args...)
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.l.Fatalln(append([]interface{}{"[fatal]"}, args...)...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.l.Panicf("[panic]"+format, args...)
}

func (l *Logger) Panicln(args ...interface{}) {
	l.l.Panicln(append([]interface{}{"[panic]"}, args...)...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.l.Printf("[error]"+format, args...)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.l.Println(append([]interface{}{"[error]"}, args...)...)
}

func (l *Logger) Warningf(format string, args ...interface{}) {
	if l.Verbosity() >= Warning {
		l.l.Printf("[warning]"+format, args...)
	}
}

func (l *Logger) Warningln(args ...interface{}) {
	if l.Verbosity() >= Warning {
		l.l.Println(append([]interface{}{"[warning]"}, args...)...)
	}
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.Verbosity() >= Debug {
		l.l.Printf("[debug]"+format, args...)
	}
}

func (l *Logger) Debugln(args ...interface{}) {
	if l.Verbosity() >= Debug {
		l.l.Println(append([]interface{}{"[debug]"}, args...)...)
	}
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if l.Verbosity() >= Info {
		l.l.Printf("[info]"+format, args...)
	}
}

func (l *Logger) Infoln(args ...interface{}) {
	if l.Verbosity() >= Info {
		l.l.Println(append([]interface{}{"[info]"}, args...)...)
	}
}

func (l *Logger) Prefix() string {
	return l.l.Prefix()
}

func (l *Logger) AddPrefix(prefix string) {
	l.SetPrefix(l.Prefix() + " " + prefix)
}

func (l *Logger) Copy() *Logger {
	logger := log.New(os.Stdout, l.l.Prefix(), l.l.Flags())
	return &Logger{
		l:   logger,
		lvl: l.lvl,
		mu:  sync.RWMutex{},
	}
}
