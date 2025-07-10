package logsWriter

import (
	"encoding/json"
	"errors"
	"fmt"
	"gift-buyer/internal/infrastructure/logsWriter/logTypes"
	"gift-buyer/pkg/logger"
	"os"
	"sync"
	"time"
)

type logsWriterImpl struct {
	File  *os.File
	mu    sync.Mutex
	level string
}

func NewLogsWriter(level string) *logsWriterImpl {
	file, err := os.OpenFile(fmt.Sprintf("%s_logs.jsonl", level), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		logger.GlobalLogger.Fatalf("Failed to open log file: %v", err)
	}

	writer := &logsWriterImpl{
		File:  file,
		level: level,
	}

	return writer
}

func (l *logsWriterImpl) Write(entry *logTypes.LogEntry) (err error) {
	bytes, err := l.marshalLogEntry(entry)
	if err != nil {
		return errors.New("failed to marshal to json")
	}

	if err := l.write(bytes); err != nil {
		return errors.New("failed to write logs to file")
	}

	return nil
}

func (l *logsWriterImpl) marshalLogEntry(entry *logTypes.LogEntry) ([]byte, error) {
	entryCopy := *entry
	entryCopy.Timestamp = time.Now().Format(time.RFC3339)

	if entryCopy.Level == "" {
		entryCopy.Level = l.level
	}

	jsonBytes, err := json.Marshal(entryCopy)
	if err != nil {
		return nil, err
	}

	jsonBytes = append(jsonBytes, '\n')
	return jsonBytes, nil
}

func (l *logsWriterImpl) write(bytes []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	_, err := l.File.Write(bytes)
	if err != nil {
		return err
	}
	return nil
}
