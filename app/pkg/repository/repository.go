package repository

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"TT-Micro-Backend-Destiny/pkg/dto"
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

func (r *DestinyRepository) RetrieveGroupDestiny(ctx context.Context, destinyIds dto.RetrieveGroupDestinyRequest) ([]models.Destiny, error) {

	retrieveRequest := func(ch chan models.Destiny, key string) {
		var destiny models.Destiny

		if err := CollectionFindOne(r.db, ctx, bson.M{}).Decode(&destiny); err != nil {
			ch <- models.Destiny{}
		}

		ch <- destiny

		close(ch)
	}

	// We allocate dynamically the channels in this variable
	var channels []chan models.Destiny

	// Iterate through the ids needed to make the request to get the info.
	for _, id := range destinyIds.IDs {
		// A channel is created per id.
		ch := make(chan models.Destiny)
		// Append the new channel to the channels array.
		channels = append(channels, ch)
		// The http request is executed.
		go retrieveRequest(ch, id)
	}

	exit := make(chan struct{})

	mainChannel := merge(channels)

	var destinyList []models.Destiny

	go func() {
		for destiny := range mainChannel {
			// we need to return here the json of destinies
			fmt.Println(destiny)
			destinyList = append(destinyList, destiny)
		}
		close(exit)
	}()

	<-exit
	fmt.Println("\n[+] All request completed.")

	return destinyList, nil
}

// merge is used to merge a number of channels in one main channel
// and finally return the new channel.
func merge(cs []chan models.Destiny) <-chan models.Destiny {
	var wg sync.WaitGroup

	// we define the main channel
	out := make(chan models.Destiny)

	// we define a function that takes a channel as argument and decrement wait count in 1.
	send := func(c <-chan models.Destiny) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	// we set up the wait counter accordly to the amount of channels.
	wg.Add(len(cs))

	// we iterate through the channels and start to execute a go routine for each one.
	for _, c := range cs {
		go send(c)
	}
	// finally this goroutine help us to wait until the process end and close the main channel.
	go func() {
		wg.Wait()

		close(out)
	}()
	// we return the main channel
	return out
}
