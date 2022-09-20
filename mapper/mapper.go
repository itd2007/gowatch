package mapper

import (
	"github.com/itd2007/gowatch/logentry"
)

type Mapper interface {
	Map(entries <-chan logentry.LogEntry) <-chan logentry.LogEntry
}
