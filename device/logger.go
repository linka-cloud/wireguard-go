/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2017-2021 WireGuard LLC. All Rights Reserved.
 */

package device

import (
	"log"
	"os"
)

// A Logger provides logging for a Device.
// The functions are Printf-style functions.
// They must be safe for concurrent use.
// They do not require a trailing newline in the format.
// If nil, that level of logging will be silent.
type Logger interface {
	Verbosef(format string, args ...any)
	Errorf(format string, args ...any)
}

type logger struct {
	verbosef func(format string, args ...any)
	errorf   func(format string, args ...any)
}

func (l *logger) Verbosef(format string, args ...any) {
	if l.verbosef != nil {
		l.verbosef(format, args...)
	}
}
func (l *logger) Errorf(format string, args ...any) {
	if l.errorf != nil {
		l.errorf(format, args...)
	}
}

// Log levels for use with NewLogger.
const (
	LogLevelSilent = iota
	LogLevelError
	LogLevelVerbose
)

// Function for use in Logger for discarding logged lines.
func DiscardLogf(format string, args ...any) {}

// NewLogger constructs a Logger that writes to stdout.
// It logs at the specified log level and above.
// It decorates log lines with the log level, date, time, and prepend.
func NewLogger(level int, prepend string) Logger {
	logger := &logger{}
	logf := func(prefix string) func(string, ...any) {
		return log.New(os.Stdout, prefix+": "+prepend, log.Ldate|log.Ltime).Printf
	}
	if level >= LogLevelVerbose {
		logger.verbosef = logf("DEBUG")
	}
	if level >= LogLevelError {
		logger.errorf = logf("ERROR")
	}
	return logger
}
