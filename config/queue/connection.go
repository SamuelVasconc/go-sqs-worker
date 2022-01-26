package queue

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

//InitSQSQueue ...
func InitSQSQueue(URL, region, ID, secret string) (*sqs.SQS, error) {
	if URL == "" {
		URL = os.Getenv("SQS_URL")
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(ID, secret, ""),
	})

	if err != nil {
		return nil, err
	}

	svc := sqs.New(sess, aws.NewConfig().WithEndpoint(URL))

	return svc, nil
}
