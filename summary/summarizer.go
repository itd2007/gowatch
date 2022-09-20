package summary

import "github.com/itd2007/gowatch/logentry"

type Summarizer interface {
	SummarizeAsync(entries <-chan logentry.LogEntry)
	StringAfterSummarizeAsyncCompleted() string
}
