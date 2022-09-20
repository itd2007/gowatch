package summary

import "github.com/itd2007/gowatch/logentry"

func givenEntriesWithMessages(messages ...string) <-chan logentry.LogEntry {
	entries := make(chan logentry.LogEntry, len(messages))
	for _, message := range messages {
		entry := logentry.New()
		entry.Message = message
		entries <- *entry
	}
	close(entries)
	return entries
}
