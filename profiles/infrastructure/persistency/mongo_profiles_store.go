package persistency

import (
	"context"
	"fmt"
	"piwi-backend-clean/authentication/core/domains/accounts"
	"piwi-backend-clean/common"
	"piwi-backend-clean/common/persistency"
	"piwi-backend-clean/profiles/core/domains/profiles"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDBProfileStore struct {
	*persistency.MongoDB
	db *mongo.Collection
}

func NewMongoDBProfileStoreRepository(db *mongo.Collection) *MongoDBProfileStore {
	return &MongoDBProfileStore{
		db:      db,
		MongoDB: persistency.NewMongoDBRepo(db),
	}
}

func (s *MongoDBProfileStore) StoreProfile(ctx context.Context, profile *profiles.Profile) (ID string, err error) {
	return r.Save(ctx, profile)
}

func (s *MongoDBProfileStore) FindProfileByID(ctx context.Context, ID string) (profile *profiles.Profile, err error) {
	query := bson.M{"_id": ID}
	profile = &accounts.Account{}
	err = s.GetBy(ctx, query, profile)
	if err != nil {
		switch err.(type) {
		case common.ErrDontExist:
			return nil,profiles.ProfileDontFoundError{}
		}
		return nil, err
	}
	return profile, nil
}

func (s *MongoDBProfileStore) FindProfileByAccountID(ctx context.Context, accountId string) (profile *profiles.Profile, err error) {
	err = s.db.FindOne(ctx, bson.M{"account_id": accountId}).Decode(&profile)
	if err != nil {
		fmt.Println(err)
		switch err.Error() {
		case "mongo: no documents in result":
			return nil, profiles.ProfileDontFoundError{}
		default:
			return nil, err
		}
	}
	return
}

func (s *MongoDBProfileStore) UpdateProfile(ctx context.Context, ID string, profile *profiles.Profile) (success bool, err error) {
	return s.Update(ctx, ID, profile)
}
