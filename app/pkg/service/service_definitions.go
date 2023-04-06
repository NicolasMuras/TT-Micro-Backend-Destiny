package service

import (
	"TT-Micro-Backend-Destiny/pkg/dto"
	"TT-Micro-Backend-Destiny/pkg/repository"
	"context"
)

// DestinyService contains the attributes for the Destiny object service.
type DestinyService struct {
	Repository repository.DestinyInterface
}

// DestinyInterface define the methods for the Destiny service.
type DestinyInterface interface {
	RetrieveDestiny(ctx context.Context, destinyID string) (dto.RetrieveDestinyResponse, error)
	UpdateDestiny(ctx context.Context, destinyReq dto.UpdateDestinyRequest, destinyID string) (dto.UpdateDestinyResponse, error)
	DeleteDestiny(ctx context.Context, destinyID string) (dto.DeleteDestinyResponse, error)
	AddDestiny(ctx context.Context, destinyReq dto.AddDestinyRequest) (dto.AddDestinyResponse, error)
	ListDestiny(ctx context.Context) (dto.ListDestinyResponse, error)
}

// MockDestinyService defines the mock Destiny Service.
type MockDestinyService struct{}

// NewDestinyService return the destiny service interface
// with "destiny" db collection.
func NewDestinyService(repo repository.Interface) DestinyInterface {

	return &DestinyService{
		Repository: repo,
	}
}
