package db

import (
	"context"
	"github.com/reward-rabieth/series-troxide/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	episodeCollection = "episodes"
)

type EpisodeRepo interface {
	StoreEpisode(ctx context.Context, episodeList []types.Episode) error
}

type MongoDbEpisodeRepo struct {
	coll *mongo.Collection
}

func NewMongoDbEpisodeRepo(db *mongo.Database) *MongoDbEpisodeRepo {
	return &MongoDbEpisodeRepo{
		coll: db.Collection(episodeCollection),
	}
}
func (m *MongoDbEpisodeRepo) StoreEpisode(ctx context.Context, episodeList []types.Episode) error {
	var write []mongo.WriteModel
	for _, episode := range episodeList {
		// Create an 'UpdateOne' model with 'Upsert' option to insert or update the document in the collection
		filter := bson.M{"_id": episode.ID}
		update := bson.M{"$set": episode}
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
