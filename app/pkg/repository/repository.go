package repository

import (
	"context"
	"log"
	"time"

	"TT-Micro-Backend-Destiny/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var CollectionInsertOne = (*mongo.Collection).InsertOne
var CollectionFind = (*mongo.Collection).Find
var CollectionFindOne = (*mongo.Collection).FindOne
var CollectionUpdateOne = (*mongo.Collection).UpdateOne
var CollectionDeleteOne = (*mongo.Collection).DeleteOne
var CursorDecode = (*mongo.Cursor).Decode

func (r *DestinyRepository) RetrieveDestiny(ctx context.Context, destinyID string) (models.Destiny, error) {
	var destiny models.Destiny

	if err := CollectionFindOne(r.db, ctx, bson.M{}).Decode(&destiny); err != nil {
		return models.Destiny{}, err
	}
	return destiny, nil
}

func (r *DestinyRepository) UpdateDestiny(ctx context.Context, destiny models.Destiny, destinyID string) error {
	var err error

	oid, _ := primitive.ObjectIDFromHex(destinyID)
	filter := bson.M{"_id": oid}

	update := bson.M{
		"$set": bson.M{
			"city":             destiny.City,
			"country":          destiny.Country,
			"icon_flag_url":    destiny.IconFlagUrl,
			"start_date":       destiny.StartDate,
			"end_date":         destiny.EndDate,
			"images":           destiny.Images,
			"overview":         destiny.OverviewData,
			"extra_activities": destiny.ExtraActivitiesData,
			"days":             destiny.Days,
			"info":             destiny.Info,
			"updated_at":       time.Now(),
		},
	}

	_, err = CollectionUpdateOne(r.db, ctx, filter, update)

	if err != nil {
		log.Printf("[ERROR] Update destiny query failed")
		return err
	}
	return nil
}

func (r *DestinyRepository) DeleteDestiny(ctx context.Context, destinyID string) ([]models.Destiny, error) {
	destiny := []models.Destiny{}
	var oid primitive.ObjectID

	oid, err := primitive.ObjectIDFromHex(destinyID)

	if err != nil {
		return []models.Destiny{}, err
	}

	filter := bson.M{"_id": oid}

	_, err = CollectionDeleteOne(r.db, ctx, filter)

	if err != nil {
		log.Printf("[ERROR] Delete destiny query failed")
		return []models.Destiny{}, err
	}
	return destiny, nil
}

func (r *DestinyRepository) GetDestinyList(ctx context.Context) ([]models.Destiny, error) {
	var destinyList []models.Destiny

	filter := bson.D{}

	cur, err := CollectionFind(r.db, ctx, filter)

	if err != nil {
		log.Printf("[ERROR] Get destiny list query failed")
		return []models.Destiny{}, err
	}

	for cur.Next(ctx) {
		var destiny models.Destiny
		err = CursorDecode(cur, &destiny)

		if err != nil {
			log.Printf("[ERROR] Problems at decoding row...")
			return []models.Destiny{}, err
		}

		destinyList = append(destinyList, destiny)
	}

	return destinyList, nil
}

func (r *DestinyRepository) AddDestiny(ctx context.Context, destiny models.Destiny) (interface{}, error) {
	result, err := CollectionInsertOne(r.db, ctx, destiny)
	id := result.InsertedID
	if err != nil {
		log.Printf("[ERROR] Insert destiny query failed")
		return primitive.ObjectID{}, err
	}
	return id, nil
}
