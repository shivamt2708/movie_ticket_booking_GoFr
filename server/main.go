package main

import (
	"context"
	"gofr.dev/pkg/gofr"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
    "fmt"
    "encoding/json"
	"log"
)

type User struct {
	email string `json:"email"`
    username string `json:"username"`
    password string `json:"password"`
    role string `json:"role"`
    location string `json:"location"`
}
type sign_response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func connectToMongoDB() (*mongo.Database, error) {
    // MongoDB connection string
    connectionString := "mongodb+srv://shivamt2708:123456778@cluster1.3y0qq5z.mongodb.net/"

    // Create a new MongoDB client
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectionString))
    if err != nil {
        return nil, err
    }

    // Check the connection
    err = client.Ping(context.Background(), nil)
    if err != nil {
        return nil, err
    }

    // Access the "test_db" database
    databaseName := "customers"
    database := client.Database(databaseName)

    log.Println("Connected to MongoDB")

    return database, nil
}

func handleSignup(ctx *gofr.Context) (interface{}, error) {
    var requestBody User
	err := json.NewDecoder(ctx.Request().Body).Decode(&requestBody)
	if err != nil {
		return nil, fmt.Errorf("error decoding request body: %w", err)
	}

    // Get the MongoDB database instance
    db, err := connectToMongoDB()
    if err != nil {
        return nil, err
    }

    // Access the "customers" collection
    collection := db.Collection("users")

    collection.InsertOne(context.Background(), bson.M{"email": requestBody.email, "username": requestBody.username, "password": requestBody.password, "role": requestBody.role, "location": requestBody.location})

    response := sign_response{
		Success: true,
		Message: "User created successfully",
	}

	// Return the response object as JSON
	return response.Success, err
}


func main() {
    app := gofr.New()

    // Register your MongoDB-connected handlers
    app.POST("/signup", handleSignup)

    // Start the server
    app.Start()
}