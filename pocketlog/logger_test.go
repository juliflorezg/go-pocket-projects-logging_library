package pocketlog_test

import (
	"github.com/juliflorezg/go-pocket-projects-logging_library/pocketlog"
)

func ExampleLogger_Debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, nil)
	debugLogger.Debugf("Hello, %s", "world")
	// Output: Hello, world
}
