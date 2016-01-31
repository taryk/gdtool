package core

import (
	"testing"
	"os"
	"flag"
)

func TestMain(m *testing.M) {
	IsTesting = true
	flag.Parse()
	InitLoggers("debug", "warn", "error")
	os.Exit(m.Run())
}
