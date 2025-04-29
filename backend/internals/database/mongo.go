package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UrlDocument struct {
	ID  string `bson:"_id,omitempty"`
	URL string `bson:"url"`
}

type MongoRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoRepository(uri, dbName, collectionName string) (*MongoRepository, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	collection := client.Database(dbName).Collection(collectionName)
	return &MongoRepository{client: client, collection: collection}, nil
}

func (m *MongoRepository) GetUrl(id string) (string, error) {
	var result UrlDocument
	err := m.collection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return "", err
	}
	return result.URL, nil
}

func (m *MongoRepository) SaveUrl(id string, url string) error {
	doc := UrlDocument{ID: id, URL: url}
	_, err := m.collection.InsertOne(context.TODO(), doc)
	return err
}
