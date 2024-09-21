package repositories

import (
	"context"
	"sync"

	"github.com/liel-almog/gala-go/backend/database"
	"github.com/liel-almog/gala-go/backend/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type EventRepository interface {
	FindAll(ctx context.Context) ([]models.Event, error)
}

type eventRepositoryImpl struct {
	db *database.MongoClient
}

var (
	initEventRepositoryOnce sync.Once
	eventRepository         *eventRepositoryImpl
)

func newEventRepository() *eventRepositoryImpl {
	return &eventRepositoryImpl{
		db: database.GetDB(),
	}
}

func GetEventController() EventRepository {
	initEventRepositoryOnce.Do(func() {
		eventRepository = newEventRepository()
	})

	return eventRepository
}

func (r *eventRepositoryImpl) FindAll(ctx context.Context) ([]models.Event, error) {
	filter := bson.D{}

	cursor, err := r.db.EventColl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var events []models.Event
	if err = cursor.All(ctx, &events); err != nil {
		return nil, err
	}

	return events, nil
}
