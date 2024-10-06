package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type ExtendedGuest struct {
	Id    bson.ObjectID `json:"id" bson:"id"`
	Name  string        `json:"name" bson:"name" validate:"min=3,max=50"`
	IsVip bool          `json:"isVip" bson:"isVip"`
}

// We are using the extended reference pattern
// https://www.mongodb.com/blog/post/building-with-patterns-the-extended-reference-pattern
type Event struct {
	Id         bson.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string          `json:"name" bson:"name" validate:"min=3,max=50"`
	DressCode  string          `json:"dressCode" bson:"dressCode"`
	Location   string          `json:"location" bson:"location"`
	Date       bson.DateTime   `json:"date" bson:"date"`
	IsVipEvent bool            `json:"isVipEvent" bson:"isVipEvent"`
	Organizers []bson.ObjectID `json:"organizers" bson:"organizers"`
	Guests     []ExtendedGuest `json:"events" bson:"events"`
}

func NewEvent() *Event {
	return &Event{
		Guests:     []ExtendedGuest{},
		Organizers: []bson.ObjectID{},
	}
}
