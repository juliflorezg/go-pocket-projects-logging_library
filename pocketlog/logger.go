package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information
type Logger struct {
	threshold Level
	output    io.Writer
}

// New returns a Logger, ready to log at the required threshold.
// The default output is Stdout.
func New(threshold Level, output io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		output:    output,
	}
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	// making sure we can safely write to the output
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold <= LevelDebug {
		_, _ = fmt.Fprintf(l.output, format+"\n", args...)
	}
}

// Infof formats and prints a message if the log level is info or higher.
func (l *Logger) Infof(format string, args ...any) {
	// making sure we can safely write to the output
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold <= LevelDebug {
		_, _ = fmt.Fprintf(l.output, format+"\n", args...)
	}
}

// Errorf formats and prints a message if the log level is error or higher.
func (l *Logger) Errorf(format string, args ...any) {
	// making sure we can safely write to the output
	if l.output == nil {
		l.output = os.Stderr
	}

	if l.threshold <= LevelDebug {
		_, _ = fmt.Fprintf(l.output, format+"\n", args...)
	}
}
