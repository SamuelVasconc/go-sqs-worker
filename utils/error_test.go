package utils_test

import (
	"testing"

	"github.com/SamuelVasconc/go-sqs-worker/utils"
	"github.com/stretchr/testify/assert"
)

func TestHandleError(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var errorIn error
		err := utils.HandleError(errorIn)
		assert.Error(t, err)
	})
}
