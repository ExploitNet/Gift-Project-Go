package authService

import (
	"fmt"
	"gift-buyer/internal/infrastructure/logsWriter/logTypes"
	"gift-buyer/pkg/logger"
)

// Helper methods for clean logging
func (f *AuthManagerImpl) logInfo(message string) {
	if err := f.infoLogsWriter.Write(&logTypes.LogEntry{
		Message: message,
	}); err != nil {
		logger.GlobalLogger.Errorf("Failed to write info log: %v", err)
	}
}

func (f *AuthManagerImpl) logError(message string) {
	if err := f.errorLogsWriter.Write(&logTypes.LogEntry{
		Message: message,
	}); err != nil {
		logger.GlobalLogger.Errorf("Failed to write error log: %v", err)
	}
}

func (f *AuthManagerImpl) logErrorf(format string, args ...interface{}) {
	f.logError(fmt.Sprintf(format, args...))
}
