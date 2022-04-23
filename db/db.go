package db

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Talent struct {
	ID        string   `json:"id"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Picture   string   `json:"picture"`
	Job       string   `json:"job"`
	Location  string   `json:"location"`
	Linkedin  string   `json:"linkedin"`
	Github    string   `json:"github"`
	Tags      []string `json:"tags"`
	Stage     string   `json:"stage"`
}

func FetchAllTalents(db *mongo.Client) ([]Talent, error) {
	talentsCollection := db.Database("crewdotwork").Collection("talents")

	cursor, err := talentsCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println("ERROR: unable to read all talents from db", err)
		return nil, err
	}

	var data []Talent
	if err = cursor.All(context.TODO(), &data); err != nil {
		log.Println("ERROR: Cannot marshal data from DB into struct", err)
		return nil, err
	}

	log.Println("Total talents fetched from DB:", len(data))
	return data, nil
}

func PopulateDBWithSampleData(db *mongo.Client) error {
	data, err := GetDataFromAPI()
	if err != nil {
		return err
	}

	var newData []interface{}
	for _, r := range data {
		newData = append(newData, r)
	}

	database := db.Database("crewdotwork")
	talentsCollection := database.Collection("talents")

	if len(data) > 0 {
		log.Println("Removing existing records from 'talents' collection")
		result, err := talentsCollection.DeleteMany(context.TODO(), bson.M{})
		if err != nil {
			log.Println("WARN: Cannot DELETE all existing records in talents collection", err)
		}
		log.Println("INFO: removed existing records from talents", result.DeletedCount)
	}

	result, err := talentsCollection.InsertMany(context.TODO(), newData)
	if err != nil {
		log.Println("ERROR: Unable to insert talents data into DB", err)
		return err
	}

	log.Println("Data inserted into DB", len(result.InsertedIDs))
	return nil
}

func GetDataFromAPI() (talents []Talent, err error) {
	url := "https://hiring.crew.work/v1/talents"
	resp, err := http.Get(url)
	log.Println("GET:", url)
	if err != nil {
		log.Println("ERROR: Unable to get JSON from API:", url)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("ERROR: Cannot read response body: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	if err = json.Unmarshal(body, &talents); err != nil {
		log.Println("ERROR: Cannot unmarshal JSON into struct", err)
	}

	log.Println("Total open positions fetched", len(talents))
	return
}
