package repository

import (
	"TT-Micro-Backend-Destiny/pkg/dto"
	"TT-Micro-Backend-Destiny/pkg/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// Repository contains attributes for the object repository.
type DestinyRepository struct {
	db *mongo.Collection
}

// DestinyRepository define the methods for the destiny object.
type DestinyInterface interface {
	RetrieveDestiny(ctx context.Context, destinyID string) (models.Destiny, error)
	UpdateDestiny(ctx context.Context, destiny models.Destiny, destinyID string) error
	DeleteDestiny(ctx context.Context, destinyID string) ([]models.Destiny, error)
	GetDestinyList(ctx context.Context) ([]models.Destiny, error)
	AddDestiny(ctx context.Context, destiny models.Destiny) (interface{}, error)
	RetrieveGroupDestiny(ctx context.Context, destinyIDs dto.RetrieveGroupDestinyRequest) ([]models.Destiny, error)
}

// Interface define the methods for the repository.
type Interface interface {
	DestinyInterface
}

// NewHandler Return the Handlder with all the handler dependencies.
func NewDestinyRepository(db *mongo.Collection) Interface {
	return &DestinyRepository{
		db: db,
	}
}
