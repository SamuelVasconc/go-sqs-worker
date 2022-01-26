package usecases

import (
	"encoding/json"
	"log"
	"time"

	"github.com/SamuelVasconc/go-sqs-worker/interfaces"
	"github.com/SamuelVasconc/go-sqs-worker/models"
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
		log.Println("[worker/Execute] Error to formate initial Date. Error: ", err.Error())
		return err
	}

	var currentDate time.Time
	if parameters.FinalDate != "" {
		currentDate, err = time.Parse(dateFormat, parameters.FinalDate)
		if err != nil {
			log.Println("[worker/Execute] Error to formate Final Date. Error: ", err.Error())
			return err
		}
	} else {
		currentDate, err = time.Parse(dateFormat, time.Now().Format(dateFormat))
		if err != nil {
			log.Println("[worker/Execute] Error to get current Date. Error: ", err.Error())
			return err
		}
	}

	tomorrow := currentDate.AddDate(0, 0, 1)

	for d := startDate; d.Before(tomorrow); d = d.AddDate(0, 0, 1) {

		moreLines := true

		for moreLines {

			protocol, err := m.SetProtocol(parameters.Status, d.Format(dateFormat), parameters.Limit)
			if err != nil {
				log.Println("[worker/Execute] Error to set Protocol. Error: ", err.Error())
				return err
			}

			lines, err := m.mysqlRepository.GetLines(protocol)
			if err != nil {
				log.Println("[worker/Execute] Error to get lines of movimentacao_caixa. Error: ", err.Error())
				return err
			}

			if len(lines) <= 0 {
				moreLines = false
				continue
			}

			for _, i := range lines {

				jsonMsg, _ := json.Marshal(i)
				err := m.sqsRepository.PublishMessage(parameters.QueueURL, string(jsonMsg))

				if err != nil {
					log.Println("[worker/Execute] Error to publish message. Error: ", err.Error())
					m.mysqlRepository.UpdateLine("error when trying to publish message", parameters.ErrorStatus, i.ID)
					return err
				}

				err = m.mysqlRepository.UpdateLine("message succefully published", parameters.SuccessStatus, i.ID)
				if err != nil {
					log.Println("[worker/Execute] Error to update line of movimentacao_caixa. Error: ", err.Error())
					return err
				}
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
