package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Event struct {
	ID         bson.ObjectID       `json:"_id" bson:"_id"`
	Name       string              `json:"name" bson:"name"`
	DressCode  string              `json:"dressCode" bson:"dressCode"`
	Location   string              `json:"location" bson:"location"`
	Date       bson.DateTime       `json:"date" bson:"date"`
	IsVipEvent bool                `json:"isVipEvent" bson:"isVipEvent"`
	Organizers []bson.ObjectID     `json:"organizers" bson:"organizers"`
	Events     []BasicRegistration `json:"events" bson:"events,omitempty"`
}
