package repositories

import (
	"context"
	"sync"

	"github.com/liel-almog/gala-go/backend/database"
	"github.com/liel-almog/gala-go/backend/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type GuestRepository interface {
	FindAll(ctx context.Context) ([]models.Guest, error)
	FindById(ctx context.Context, id bson.ObjectID) (*models.Guest, error)
	Create(ctx context.Context, guest *models.Guest) error
}

type guestRepositoryImpl struct {
	db *database.MongoClient
}

var (
	initGuestRepositoryOnce sync.Once
	guestRepository         GuestRepository
)

func newGuestRepository() *guestRepositoryImpl {
	return &guestRepositoryImpl{
		db: database.GetDB(),
	}
}

func GetGuestRepository() GuestRepository {
	initGuestRepositoryOnce.Do(func() {
		guestRepository = newGuestRepository()
	})

	return guestRepository
}

func (r *guestRepositoryImpl) FindAll(ctx context.Context) ([]models.Guest, error) {
	filter := bson.D{}

	cursor, err := r.db.GuestColl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var guests []models.Guest
	if err = cursor.All(ctx, &guests); err != nil {
		return nil, err
	}

	return guests, nil
}

func (r *guestRepositoryImpl) FindById(ctx context.Context, id bson.ObjectID) (*models.Guest, error) {
	filter := bson.D{{Key: "_id", Value: id}}

	var guest models.Guest
	if err := r.db.GuestColl.FindOne(ctx, filter).Decode(&guest); err != nil {
		return nil, err
	}

	return &guest, nil
}

func (r *guestRepositoryImpl) Create(ctx context.Context, guest *models.Guest) error {
	_, err := r.db.GuestColl.InsertOne(ctx, guest)
	return err
}
