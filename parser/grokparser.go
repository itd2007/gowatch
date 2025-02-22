package parser

import (
	"github.com/itd2007/gowatch/logentry"
	"github.com/vjeantet/grok"
	"log"
	"time"
)

type GrokParser struct {
	linesource LineSource
	grok       *grok.Grok
	pattern    string
	predicate  logentry.Predicate
	timeLayout string
}

func NewGrokParser(linesource LineSource, pattern string, timeLayout string, predicate logentry.Predicate) (p *GrokParser) {
	return &GrokParser{
		grok:       grok.New(),
		linesource: linesource,
		pattern:    pattern,
		timeLayout: timeLayout,
		predicate:  predicate}
}

func (p *GrokParser) AddPattern(name string, pattern string) {
	p.grok.AddPattern(name, pattern)
}

func (p *GrokParser) Parse() <-chan logentry.LogEntry {
	return parse(p.linesource, p.predicate, p.lineToLogEntry)
}

func (p *GrokParser) lineToLogEntry(line string, entry *logentry.LogEntry) {
	matches, err := p.grok.ParseToMultiMap(p.pattern, line)
	if err != nil {
		log.Print(err)
		return
	}

	for field, values := range matches {
		for _, value := range values {
			if entry.IsTimestamp(field) {
				timestamp, err := time.Parse(p.timeLayout, value)
				if err != nil {
					log.Fatalf("Error on parsing with time layout \"%s\": %s", p.timeLayout, err)
				}
				if timestamp.Year() == 0 {
					timestamp = timestamp.AddDate(time.Now().Year(), 0, 0)
					if time.Now().Before(timestamp) {
						timestamp = timestamp.AddDate(-1, 0, 0) // minus one year
					}
				}
				entry.Timestamp = timestamp
			} else if err := entry.AssignValue(field, value); err != nil {
				log.Fatal(err)
			}
		}
	}
}
