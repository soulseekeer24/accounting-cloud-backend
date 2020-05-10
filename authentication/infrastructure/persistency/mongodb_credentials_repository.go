package persistency

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"piwi-backend-clean/authentication/core/domains/accounts"
	"piwi-backend-clean/common"
	"piwi-backend-clean/common/persistency"
)

type MongoDBAccountsRepository struct {
	*persistency.MongoDB
	db *mongo.Collection
}

func NewMongoDBAccountsRepository(db *mongo.Collection) *MongoDBAccountsRepository {
	return &MongoDBAccountsRepository{
		db:      db,
		MongoDB: persistency.NewMongoDBRepo(db),
	}
}

func (r *MongoDBAccountsRepository) SaveAccount(ctx context.Context, acc *accounts.Account) (ID string, err error) {
	return r.Save(ctx, acc)
}

func (r *MongoDBAccountsRepository) UpdateAccount(ctx context.Context, ID string, account *accounts.Account) (success bool, err error) {
	return r.Update(ctx, ID, account)
}

func (r *MongoDBAccountsRepository) GetAccountsByUserName(ctx context.Context, username string) (account *accounts.Account, err error) {
	query := bson.M{"username": username}
	account = &accounts.Account{}
	err = r.GetBy(ctx, query, account)
	if err != nil {
		switch err.(type) {
		case common.ErrDontExist:
			return nil,accounts.ErrAccountDontExist{}
		}
		return nil, err
	}
	return account, nil
}

func (r *MongoDBAccountsRepository) GetAccountsByValidationHash(ctx context.Context, hash string) (account *accounts.Account, err error) {

	query := bson.M{"validation_hash": hash}
	account = &accounts.Account{}

	err = r.GetBy(ctx, query, account)
	if err != nil {
		switch err.(type) {
		case common.ErrDontExist:
			return nil,accounts.ErrAccountDontExist{}
		}
		return nil, err
	}

	return account, nil
}

func (r *MongoDBAccountsRepository) GetAccountsByEmail(ctx context.Context, email string) (account *accounts.Account, err error) {
	query := bson.M{"email": email}
	account = &accounts.Account{}
	err = r.GetBy(ctx, query, account)
	if err != nil {
		switch err.(type) {
		case common.ErrDontExist:
			return nil,accounts.ErrAccountDontExist{}
		}
		return nil, err
	}

	return account, nil
}
