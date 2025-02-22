package config

import (
	"fmt"
	"github.com/itd2007/gowatch/logentry"
	"github.com/itd2007/gowatch/parser"
	"log"
)

func (logfile *LogfileConfig) CreateParser(linesource parser.LineSource, predicate logentry.Predicate) parser.Parser {
	timeLayout := parseTimeLayout(logfile.TimeLayout)

	switch logfile.Parser {
	case "simple":
		return parser.NewSimpleParser(linesource, predicate)
	case "", "grok":
		if pattern, ok := logfile.With["pattern"]; ok {
			return parser.NewGrokParser(linesource, fmt.Sprint(pattern), timeLayout, predicate)
		}
		log.Fatal("Grok parser used without pattern on logfile '", logfile.Filename, "'")
		return nil // actually never reached
	default:
		log.Fatal("Unrecognized parser '", logfile.Parser, "' on logfile '", logfile.Filename, "'")
		return nil // actually never reached
	}
}

func parseTimeLayout(givenTimeLayout string) string {
	if interpretedTimeLayout, ok := PredefinedTimeLayouts[givenTimeLayout]; ok {
		return interpretedTimeLayout
	}
	return givenTimeLayout
}
