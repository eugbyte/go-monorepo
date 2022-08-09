package mongolib

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/web-notify/api/monorepo/libs/utils/formats"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MonogoServicer interface {
	DB() *mongo.Database
	CreatedShardedCollection(collectionName string, field string, unique bool)
	CreateIndex(collectionName string, field string, unique bool) error
	Find(collectionName string, filter primitive.D, items []interface{}) error
	FindOne(collectionName string, filter primitive.D, item interface{}) error
	InsertOne(collectionName string, item interface{}) error
	UpdateOne(collectionName string, filter primitive.D, item interface{}, upsert bool) error
}

type mongoService struct {
	Database *mongo.Database
}

func NewMongoService(dbName string, connectionString string) MonogoServicer {
	ms := mongoService{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectionString).SetDirect(true)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("cannot create client")
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("unable to initialize connection %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("unable to connect %v", err)
	}
	ms.Database = client.Database(dbName)
	return &ms
}

func (ms *mongoService) DB() *mongo.Database {
	return ms.Database
}

// Create sharded collection. If sharded collection already exists, operation is skipped
// https://www.mongodb.com/community/forums/t/how-do-you-shard-a-collection-with-the-go-driver/4676
func (ms *mongoService) CreatedShardedCollection(collectionName string, field string, unique bool) {
	ctx := context.Background()

	existingCollectionNames, err := ms.Database.ListCollectionNames(
		context.TODO(),
		bson.D{{Key: "options.capped", Value: true}},
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, coltnName := range existingCollectionNames {
		if coltnName == collectionName {
			formats.Trace("collection already exists, skipping creating sharded collection ...")
			return
		}
	}

	cmd := bson.D{
		{Key: "shardCollection", Value: fmt.Sprintf("%s.%s", ms.Database.Name(), collectionName)},
		{Key: "key", Value: bson.M{field: "hashed"}}, // Hashed sharding requires a field hashed index
		{Key: "unique", Value: unique},
	}
	err = ms.Database.RunCommand(ctx, cmd).Err()

	if err != nil {
		log.Fatalf("sharding failed. %v", err)
	}
}

// From https://christiangiacomi.com/posts/mongodb-index-using-go/
func (ms *mongoService) CreateIndex(collectionName string, field string, unique bool) error {
	indexModel := mongo.IndexModel{
		Keys:    bson.M{field: 1}, // index in ascending order or -1 for descending order
		Options: options.Index().SetUnique(unique),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := ms.Database.Collection(collectionName)

	res, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		fmt.Println(err.Error())
	}
	formats.Trace(res)

	return err
}

func (ms *mongoService) Find(collectionName string, filter primitive.D, items []interface{}) error {
	collection := ms.Database.Collection(collectionName)
	ctx := context.Background()
	rs, err := collection.Find(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "no item found")
	}
	err = rs.All(ctx, &items)
	if err != nil {
		log.Fatalf("failed to list item(s) %v", err)
	}
	return err
}

func (ms *mongoService) FindOne(collectionName string, filter primitive.D, item interface{}) error {
	collection := ms.Database.Collection(collectionName)
	ctx := context.Background()
	rs := collection.FindOne(ctx, filter)
	err := rs.Decode(item)
	if err != nil {
		log.Printf("failed to list item. %v", err)
	}
	return err
}

func (ms *mongoService) InsertOne(collectionName string, item interface{}) error {
	ctx := context.Background()
	err := ms.Database.Client().Connect(ctx)
	if err != nil {
		return err
	}
	//nolint
	defer ms.Database.Client().Disconnect(ctx)

	collection := ms.Database.Collection(collectionName)
	result, err := collection.InsertOne(ctx, item)
	formats.Trace(result)
	if err != nil {
		log.Printf("failed to add item. %v", err)
	}
	return err
}

// upsert = true means that a new record will be created if none exists
func (ms *mongoService) UpdateOne(collectionName string, filter primitive.D, item interface{}, upsert bool) error {
	ctx := context.Background()
	err := ms.Database.Client().Connect(ctx)
	if err != nil {
		return err
	}
	//nolint
	defer ms.Database.Client().Disconnect(ctx)

	collection := ms.Database.Collection(collectionName)

	update := bson.D{
		{Key: "$set", Value: item},
	}
	opts := options.UpdateOptions{Upsert: &upsert}
	result, err := collection.UpdateOne(ctx, filter, update, &opts)
	formats.Trace(result)
	if err != nil {
		log.Printf("failed to add item. %v", err)
	}
	return err
}
