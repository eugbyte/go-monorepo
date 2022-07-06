package queue

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/Azure/azure-storage-queue-go/azqueue"
	"github.com/web-notify/api/monorepo/libs/utils/logs"
)

type QueueService struct {
	queueUrl azqueue.QueueURL
	cxt      context.Context
}

type QueueServiceImpl interface {
	Init(cxt context.Context, queueName string, accountName string, accountKey string, connectionString string)
	QueueExist() bool
	CreateQueue(metaData azqueue.Metadata) error
	Enqueue(messageText string, visibilityTimeout time.Duration, timeToLive time.Duration) (*azqueue.EnqueueMessageResponse, error)
}

/*
	Must use this method to initialise the queue. Each queue is unique to the queue name
	To change to a different queue, must call this method again with a different queueName
*/
func (qService *QueueService) Init(cxt context.Context, queueName string, accountName string, accountKey string, connectionString string) {
	qService.cxt = cxt
	// http://localhost/devstoreaccount1/my-queue
	qService.queueUrl = GenQueueUrl(queueName, accountName, accountKey, connectionString)
	logs.Trace("QueueUrl", qService.queueUrl.URL())
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
