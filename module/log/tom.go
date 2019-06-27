package log

import (
	"bytes"
	"io"
	"sync"
	"time"

	"fmt"
)

type Lvl uint8

const (
	DEBUG Lvl = iota + 1
	INFO
)

type TomLogger struct {
	Writer io.Writer
	mutex  sync.Mutex
	levels []string
}

func NewTomLogger() (logger *TomLogger) {
	logger = new(TomLogger)
	logger.levels = []string{
		"-",
		"DEBUG",
		"INFO",
	}
	return
}

func (p *TomLogger) Output() (w io.Writer) {
	return
}
func (p *TomLogger) SetOutput(w io.Writer) {
	p.Writer = w
}
func (p *TomLogger) Prefix() (prefix string) {
	return
}
func (p *TomLogger) SetPrefix(prefix string)                   {}
func (p *TomLogger) SetHeader(h string)                        {}
func (p *TomLogger) Print(i ...interface{})                    {}
func (p *TomLogger) Printf(format string, args ...interface{}) {}
func (p *TomLogger) Debug(i ...interface{})                    {}
func (p *TomLogger) Info(i ...interface{}) {
	p.log(INFO, i)
}
func (p *TomLogger) Warn(i ...interface{})                     {}
func (p *TomLogger) Error(i ...interface{})                    {}
func (p *TomLogger) Fatal(i ...interface{})                    {}
func (p *TomLogger) Fatalf(format string, args ...interface{}) {}
func (p *TomLogger) Panic(i ...interface{})                    {}

func (p *TomLogger) log(level Lvl, args ...interface{}) {
	buf := bytes.Buffer{}

	buf.Write([]byte(time.Now().Format(time.RFC3339)))
	buf.WriteByte(' ')
	buf.Write([]byte(p.levels[level]))
	buf.WriteByte(' ')
	buf.WriteString(fmt.Sprint(args...))
	buf.WriteByte('\n')

	p.mutex.Lock()
	defer p.mutex.Unlock()
	_, _ = p.Writer.Write(buf.Bytes())
}
