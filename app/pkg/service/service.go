package service

import (
	"TT-Micro-Backend-Destiny/pkg/dto"
	"TT-Micro-Backend-Destiny/pkg/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *DestinyService) RetrieveDestiny(ctx context.Context, destinyID string) (dto.RetrieveDestinyResponse, error) {
	destiny, err := s.Repository.RetrieveDestiny(ctx, destinyID)
	if err != nil {
		return dto.RetrieveDestinyResponse{}, err
	}
	oid, _ := primitive.ObjectIDFromHex(destinyID)
	return dto.RetrieveDestinyResponse{
		ID:                  oid,
		City:                destiny.City,
		Country:             destiny.Country,
		IconFlagUrl:         destiny.IconFlagUrl,
		StartDate:           destiny.StartDate,
		EndDate:             destiny.EndDate,
		Images:              destiny.Images,
		OverviewData:        destiny.OverviewData,
		ExtraActivitiesData: destiny.ExtraActivitiesData,
		Days:                destiny.Days,
		Info:                destiny.Info,
		CreatedAt:           destiny.CreatedAt,
		UpdatedAt:           destiny.UpdatedAt,
	}, nil
}

func (s *DestinyService) UpdateDestiny(ctx context.Context, destinyReq dto.UpdateDestinyRequest, destinyID string) (dto.UpdateDestinyResponse, error) {
	oid, _ := primitive.ObjectIDFromHex(destinyID)
	destiny := models.Destiny{
		ID:                  oid,
		City:                destinyReq.City,
		Country:             destinyReq.Country,
		IconFlagUrl:         destinyReq.IconFlagUrl,
		StartDate:           destinyReq.StartDate,
		EndDate:             destinyReq.EndDate,
		Images:              destinyReq.Images,
		OverviewData:        destinyReq.OverviewData,
		ExtraActivitiesData: destinyReq.ExtraActivitiesData,
		Days:                destinyReq.Days,
		Info:                destinyReq.Info,
		CreatedAt:           destinyReq.CreatedAt,
		UpdatedAt:           destinyReq.UpdatedAt,
	}

	err := s.Repository.UpdateDestiny(ctx, destiny, destinyID)
	if err != nil {
		return dto.UpdateDestinyResponse{}, err
	}

	return dto.UpdateDestinyResponse{
		Message: "[+] Destiny updated successfully",
	}, nil
}

func (s *DestinyService) DeleteDestiny(ctx context.Context, destinyID string) (dto.DeleteDestinyResponse, error) {
	_, err := s.Repository.DeleteDestiny(ctx, destinyID)
	if err != nil {
		return dto.DeleteDestinyResponse{}, err
	}
	return dto.DeleteDestinyResponse{
		Message: "[-] Destiny deleted successfully",
	}, nil
}

func (s *DestinyService) AddDestiny(ctx context.Context, destinyReq dto.AddDestinyRequest) (dto.AddDestinyResponse, error) {
	destiny := models.Destiny{
		City:                destinyReq.City,
		Country:             destinyReq.Country,
		IconFlagUrl:         destinyReq.IconFlagUrl,
		StartDate:           destinyReq.StartDate,
		EndDate:             destinyReq.EndDate,
		Images:              destinyReq.Images,
		OverviewData:        destinyReq.OverviewData,
		ExtraActivitiesData: destinyReq.ExtraActivitiesData,
		Days:                destinyReq.Days,
		Info:                destinyReq.Info,
		CreatedAt:           destinyReq.CreatedAt,
		UpdatedAt:           destinyReq.UpdatedAt,
	}
	id, err := s.Repository.AddDestiny(ctx, destiny)
	if err != nil {
		return dto.AddDestinyResponse{}, err
	}
	return dto.AddDestinyResponse{
		ID:      id,
		Message: "[+] Destiny created successfully",
	}, nil
}

func (s *DestinyService) ListDestiny(ctx context.Context) (dto.ListDestinyResponse, error) {
	destinyList, err := s.Repository.GetDestinyList(ctx)
	if err != nil {
		return dto.ListDestinyResponse{}, err
	}
	return dto.ListDestinyResponse{
		DestinyList: destinyList,
	}, nil
}
