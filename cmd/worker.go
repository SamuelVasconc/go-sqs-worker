package cmd

import (
	"log"
	"os"
	"time"

	"github.com/SamuelVasconc/go-sqs-worker/config/db"
	"github.com/SamuelVasconc/go-sqs-worker/config/queue"
	"github.com/SamuelVasconc/go-sqs-worker/interfaces"
	"github.com/SamuelVasconc/go-sqs-worker/repositories"
	"github.com/SamuelVasconc/go-sqs-worker/usecases"
	"github.com/SamuelVasconc/go-sqs-worker/utils"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Worker struct {
	SQS struct {
		Region string
		ID     string
		Secret string
		Queue  string
		Conn   *sqs.SQS
	}
	SQSUseCase   interfaces.SqsUseCase
	MySqlUseCase interfaces.MySqlUseCase
}

func (w *Worker) Initialization() {
	//Inicialize Database Connection
	db.InitDb()

	//Inicialize SQS Connection
	w.SQS.Region = os.Getenv("SQS_REGION")
	w.SQS.ID = os.Getenv("SQS_ID")
	w.SQS.Secret = os.Getenv("SQS_SECRET")
	w.SQS.Queue = os.Getenv("SQS_URL")

	var err error
	w.SQS.Conn, err = queue.InitSQSQueue("", w.SQS.Region, w.SQS.ID, w.SQS.Secret)
	if err != nil {
		log.Println("[worker/Initialization] Error to connect on SQS queue: ", utils.HandleError(err).Error())
		os.Exit(-1)
	}

	//instances repositories and usecases
	sqsRepository := repositories.NewSqsRepository(w.SQS.Conn)
	w.SQSUseCase = usecases.NewSqsUseCase(sqsRepository)

	mySqlRepository := repositories.NewMySqlRepository(db.DBConn)
	w.MySqlUseCase = usecases.NewMySqlUseCase(mySqlRepository, sqsRepository)
}

func (w *Worker) Execute() {

	for {
		//Load Parameters
		parameters, err := utils.PrepareParameters()
		if err != nil {
			log.Println("[worker/Initialization] Error to connect on SQS queue: ", utils.HandleError(err).Error())
			return
		}

		w.MySqlUseCase.GetLines(parameters)

		//Alternativa para o uso de cronjobs
		//uma vez que o processo principal se encerra o worker dorme por um tempo determinado e retorna a execução atualizando os parametros
		time.Sleep(time.Duration(parameters.SleepTime) * time.Minute)
	}
}
