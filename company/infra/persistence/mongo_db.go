package persistence

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	company "piwi-backend-clean/company/domain"
	"time"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout int) (company.Repository, error) {
	repo := &mongoRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, err
	}
	repo.client = client
	return repo, nil
}

func (m *mongoRepository) GetAll() (companies []company.Company, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	collection := m.client.Database(m.database).Collection("companies")

	cursor, err := collection.Find(ctx, bson.D{{}}, options.Find())
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var company company.Company
		err := cursor.Decode(&company)

		if err != nil {
			log.Fatal(err)
		}

		companies = append(companies, company)
	}

	return companies, nil
}

func (m *mongoRepository) Find(id string) (company *company.Company, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	collection := m.client.Database(m.database).Collection("companies")
	filter := bson.M{"id": id}

	err = collection.FindOne(ctx, filter).Decode(&company)
	if err != nil {
		return nil, err
	}
	return company, nil

}

func (m *mongoRepository) Store(company *company.Company) (companyStored *company.Company, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	collection := m.client.Database(m.database).Collection("companies")

	result, err := collection.InsertOne(
		ctx,
		bson.M{
			"name":                      company.Name,
			"tax_identification_number": company.TaxIdentificationNumber,
			"created_at":                company.CreatedAt,
		},
	)
	if err != nil {
		return nil, err
	}

	company.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return company, nil
}

func (m *mongoRepository) Delete(companyID string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), m.timeout)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(companyID)
	if err != nil {
		return err
	}

	collection := m.client.Database(m.database).Collection("companies")

	_, err = collection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		return err
	}
	return nil
}