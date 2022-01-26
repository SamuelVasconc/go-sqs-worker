package utils_test

import (
	"os"
	"testing"

	"github.com/SamuelVasconc/go-sqs-worker/utils"
	"github.com/stretchr/testify/assert"
)

func TestPrepareParameters(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		os.Setenv("STARTDATE", "2022-01-18")
		os.Setenv("FINALDATE", "2022-01-21")
		os.Setenv("STATUS", "N")
		os.Setenv("QUEUEURL", "https://sqs.")
		os.Setenv("SUCCESSSTATUS", "P")
		os.Setenv("ERRORSTATUS", "E")
		os.Setenv("LIMIT", "1000")
		os.Setenv("SLEEPTIME", "30")
		os.Setenv("LOGGING_LEVEL", "DEBUG")

		_, err := utils.PrepareParameters()
		assert.NoError(t, err)
	})

	t.Run("error limit", func(t *testing.T) {
		os.Setenv("LIMIT", "fjgdjfg")

		_, err := utils.PrepareParameters()
		assert.Error(t, err)
	})

	t.Run("error sleep", func(t *testing.T) {
		os.Setenv("LIMIT", "30")
		os.Setenv("SLEEPTIME", "fgfjgh")

		_, err := utils.PrepareParameters()
		assert.Error(t, err)
	})
}
