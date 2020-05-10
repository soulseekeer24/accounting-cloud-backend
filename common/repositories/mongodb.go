package repositories

import (
	"fmt"
	"log"
	"reflect"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"piwi-backend-clean/common/errors"
)

type MongoDB struct {

	dbContext *mongo.Collection
}

func NewMongoDBRepo(db *mongo.Collection) *MongoDB {
	return &MongoDB{dbContext: db}
}

func (r MongoDB) GetAll(ctx context.Context) (list []interface{}, err error) { return }
func (r MongoDB) Save(ctx context.Context, entity interface{}) (ID string, err error) {
	result, err := r.dbContext.InsertOne(ctx, entity)
	if err != nil {
		return "", err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		return oid.Hex(), nil
	}

	return fmt.Sprintf("%v", result.InsertedID), nil
}

func (r MongoDB) Update(ctx context.Context, ID string, entity interface{}) (ok bool, err error) {
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updating ID %v with %v \n", ID, entity)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": entity}
	result, err := r.dbContext.UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}
	fmt.Println(result.ModifiedCount)
	ok = result.ModifiedCount == 1
	return ok, nil
}

func (r MongoDB) Delete(ctx context.Context, ID string) (ok bool, err error) { return }

func (r MongoDB) GetByID(ctx context.Context, ID string, output interface{}) (err error) {
	return
}

func (r MongoDB) GetBy(ctx context.Context, query interface{}, output interface{}) (err error) {

	err = r.dbContext.FindOne(ctx, query).Decode(output)
	if err != nil {
		switch err.Error() {
		case "mongo: no documents in result":
			return errors.DontExist{}
		default:
			return err
		}

	}
	fmt.Printf("Result of mongodb : %v", output)
	return nil
}

func (r MongoDB) GetAllBy(ctx context.Context, query interface{}, schema interface{}) (list []interface{}, err error) {
	cursor, err := r.dbContext.Find(ctx, query)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		el := reflect.New(reflect.TypeOf(schema))
		err := cursor.Decode(&el)
		if err != nil {
			log.Fatal(err)
		}

		list = append(list, el)
	}

	return list, nil
}
