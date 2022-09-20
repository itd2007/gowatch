package config

import (
	"github.com/itd2007/gowatch/logentry"
	"github.com/itd2007/gowatch/parser"
	"testing"
)

func acceptAllPredicate() logentry.AcceptAllPredicate {
	return logentry.AcceptAllPredicate{}
}

func givenLineSource(t *testing.T, lines ...string) parser.LineSource {
	linesource := parser.NewSimpleLineSource()
	for _, line := range lines {
		linesource.AddLine(line)
	}
	return linesource
}
