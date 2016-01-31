package cache

import (
	"testing"
	"os"
	"flag"

	. "github.com/taryk/gdtool/core"
)

func TestMain(m *testing.M) {
	IsTesting = true
	flag.Parse()
	InitLoggers("debug", "warn", "error")
	os.Exit(m.Run())
}
