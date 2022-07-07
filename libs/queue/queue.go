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

type QueueServiceImpl interface {
	/*
		Must use this method to initialise the queue. Each queue is unique to the queue name
		To change to a different queue, must call this method again with a different queueName
	*/
	Init(cxt context.Context, queueName string, accountName string, accountKey string)
	QueueExist() bool
	// metaData argument is optional
	CreateQueue(metaData azqueue.Metadata) error
	Enqueue(messageText string, visibilityTimeout time.Duration, timeToLive time.Duration) (*azqueue.EnqueueMessageResponse, error)
}

type QueueService struct {
	queueUrl azqueue.QueueURL
	cxt      context.Context
}

func (qService *QueueService) Init(cxt context.Context, queueName string, accountName string, accountKey string) {
	qService.cxt = cxt
	// http://localhost/devstoreaccount1/my-queue
	connection := getConnectionString(accountName, queueName)
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
}

func (qService *QueueService) QueueExist() bool {
	_, err := qService.queueUrl.GetProperties(qService.cxt)
	return err == nil
}

func (qService *QueueService) CreateQueue(metaData azqueue.Metadata) error {
	if metaData == nil {
		metaData = azqueue.Metadata{}
	}
	_, err := qService.queueUrl.Create(qService.cxt, metaData)
	return err
}

func (qService *QueueService) Enqueue(messageText string, visibilityTimeout time.Duration, timeToLive time.Duration) (*azqueue.EnqueueMessageResponse, error) {
	messageUrl := qService.queueUrl.NewMessagesURL()
	return messageUrl.Enqueue(qService.cxt, messageText, visibilityTimeout, timeToLive)
}

// To generate the QueueURL to the queue
// connection string: "http://127.0.0.1:10001/devstoreaccount1/{queueName}"
func getConnectionString(accountName string, queueName string) string {
	if config.STAGE == config.DEV {
		return fmt.Sprintf("%s/%s/%s", "http://127.0.0.1:10001", accountName, queueName)
	} else {
		return fmt.Sprintf("https://%s.queue.core.windows.net/%s", accountName, queueName)
	}
}
