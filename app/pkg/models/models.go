package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Overview struct {
	About string `json:"about"`
}

type ExtraActivities struct {
	ActivityName        string   `json:"activity_name"`
	Address             string   `json:"address"`
	LocationCoordinates []string `json:"location_coordinates"`
	Image               string   `json:"image"`
	About               string   `json:"about"`
	OpenHour            string   `json:"open_hour"`
	CloseHour           string   `json:"close_hour"`
	Price               string   `json:"price"`
}

type TransportData struct {
	OriginLocationName  string    `json:"origin_location_name"`
	DestinyLocationName string    `json:"destiny_location_name"`
	LocationCoordinates []string  `json:"location_coordinates"`
	StartHour           string    `json:"start_hour"`
	EndHour             string    `json:"end_hour"`
	StartDate           time.Time `json:"start_date"`
	EndDate             time.Time `json:"end_date"`
}

type ActivityData struct {
	ActivityName        string   `json:"activity_name"`
	Address             string   `json:"address"`
	Images              []string `json:"images"`
	LocationCoordinates []string `json:"location_coordinates"`
	Dscription          string   `json:"description"`
	StartHour           string   `json:"start_hour"`
	EndHour             string   `json:"end_hour"`
}

type HotelData struct {
	HotelName           string    `json:"hotel_name"`
	SpokenLanguages     []string  `json:"spoken_languages"`
	Stars               int       `json:"stars"`
	Images              []string  `json:"images"`
	LocationCoordinates []string  `json:"location_coordinates"`
	CheckinHour         string    `json:"checkin_hour"`
	CheckoutHour        string    `json:"checkout_hour"`
	StartDate           time.Time `json:"start_date"`
	EndDate             time.Time `json:"end_date"`
	Rooms               []string  `json:"rooms"`
	Breakfast           bool      `json:"breakfast"`
	Wifi                bool      `json:"wifi"`
	Address             string    `json:"address"`
	PhoneNumber         string    `json:"phone_number"`
	ContactEmail        string    `json:"contact_email"`
}

type EventData struct {
	HotelData     HotelData     `json:"hotel_data"`
	ActivityData  ActivityData  `json:"activity_data"`
	TransportData TransportData `json:"transport_data"`
}

type Event struct {
	EventType string    `json:"event_type"`
	HoverInfo string    `json:"hover_info"`
	EventIcon string    `json:"event_icon"`
	Date      time.Time `json:"date"`
	EventData EventData `json:"event_data"` // EventData -> interface{}
}

type Day struct {
	Events []Event `json:"events"`
}

type Section struct {
	Title       string `json:"title"`
	Information string `json:"information"`
}

type Info struct {
	Sections []Section `json:"sections"`
}

type Destiny struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	City                string             `json:"city"`
	Country             string             `json:"country"`
	IconFlagUrl         string             `json:"icon_flag_url"`
	StartDate           time.Time          `json:"start_date"`
	EndDate             time.Time          `json:"end_date"`
	Images              []string           `json:"images"`
	OverviewData        Overview           `json:"overview"`
	ExtraActivitiesData []ExtraActivities  `json:"extra_activities"`
	Days                []Day              `json:"days"`
	Info                Info               `json:"info"`
	CreatedAt           time.Time          `bson:"created_at" json:"created_at,omitempty"`
	UpdatedAt           time.Time          `bson:"updated_at" json:"updated_at,omitempty"`
}
