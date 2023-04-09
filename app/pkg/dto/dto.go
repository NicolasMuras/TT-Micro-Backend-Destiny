package dto

import (
	"TT-Micro-Backend-Destiny/pkg/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RetrieveGroupDestinyResponse struct {
	DestinyGroup []models.Destiny `json:"destiny_group"`
}

type RetrieveGroupDestinyRequest struct {
	IDs []string `json:"destiny_ids"`
}

type RetrieveDestinyResponse struct {
	ID                  primitive.ObjectID       `bson:"_id,omitempty" json:"id,omitempty"`
	City                string                   `json:"city"`
	Country             string                   `json:"country"`
	IconFlagUrl         string                   `json:"icon_flag_url"`
	StartDate           time.Time                `json:"start_date"`
	EndDate             time.Time                `json:"end_date"`
	Images              []string                 `json:"images"`
	OverviewData        models.Overview          `json:"overview"`
	ExtraActivitiesData []models.ExtraActivities `json:"extra_activities"`
	Days                []models.Day             `json:"days"`
	Info                models.Info              `json:"info"`
	CreatedAt           time.Time                `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt           time.Time                `bson:"updated_at" json:"updated_at,omitempty"`
}

type UpdateDestinyResponse struct {
	Message string `json:"message"`
}

type UpdateDestinyRequest struct {
	City                string                   `json:"city"`
	Country             string                   `json:"country"`
	IconFlagUrl         string                   `json:"icon_flag_url"`
	StartDate           time.Time                `json:"start_date"`
	EndDate             time.Time                `json:"end_date"`
	Images              []string                 `json:"images"`
	OverviewData        models.Overview          `json:"overview"`
	ExtraActivitiesData []models.ExtraActivities `json:"extra_activities"`
	Days                []models.Day             `json:"days"`
	Info                models.Info              `json:"info"`
	CreatedAt           time.Time                `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt           time.Time                `bson:"updated_at" json:"updated_at,omitempty"`
}

type DeleteDestinyResponse struct {
	Message string `json:"message"`
}

type ListDestinyResponse struct {
	DestinyList []models.Destiny `json:"destiny_list"`
}

type AddDestinyRequest struct {
	City                string                   `json:"city"`
	Country             string                   `json:"country"`
	IconFlagUrl         string                   `json:"icon_flag_url"`
	StartDate           time.Time                `json:"start_date"`
	EndDate             time.Time                `json:"end_date"`
	Images              []string                 `json:"images"`
	OverviewData        models.Overview          `json:"overview"`
	ExtraActivitiesData []models.ExtraActivities `json:"extra_activities"`
	Days                []models.Day             `json:"days"`
	Info                models.Info              `json:"info"`
	CreatedAt           time.Time                `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt           time.Time                `bson:"updated_at" json:"updated_at,omitempty"`
}

type AddDestinyResponse struct {
	ID      interface{} `bson:"_id,omitempty" json:"id,omitempty"`
	Message string      `json:"message"`
}
