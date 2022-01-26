package models

type Parameter struct {
	StartDate     string
	FinalDate     string
	Status        string
	Limit         int
	QueueURL      string
	SuccessStatus string
	ErrorStatus   string
	SleepTime     int
}
