package log

import (
	"io"
)

type (
	Logger interface {
		Output() io.Writer
		SetOutput(w io.Writer)
		Prefix() string
		SetPrefix(p string)
		SetHeader(h string)
		Print(i ...interface{})
		Printf(format string, args ...interface{})
		Debug(i ...interface{})
		Info(i ...interface{})
		Warn(i ...interface{})
		Error(i ...interface{})
		Fatal(i ...interface{})
		Fatalf(format string, args ...interface{})
		Panic(i ...interface{})
	}
)
