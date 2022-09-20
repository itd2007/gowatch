package parser

import (
	"github.com/itd2007/gowatch/logentry"
	"testing"
)

func acceptAllPredicate() logentry.AcceptAllPredicate {
	return logentry.AcceptAllPredicate{}
}

func givenLineSource(t *testing.T, lines ...string) LineSource {
	linesource := NewSimpleLineSource()
	for _, line := range lines {
		linesource.AddLine(line)
	}
	return linesource
}
