package db

import (
	"context"
	"github.com/reward-rabieth/series-troxide/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	showCollection = "shows"
)

type ShowHandler interface {
	StoreShows(ctx context.Context, seriesList []types.Show) error
}

type MongoDBShowRepo struct {
	coll *mongo.Collection
}

func NewMongoDBShowRepo(db *mongo.Database) *MongoDBShowRepo {
	return &MongoDBShowRepo{
		coll: db.Collection(showCollection),
	}
}

func (m *MongoDBShowRepo) StoreShows(ctx context.Context, shows []types.Show) error {
	var write []mongo.WriteModel
	for _, shows := range shows {
		// Create an 'UpdateOne' model with 'Upsert' option to insert or update the document in the collection
		filter := bson.M{"_id": shows.Id}
		update := bson.M{"$set": shows}
		updateModel := mongo.NewUpdateOneModel().SetFilter(filter).SetUpdate(update).SetUpsert(true)
		writes := append(write, updateModel)
		// Perform the bulk write operation to the collection
		_, err := m.coll.BulkWrite(ctx, writes, options.BulkWrite().SetOrdered(false))
		if err != nil {
			return err
		}
	}
	return nil
}
