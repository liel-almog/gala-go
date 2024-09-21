package models

import "go.mongodb.org/mongo-driver/v2/bson"

type BasicRegistration struct {
	GuestId bson.ObjectID `json:"guestId" bson:"guestId"`
	EventId bson.ObjectID `json:"eventId" bson:"eventId"`
}
