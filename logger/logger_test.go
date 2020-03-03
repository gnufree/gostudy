package logger

import (
	"testing"
)

func TestFileLogger(t *testing.T) {
	logger := NewFileLogger(LogLevelDebug, "./", "test")
	logger.Debug("user id[%d] is come form china", 32121)
	logger.Trace("Trace id [%d] is come form china", 32121)
	logger.Info("info log [%d]", 32121)
	logger.Warn("test warn log")
	logger.Error("test error log")
	logger.Fatal("test fatal log")
	logger.Close()
}

func TestConsoleLogger(t *testing.T) {
	logger := NewConsoleLogger(LogLevelDebug)
	logger.Debug("user id[%d] is come form china", 32121)
	logger.Trace("Trace id [%d] is come form china", 32121)
	logger.Info("info log [%d]", 32121)
	logger.Warn("test warn log")
	logger.Error("test error log")
	logger.Fatal("test fatal log")
	logger.Close()
}
