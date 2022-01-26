package utils

import (
	"os"
	"strconv"

	"github.com/SamuelVasconc/go-sqs-worker/models"
)

// PrepareParameters ...
func PrepareParameters() (models.Parameter, error) {
	parameters := models.Parameter{
		StartDate:     os.Getenv("STARTDATE"),
		FinalDate:     os.Getenv("FINALDATE"),
		Status:        os.Getenv("STATUS"),
		QueueURL:      os.Getenv("QUEUEURL"),
		SuccessStatus: os.Getenv("SUCCESSSTATUS"),
		ErrorStatus:   os.Getenv("ERRORSTATUS"),
	}

	limit, err := strconv.Atoi(os.Getenv("LIMIT"))
	if err != nil {
		return parameters, err
	}

	sleep, err := strconv.Atoi(os.Getenv("SLEEPTIME"))
	if err != nil {
		return parameters, err
	}

	parameters.Limit = limit
	parameters.SleepTime = sleep
	return parameters, nil
}
