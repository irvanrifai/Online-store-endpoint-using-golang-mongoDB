package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Item struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Size      int64              `json:"size,omitempty" bson:"size,omitempty"`
	Dimension string             `json:"dimension,omitempty" bson:"dimension,omitempty"`
	Price     string             `json:"price,omitempty" bson:"price,omitempty"`
	Quantity  int64              `json:"quantity,omitempty" bson:"quantity,omitempty"`
}

var client *mongo.Client

func Test() {

}

func CreateItemEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var item Item
	json.NewDecoder(request.Body).Decode(&item)
	collection := client.Database("online_store").Collection("items")
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(context.TODO(), item)
	json.NewEncoder(response).Encode(result)
	fmt.Println(result)
}

func GetAllItemEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var items []Item
	collection := client.Database("online_store").Collection("items")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"messsage":"` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var item Item
		cursor.Decode(&item)
		items = append(items, item)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"messsage":"` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(items)
}

func GetItemEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var item Item
	collection := client.Database("online_store").Collection("items")
	err := collection.FindOne(context.TODO(), Item{ID: id}).Decode(&item)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"messsage":"` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(item)
}

func UpdateItemEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
}

func DeleteItemEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var item Item
	collection := client.Database("online_store").Collection("items")
	err := collection.FindOne(context.TODO(), Item{ID: id}).Decode(&item)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"messsage":"` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(item)
}

func main() {
	fmt.Println("Starting Application...")
	client, _ = mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	client.Connect(context.TODO())
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		// Can't connect to Mongo server
		log.Fatal(err)
	}
	router := mux.NewRouter()
	router.HandleFunc("/", GetAllItemEndpoint).Methods("GET")
	router.HandleFunc("/{id}", GetItemEndpoint).Methods("GET")
	router.HandleFunc("/item", CreateItemEndpoint).Methods("POST")
	http.ListenAndServe(":3000", router)
}
