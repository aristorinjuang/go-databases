package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type article struct {
	ID      primitive.ObjectID `bson:"_id"`
	Title   string             `bson:"Title"`
	Content string             `bson:"Content"`
}

var (
	articles   []article
	client     *mongo.Client
	collection *mongo.Collection
	ctx        context.Context
	err        error
	objectID   primitive.ObjectID
)

func index(w http.ResponseWriter, r *http.Request) {
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var document article
		err := cur.Decode(&document)
		if err != nil {
			log.Fatal(err)
		}
		articles = append(articles, document)
	}

	response, err := json.Marshal(articles)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func create(w http.ResponseWriter, r *http.Request) {
	var request article
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	request.ID = primitive.NewObjectID()

	insertResult, err := collection.InsertOne(context.TODO(), request)
	if err != nil {
		log.Fatal(err)
	}
	if insertResult.InsertedID != nil {
		response, err := json.Marshal(article{request.ID, request.Title, request.Content})
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

func read(w http.ResponseWriter, r *http.Request, key string) {
	objectID, err = primitive.ObjectIDFromHex(key)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{Key: "_id", Value: objectID}}

	var document article
	err = collection.FindOne(ctx, filter).Decode(&document)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	response, err := json.Marshal(document)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func update(w http.ResponseWriter, r *http.Request, key string) {
	objectID, err = primitive.ObjectIDFromHex(key)
	if err != nil {
		log.Fatal(err)
	}

	var request article
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	request.ID = objectID

	filter := bson.D{{Key: "_id", Value: objectID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "Title", Value: request.Title},
			{Key: "Content", Value: request.Content},
		}},
	}

	updateResult, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	if updateResult.MatchedCount != 0 {
		response, err := json.Marshal(request)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}
}

func delete(w http.ResponseWriter, r *http.Request, key string) {
	objectID, err = primitive.ObjectIDFromHex(key)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{Key: "_id", Value: objectID}}
	deleteResult, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	if deleteResult.DeletedCount != 1 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected!")

	collection = client.Database("go_database").Collection("articles")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		key := strings.Split(r.URL.Path, "/")[1]

		switch r.Method {
		case http.MethodGet:
			if key == "" {
				index(w, r)
			} else {
				read(w, r, key)
			}
		case http.MethodPost:
			create(w, r)
		case http.MethodPut:
			update(w, r, key)
		case http.MethodDelete:
			delete(w, r, key)
		default:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	})
	http.ListenAndServe(":8000", nil)
}
