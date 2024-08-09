package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Trainer struct {
    Name string
    Age  int
    City string
}

func main() {
    // Load environment variables from .env file
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
    uri := os.Getenv("MONGODB_URI")
    serverAPI := options.ServerAPI(options.ServerAPIVersion1)
    opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

    client, err := mongo.Connect(context.TODO(), opts)
    if err != nil {
      panic(err)
    }

    
    err = client.Ping(context.Background(), nil)
    
    if err != nil{
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB!")

    collection := client.Database("test").Collection("trainers")

    // This will be all about Eson and bson

    // Inserting a document
    
    ash := Trainer{"Ash", 10, "Pallet Town"}
    misty := Trainer{"Misty", 10, "Cerulean City"}
    brock := Trainer{"Brock", 15, "Pewter City"}
    trainers := []interface{}{ash, misty, brock}
    insertResult, err := collection.InsertMany(context.TODO(), trainers)

    if err != nil{
        log.Fatal(err)
    }
    fmt.Println("Inserted multiple documents: ", insertResult.InsertedIDs)

    // Updating document 

    filter := map[string]interface{}{"name": "Ash"}
    update := map[string]interface{}{
        "$inc": map[string]interface{}{"age": 1},
    }


   collection.UpdateOne(context.TODO(), filter, update)

//    finding documents 
    var result Trainer

    err = collection.FindOne(context.TODO(), filter).Decode(&result)
    if err != nil{
        log.Fatal(err)
    }
    fmt.Printf("Found a single document: %+v\n", result)

    findOptions := options.Find()
    findOptions.SetLimit(2)

    var results []*Trainer

    cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions) 
    if err != nil{
        log.Fatal(err)
    }
    for cur.Next(context.TODO()){
        var elem Trainer
        err := cur.Decode(&elem)
        if err != nil{
            log.Fatal(err)
        }
        results = append(results, &elem)

    }
    if err := cur.Err(); err != nil{
        log.Fatal(err)
    }
    cur.Close(context.TODO())

    fmt.Printf("Found multiple documents (array of pointers): %+v\n", results)

    deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
    if err != nil{
        log.Fatal(err)
    }

    fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)



}

