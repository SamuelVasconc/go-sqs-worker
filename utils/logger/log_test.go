package logger_test

import (
	"errors"
	"testing"

	"github.com/SamuelVasconc/go-sqs-worker/utils/logger"
	"github.com/stretchr/testify/assert"
)

func TestDebug(t *testing.T) {
	t.Run("test log a debug message", func(t *testing.T) {
		logger.Level = logger.Enum("DEBUG")
		logger.Debug("Debug message")

		assert.Equal(t, logger.Level, logger.Severity("DEBUG"))
	})
}

func TestInfo(t *testing.T) {
	t.Run("test log a info message", func(t *testing.T) {
		logger.Level = logger.Enum("INFO")
		logger.Info("Info message")

		assert.Equal(t, logger.Level, logger.Severity("INFO"))
	})
}

func TestWarn(t *testing.T) {
	t.Run("test log a warn message", func(t *testing.T) {
		logger.Level = logger.Enum("WARN")
		logger.Warn("Warn message")

		assert.Equal(t, logger.Level, logger.Severity("WARN"))
	})
}

func TestError(t *testing.T) {
	logger.Level = logger.Enum("ERROR")
	t.Run("test log a error message", func(t *testing.T) {
		err := errors.New("error message")
		logger.Error(err)

		assert.Equal(t, logger.Level, logger.Severity("ERROR"))
	})

	t.Run("test should not log when error is nil", func(t *testing.T) {
		var err error = nil
		logger.Error(err)

		assert.Equal(t, logger.Level, logger.Severity("ERROR"))
	})

	t.Run("test should log multiple messages", func(t *testing.T) {
		message := "message1"
		message2 := "{error}"

		logger.Error(message, message2)

		assert.Equal(t, logger.Level, logger.Severity("ERROR"))
	})
}
