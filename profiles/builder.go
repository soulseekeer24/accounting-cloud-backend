package profiles

import (
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"piwi-backend-clean/profiles/core"
	"piwi-backend-clean/profiles/infrastructure/gateway"
	"piwi-backend-clean/profiles/infrastructure/persistency"
)

func BuildModule(client *mongo.Client, r *chi.Mux) *core.Module {

	mongoStore := persistency.NewMongoDBProfileStoreRepository(client.Database("accounting-app"))
	users := core.BuildModule(mongoStore)

	//Http Controller
	httpController := gateway.NewHttpController(users)

	r.Get("/profiles/me", httpController.Me)
	r.Post("/profiles",httpController.CreateProfile)

	return users
}
