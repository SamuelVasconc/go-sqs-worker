package usecases

import (
	"encoding/json"
	"time"

	"github.com/SamuelVasconc/go-sqs-worker/interfaces"
	"github.com/SamuelVasconc/go-sqs-worker/models"
	"github.com/SamuelVasconc/go-sqs-worker/utils/logger"
)

const dateFormat = "2006-01-02"

type mysqlUseCase struct {
	mysqlRepository interfaces.MySqlRepository
	sqsRepository   interfaces.SqsRepository
}

func NewMySqlUseCase(mysqlRepository interfaces.MySqlRepository, sqsRepository interfaces.SqsRepository) interfaces.MySqlUseCase {
	return &mysqlUseCase{mysqlRepository, sqsRepository}
}

func (m *mysqlUseCase) GetLines(parameters models.Parameter) error {

	startDate, err := time.Parse(dateFormat, parameters.StartDate)
	if err != nil {
		logger.Error("Error to formate initial Date. Error:", err.Error())
		return err
	}

	var currentDate time.Time
	if parameters.FinalDate != "" {
		currentDate, err = time.Parse(dateFormat, parameters.FinalDate)
		if err != nil {
			logger.Error("Error to formate Final Date. Error:", err.Error())
			return err
		}
	} else {
		currentDate, err = time.Parse(dateFormat, time.Now().Format(dateFormat))
		if err != nil {
			logger.Error("Error to get current Date. Error:", err.Error())
			return err
		}
	}

	tomorrow := currentDate.AddDate(0, 0, 1)

	for d := startDate; d.Before(tomorrow); d = d.AddDate(0, 0, 1) {

		logger.Debug("searching lines for day:", d.Format(dateFormat))

		moreLines := true

		for moreLines {

			protocol, err := m.SetProtocol(parameters.Status, d.Format(dateFormat), parameters.Limit)
			if err != nil {
				logger.Error("Error to set Protocol. Error:", err.Error())
				return err
			}
			logger.Debug("Getting protocol", protocol)

			lines, err := m.mysqlRepository.GetLines(protocol)
			if err != nil {
				logger.Error("Error to get lines of t_transactions. Error:", err.Error())
				return err
			}
			logger.Debug("Lines Gotted", len(lines))

			if len(lines) <= 0 {
				logger.Debug("There isn't more lines")
				moreLines = false
				continue
			}

			for _, i := range lines {

				jsonMsg, _ := json.Marshal(i)

				logger.Debug("Publishing message id:", i.ID)
				err := m.sqsRepository.PublishMessage(parameters.QueueURL, string(jsonMsg))

				if err != nil {
					logger.Error("Error to publish message. Error:", err.Error())
					m.mysqlRepository.UpdateLine("error when trying to publish message", parameters.ErrorStatus, i.ID)
					return err
				}

				logger.Debug("Updating message id:", i.ID)
				err = m.mysqlRepository.UpdateLine("message succefully published", parameters.SuccessStatus, i.ID)
				if err != nil {
					logger.Error("Error to update line of t_transactions. Error:", err.Error())
					return err
				}
				logger.Debug("Successfully processed message id:", i.ID)
			}
		}
	}

	return nil
}

func (m *mysqlUseCase) SetProtocol(status, date string, limit int) (string, error) {
	protocol, err := m.mysqlRepository.GenerateProtocol()
	if err != nil {
		return "", err
	}

	err = m.mysqlRepository.SetProtocol(status, date, protocol, limit)
	if err != nil {
		return "", err
	}

	return protocol, nil
}
