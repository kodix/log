package log

import (
	"log"
	"os"
	"runtime"
	"testing"
)

func TestRace(t *testing.T) {
	out, err := os.Open(os.DevNull)
	if err != nil {
		t.Fatal(err)
	}
	SetOutput(out)
	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			SetVerbosity(Debug)
			Debugln("")
			Debugf("")
			SetVerbosity(Error)
			SetVerbosity(Warning)
			Warningln("")
			Warningf("")
			_ = std.Verbosity()
			SetOutput(out)
			Errorln("")
			Errorf("")
			SetPrefix("")
			SetVerbosity(Info)
			SetFlags(log.LstdFlags | log.Lshortfile)
			Infoln("")
			Infof("")
		}()
	}
}

func TestLogger_SetVerbosity(t *testing.T) {
	var name = "Verbosity setter"
	t.Run(name, func(t *testing.T) {
		l := New(os.Stdout, "", Error, log.LstdFlags)
		l.SetVerbosity(Info)
		if l.Verbosity() != Info {
			t.Errorf("%s: want: %v, got: %v", name, Info, l.Verbosity())
		}
	})
}

func TestLogger_Fatalln(t *testing.T) {
	Fatalln("i am fatal", "blablabla")
}
