package database

import (
	"context"
	"fmt"

	"github.com/RINOHeinrich/multiserviceauth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// struct for MongoDB database
type MongoDB struct {
	config models.Dbconfig
	DB     *mongo.Client
}

// Connect to MongoDB
func (m *MongoDB) Connect() error {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", m.config.DBUser, m.config.DBHost, m.config.DBPassword, m.config.DBPort))
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}
	m.DB = client
	return nil
}

// Disconnect from MongoDB
func (m *MongoDB) Disconnect() error {
	err := m.DB.Disconnect(context.Background())
	if err != nil {
		return err
	}
	return nil
}

// Implement the methods of the Database interface for MongoDB
// Insert an user
func (m *MongoDB) Insert(data *models.User) error {
	userCollection := m.DB.Database(m.config.DBName).Collection("User")
	_, err := userCollection.InsertOne(context.Background(), data)
	if err != nil {
		return err
	}
	return nil
}

// Delete an user
func (m *MongoDB) Delete(id string) error {
	userCollection := m.DB.Database(m.config.DBName).Collection("User")
	filter := bson.E{Key: "id", Value: id}
	_, err := userCollection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

// Update an user
func (m *MongoDB) Update(id string, data *models.User) error {
	userCollection := m.DB.Database(m.config.DBName).Collection("User")
	filter := bson.E{Key: "id", Value: id}
	update := bson.D{{Key: "$set", Value: data}}
	_, err := userCollection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

// Find an user
func (m *MongoDB) Find(id string) (*models.User, error) {
	userCollection := m.DB.Database(m.config.DBName).Collection("User")
	filter := bson.E{Key: "id", Value: id}
	var user models.User
	err := userCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Find all users
func (m *MongoDB) FindAll() ([]models.User, error) {
	userCollection := m.DB.Database(m.config.DBName).Collection("User")
	cursor, err := userCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var users []models.User
	if err = cursor.All(context.Background(), &users); err != nil {
		return nil, err
	}
	return users, nil
}
func (m *MongoDB) LoadConfig(filename string) error {

	return nil
}
