package models

import "go.mongodb.org/mongo-driver/v2/bson"

type ExtendedEvent struct {
	Id   bson.ObjectID `json:"id" bson:"id"`
	Name string        `json:"name" bson:"name" validate:"min=3,max=50"`
}

type CustomRequest struct {
	Fulfilled   bool          `json:"fulfilled" bson:"fulfilled"`
	Description string        `json:"description" bson:"description" validate:"min=3,max=250"`
	Id          bson.ObjectID `json:"id" bson:"id"`
}

// We are using the extended reference pattern
// https://www.mongodb.com/blog/post/building-with-patterns-the-extended-reference-pattern
type Guest struct {
	Id             bson.ObjectID   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name           string          `json:"name" bson:"name" validate:"min=3,max=50"`
	IsVip          bool            `json:"isVip" bson:"isVip"`
	Age            int8            `json:"age" bson:"age" validate:"gte=18,lte=120"`
	Events         []ExtendedEvent `json:"events" bson:"events"`
	CustomRequests []CustomRequest `json:"customRequests" bson:"customRequests"`
}

func NewGuest() *Guest {
	return &Guest{
		Events:         []ExtendedEvent{},
		CustomRequests: []CustomRequest{},
	}
}
