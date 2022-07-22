package queue

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/Azure/azure-storage-queue-go/azqueue"
	"github.com/web-notify/api/monorepo/libs/utils/config"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
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

/*
	Must use this method to initialise the queue. Each queue is unique to the queue name
	To change to a different queue, must call this method again with a different queueName
*/
func NewQueueService(cxt context.Context, queueName string, baseConnectionString string, accountName string, accountKey string) QueueServicer {
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

// To generate the QueueURL to the queue
// connection string: "http://127.0.0.1:10001/devstoreaccount1/{queueName}"
func GetBaseConnectionString(stage string, accountName string) string {
	if stage == config.DEV {
		return fmt.Sprintf("%s/%s", "http://127.0.0.1:10001", accountName)
	} else {
		return fmt.Sprintf("https://%s.queue.core.windows.net", accountName)
	}
}
