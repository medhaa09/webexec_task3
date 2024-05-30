package store

import (
	"context"
	"fmt"
	"task3/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type MongoStore struct {
	Collection1 *mongo.Collection
}

const uri = "mongodb+srv://medha:drumDRO67%23%24@cluster0.qj0tdiv.mongodb.net/"

func (m *MongoStore) OpenConnectionWithMongoDB() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("webexec").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	m.Collection1 = client.Database("webexec").Collection("user")

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func (m *MongoStore) StoreUserData(user models.User) error {
	fmt.Println("Trying to insert user data into MongoDB")

	// Hash the password before storing
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error hashing password:", err)
		return err
	}
	user.Password = string(hashedPassword)

	_, err = m.Collection1.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println("Error inserting user data:", err)
		return err
	}
	fmt.Println("Insertion of user data successful")
	return nil
}

func (m *MongoStore) UserLogin(username string, password string) bool {

	var foundUser models.User
	err := m.Collection1.FindOne(context.TODO(), bson.M{
		"username": username,
	}).Decode(&foundUser)

	if err != nil {
		fmt.Println("Wrong credentials: ", err)
		return false
	}

	// Compare the hashed password with the plain text password
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		fmt.Println("Invalid password: ", err)
		return false
	}

	return true
}
