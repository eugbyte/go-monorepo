package queue

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/Azure/azure-storage-queue-go/azqueue"
)

type QueueService struct {
	QueueUrl azqueue.QueueURL
	CXT      context.Context
}

type QueueServiceImpl interface {
	Init(queueName string, cxt context.Context, accountName string, accountKey string, connectionString string)
	IsQueueExist() bool
	CreateQueue() error
	GetMessageURL() azqueue.MessagesURL
}

/*
	Must use this method to initialise the queue client as a state
*/
func (qService *QueueService) Init(queueName string, cxt context.Context, accountName string, accountKey string, connectionString string) {
	qService.CXT = cxt
	qService.QueueUrl = GenQueueUrl(queueName, accountName, accountKey, connectionString)
}

// To generate the QueueURL to the queue
// connection string: "http://127.0.0.1:10001/devstoreaccount1/{queueName}"
func GenQueueUrl(queueName string, accountName string, accountKey string, connection string) azqueue.QueueURL {

	_url, err := url.Parse(fmt.Sprintf("%s/%s", connection, queueName))
	if err != nil {
		log.Fatal("Error parsing url: ", err)
	}

	credential, err := azqueue.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal("Error creating credentials: ", err)
	}

	queueUrl := azqueue.NewQueueURL(*_url, azqueue.NewPipeline(credential, azqueue.PipelineOptions{}))
	return queueUrl
}

func (qService *QueueService) IsQueueExist() bool {
	_, err := qService.QueueUrl.GetProperties(qService.CXT)
	return err == nil
}

func (qService *QueueService) CreateQueue() error {
	ctx := context.Background()
	_, err := qService.QueueUrl.Create(ctx, azqueue.Metadata{})

	return err
}

func (qService QueueService) GetMessageURL() azqueue.MessagesURL {
	return qService.QueueUrl.NewMessagesURL()
}
