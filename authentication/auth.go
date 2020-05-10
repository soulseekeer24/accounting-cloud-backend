package authentication

import (
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"piwi-backend-clean/authentication/core"
	"piwi-backend-clean/authentication/infrastructure/controllers"
	"piwi-backend-clean/authentication/infrastructure/persistency"
	"piwi-backend-clean/authentication/infrastructure/utils"
)


func BuildAuthModule(client *mongo.Client, r *chi.Mux) *core.Module {

	mongoCredsRepo := persistency.NewMongoDBAccountsRepository(client.Database("accounting-app").Collection("accounts"))
	bcryptEncrypter := utils.BcryptEncripter{}
	jwtToken := &utils.JWTTokenManager{}
	auth := core.NewAuthentication(mongoCredsRepo, bcryptEncrypter, jwtToken)

	//Http Controller
	httpController := controllers.NewAuthHTTP(auth)

	//we add endpoints here to mux

	r.Post("/auth/signin", httpController.Signin)
	r.Post("/auth/signup", httpController.SignUp)
	r.Get("/auth/validate/{validation_code}", httpController.ValidateAccount)


	return auth
}
