package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Event struct {
	// ID         bson.ObjectID       `json:"_id" bson:"_id"`
	Name       string              `json:"name" bson:"name" validate:"min=3,max=50"`
	DressCode  string              `json:"dressCode" bson:"dressCode"`
	Location   string              `json:"location" bson:"location"`
	Date       bson.DateTime       `json:"date" bson:"date"`
	IsVipEvent bool                `json:"isVipEvent" bson:"isVipEvent"`
	Organizers []bson.ObjectID     `json:"organizers" bson:"organizers"`
	Guests     []BasicRegistration `json:"events" bson:"events"`
}

func NewEvent() *Event {
	return &Event{
		Guests:     []BasicRegistration{},
		Organizers: []bson.ObjectID{},
	}
}
