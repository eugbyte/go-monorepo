package mongo

import (
	"context"
	"log"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MonogoServiceImp interface {
	Init(dbName string, connectionString string)
	Find(collectionName string, filter primitive.D, items []interface{}) error
	InsertOne(collectionName string, item interface{}) error
}

type MongoService struct {
	database *mongo.Database
}

func (ms *MongoService) Init(dbName string, connectionString string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectionString).SetDirect(true)

	client, _ := mongo.NewClient(clientOptions)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatalf("unable to initialize connection %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("unable to connect %v", err)
	}
	ms.database = client.Database(dbName)
}

func (ms *MongoService) Find(collectionName string, filter primitive.D, items []interface{}) error {
	collection := ms.database.Collection(collectionName)
	ctx := context.Background()
	rs, err := collection.Find(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "no item found")
	}
	err = rs.All(ctx, &items)
	if err != nil {
		log.Fatalf("failed to list todo(s) %v", err)
	}
	return err
}

func (ms *MongoService) InsertOne(collectionName string, item interface{}) error {
	ctx := context.Background()
	ms.database.Client().Connect(ctx)
	defer ms.database.Client().Disconnect(ctx)
	collection := ms.database.Collection(collectionName)
	result, err := collection.InsertOne(ctx, item)
	if err != nil {
		log.Fatalf("failed to add todo %v", err)
	}
	log.Print(result)
	return err
}
