package main

import (
	//"fmt"
	"log"
	"net/http"
	//"context"

	//"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"

)


//var client *mongo.Client

func main() {
	// Database connection
	//ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	//defer cancel()

	//credential := options.Credential{
	//	Username: "user",
	//	Password: "password",
	//}
	//clientOpts := options.Client().ApplyURI("mongodb://localhost:27017").
	//	SetAuth(credential)

	//var err error
	//client, err = mongo.Connect(ctx, clientOpts)

	//if err != nil{
	//	log.Default()
	//}


	//ctx := context.TODO()

	//// Connect to MongoDB
	//mongoconn := options.Client().ApplyURI("mongodb://root:password123@db:27017")
	//mongoclient, err := mongo.Connect(ctx, mongoconn)

	//if err != nil {
	//	panic(err)
	//}

	//if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
	//	panic(err)
	//}

	//fmt.Println("MongoDB successfully connected...")

	log.Fatal(http.ListenAndServe(":8079", routes()))
}