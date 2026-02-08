package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jcfullmer/api2db/internal/database"
)

type NpsBody struct {
	Data []struct {
		ID          string          `json:"id"`
		FullName    string          `json:"fullName"`
		ParkCode    string          `json:"parkCode"`
		States      string          `json:"states"`
		Description string          `json:"description"`
		Designation string          `json:"designation"`
		Activities  json.RawMessage `json:"activities"`
		Topics      json.RawMessage `json:"topics"`

		WeatherInfo    string          `json:"weatherInfo"`
		Contacts       json.RawMessage `json:"contacts"`
		EntranceFees   json.RawMessage `json:"entranceFees"`
		EntrancePasses json.RawMessage `json:"entrancePasses"`
		Fees           json.RawMessage `json:"fees"`
		OperatingHours json.RawMessage `json:"operatingHours"`
		Addresses      json.RawMessage `json:"addresses"`
		Images         json.RawMessage `json:"images"`
		Multimedia     json.RawMessage `json:"multimedia"`
	} `json:"data"`
}

type ParkDetails struct {
	WeatherInfo    string          `json:"weatherInfo"`
	Contacts       json.RawMessage `json:"contacts"`
	EntranceFees   json.RawMessage `json:"entranceFees"`
	EntrancePasses json.RawMessage `json:"entrancePasses"`
	Fees           json.RawMessage `json:"fees"`
	OperatingHours json.RawMessage `json:"operatingHours"`
	Addresses      json.RawMessage `json:"addresses"`
	Images         json.RawMessage `json:"images"`
	Multimedia     json.RawMessage `json:"multimedia"`
}

func ResponseToDB(res *http.Response, db *database.Queries) error {
	ParkStruct, err := jsonToStruct(res)
	if err != nil {
		return err
	}
	err = structToDB(db, ParkStruct)
	if err != nil {
		return err
	}
	return nil
}

func jsonToStruct(res *http.Response) (NpsBody, error) {
	result := NpsBody{}
	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(&result)
	if err != nil {
		return NpsBody{}, err
	}
	return result, nil
}

func structToDB(db *database.Queries, parks NpsBody) error {
	for _, park := range parks.Data {
		exists, err := db.CheckExists(context.Background(), park.ID)
		if err != nil {
			return err
		}
		if exists {
			log.Printf("%s already exists.", park.FullName)
			continue
		}
		details := ParkDetails{
			WeatherInfo:    park.WeatherInfo,
			Contacts:       park.Contacts,
			EntranceFees:   park.EntranceFees,
			EntrancePasses: park.EntrancePasses,
			Fees:           park.Fees,
			OperatingHours: park.OperatingHours,
			Addresses:      park.Addresses,
			Images:         park.Images,
			Multimedia:     park.Multimedia,
		}
		jsonDetails, err := json.Marshal(details)
		if err != nil {
			return err
		}
		params := database.CreateUserParams{
			ID:          uuid.New(),
			NpsID:       park.ID,
			FullName:    park.FullName,
			ParkCode:    park.ParkCode,
			States:      park.States,
			Description: park.Description,
			Designation: park.Designation,
			Activities:  park.Activities,
			Topics:      park.Topics,
			Details:     jsonDetails,
			CreatedAt:   time.Now(),
		}
		_, err = db.CreateUser(context.Background(), params)
		if err != nil {
			return err
		}
		log.Printf("Added %s to database\n", park.FullName)
	}
	return nil
}
