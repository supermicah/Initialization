package mlog

import (
	"runtime"
)

// The plog is a public function combiation for other log objects.
type plog struct {
	isOpen bool
}

// GetCaller returns file name and line number at the third step of runtime.
func (*plog) getCaller() (string, int) {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "???"
		line = 0
	}
	return file, line
}

// Open makes log open.
func (p *plog) Open() {
	p.isOpen = true
}

// Close makes log close.
func (p *plog) Close() {
	p.isOpen = false
}
