package logsWriter

import "gift-buyer/internal/infrastructure/logsWriter/logTypes"

type LogsWriter interface {
	Write(entry *logTypes.LogEntry) error
}
