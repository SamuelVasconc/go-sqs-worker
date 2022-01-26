package usecases_test

import (
	"errors"
	"testing"

	"github.com/SamuelVasconc/go-sqs-worker/repositories/mocks"
	"github.com/SamuelVasconc/go-sqs-worker/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSetProtocol(t *testing.T) {
	mockRepo := new(mocks.MysqlRepository)
	mockRepoSQS := new(mocks.SQSRepository)

	t.Run("error-set-protocol", func(t *testing.T) {
		mockRepo.On("GenerateProtocol").Return("", nil).Once()
		mockRepo.On("SetProtocol", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(errors.New("")).Once()

		u := usecases.NewMySqlUseCase(mockRepo, mockRepoSQS)
		_, err := u.SetProtocol("", "", 0)
		assert.Error(t, err)
	})

	t.Run("error-generate", func(t *testing.T) {
		mockRepo.On("GenerateProtocol").Return("", errors.New("")).Once()
		mockRepo.On("SetProtocol", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(nil).Once()

		u := usecases.NewMySqlUseCase(mockRepo, mockRepoSQS)
		_, err := u.SetProtocol("", "", 0)
		assert.Error(t, err)
	})

	t.Run("success", func(t *testing.T) {
		mockRepo.On("GenerateProtocol").Return("", nil)
		mockRepo.On("SetProtocol", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("int")).Return(nil)

		u := usecases.NewMySqlUseCase(mockRepo, mockRepoSQS)
		_, err := u.SetProtocol("", "", 0)
		assert.NoError(t, err)
	})

}
