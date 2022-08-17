package queue

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/Azure/azure-storage-queue-go/azqueue"
	"github.com/eugbyte/monorepo/libs/formats"
)

type QueueServicer interface {
	QueueExist() bool
	// metaData argument is optional
	CreateQueue(metaData azqueue.Metadata) error
	Enqueue(messageText string, visibilityTimeout time.Duration, timeToLive time.Duration) (*azqueue.EnqueueMessageResponse, error)
}

type queueService struct {
	queueUrl azqueue.QueueURL
	cxt      context.Context
}

// Get the baseConnectionString from config.QueueBaseURL()
func New(cxt context.Context, queueName string, baseConnectionString string, accountName string, accountKey string) QueueServicer {
	qService := queueService{}
	qService.cxt = cxt
	// http://localhost/devstoreaccount1/my-queue
	connection := fmt.Sprintf("%s/%s", baseConnectionString, queueName)
	formats.Trace("connectionString:", connection)
	urlObj, err := url.Parse(connection)
	if err != nil {
		log.Fatal("Error parsing url: ", err)
	}

	credential, err := azqueue.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		log.Fatal("Error creating credentials: ", err)
	}
	var queueUrl azqueue.QueueURL = azqueue.NewQueueURL(*urlObj, azqueue.NewPipeline(credential, azqueue.PipelineOptions{}))
	qService.queueUrl = queueUrl
	formats.Trace("QueueUrl", qService.queueUrl.URL())
	return &qService
}

func (qService *queueService) QueueExist() bool {
	_, err := qService.queueUrl.GetProperties(qService.cxt)
	return err == nil
}

func (qService *queueService) CreateQueue(metaData azqueue.Metadata) error {
	if metaData == nil {
		metaData = azqueue.Metadata{}
	}
	_, err := qService.queueUrl.Create(qService.cxt, metaData)
	return err
}

func (qService *queueService) Enqueue(messageText string, visibilityTimeout time.Duration, timeToLive time.Duration) (*azqueue.EnqueueMessageResponse, error) {
	messageUrl := qService.queueUrl.NewMessagesURL()
	return messageUrl.Enqueue(qService.cxt, messageText, visibilityTimeout, timeToLive)
}
