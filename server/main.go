package main

import (
	"context"
	"gofr.dev/pkg/gofr"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Customer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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

func handleCreateCustomer(ctx *gofr.Context) (interface{}, error) {
    name := ctx.PathParam("name")

    // Get the MongoDB database instance
    db, err := connectToMongoDB()
    if err != nil {
        return nil, err
    }

    // Access the "customers" collection
    collection := db.Collection("customers")

    // Insert a new customer document
    result, err := collection.InsertOne(context.Background(), bson.M{"name": name})
    if err != nil {
        return nil, err
    }

    log.Printf("Inserted customer with ID: %v", result.InsertedID)

    return nil, nil
}

func handleGetCustomers(ctx *gofr.Context) (interface{}, error) {
    // Get the MongoDB database instance
    db, err := connectToMongoDB()
    if err != nil {
        return nil, err
    }

    // Access the "customers" collection
    collection := db.Collection("customers")

    // Find all customers
    cur, err := collection.Find(context.Background(), bson.D{})
    if err != nil {
        return nil, err
    }
    defer cur.Close(context.Background())

    var customers []Customer
    for cur.Next(context.Background()) {
        var customer Customer
        if err := cur.Decode(&customer); err != nil {
            return nil, err
        }
        customers = append(customers, customer)
    }

    if err := cur.Err(); err != nil {
        return nil, err
    }

    return customers, nil
}

func main() {
    app := gofr.New()

    // Register your MongoDB-connected handlers
    app.POST("/customer/{name}", handleCreateCustomer)
    app.GET("/customer", handleGetCustomers)

    // Start the server
    app.Start()
}