package controller

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

const connectionString = "mongodb+srv://pranavkonde:<Pradip@123>@cluster0.3609jz4.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

// Most important

var collection *mongo.Collection

// connect with mongoDB
// init method only executes only first time in start for the initialization
func init() {
	//client option
	clientOption := options.Client().ApplyURI(connectionString)
	// connect to mongodb
	mongo.connect(context.T)
}
