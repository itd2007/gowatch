package parser

import (
	"github.com/itd2007/gowatch/logentry"
)

// converts one line into one message, so it doesn't actually parse something
type SimpleParser struct {
	linesource LineSource
	predicate  logentry.Predicate
}

func NewSimpleParser(linesource LineSource, predicate logentry.Predicate) (p *SimpleParser) {
	p = new(SimpleParser)

	p.linesource = linesource
	p.predicate = predicate

	return
}

func (p *SimpleParser) Parse() <-chan logentry.LogEntry {
	return parse(p.linesource, p.predicate, p.lineToLogEntry)
}

func (p *SimpleParser) lineToLogEntry(line string, entry *logentry.LogEntry) {
	entry.Message = line
}
