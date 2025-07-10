package giftServiceHelpers

import (
	"fmt"
	"gift-buyer/internal/infrastructure/logsWriter"
	"gift-buyer/internal/infrastructure/logsWriter/logTypes"
	"gift-buyer/pkg/logger"
)

// Helper methods for clean logging
func LogInfo(infoLogsWriter logsWriter.LogsWriter, message string) {
	if err := infoLogsWriter.Write(&logTypes.LogEntry{
		Message: message,
	}); err != nil {
		logger.GlobalLogger.Errorf("Failed to write info log: %v", err)
	}
}

func LogError(errorLogsWriter logsWriter.LogsWriter, message string) {
	if err := errorLogsWriter.Write(&logTypes.LogEntry{
		Message: message,
	}); err != nil {
		logger.GlobalLogger.Errorf("Failed to write error log: %v", err)
	}
}

func LogErrorf(errorLogsWriter logsWriter.LogsWriter, format string, args ...interface{}) {
	LogError(errorLogsWriter, fmt.Sprintf(format, args...))
}
