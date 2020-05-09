package authentication

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"piwi-backend-clean/authentication/core"
	"piwi-backend-clean/authentication/infrastructure/controllers"
	"piwi-backend-clean/authentication/infrastructure/persistency"
	"piwi-backend-clean/authentication/infrastructure/utils"

)

const DB_URI = "mongodb://localhost:27017"

	func BuildAuthModule(client *mongo.Client, r *gin.IRouter) *core.Module {

	mongoCredsRepo := persistency.NewMongoDBAccountsRepository(client.Database("m_market").Collection("accounts"))
	bcryptEncrypter := utils.BcryptEncripter{}
	jwtToken := &utils.JWTTokenManager{}
	auth := core.NewAuthentication(mongoCredsRepo, bcryptEncrypter, jwtToken)

	//Http Controller
	httpController := controllers.NewAuthHTTP(auth)

	//we add endpoints here to mux

	//r.HandleFunc("/signin", httpController.Signin).Methods("POST", "OPTIONs")
	//r.HandleFunc("/signup", httpController.SignUp).Methods("POST", "OPTIONs")
	//r.HandleFunc("/validate/{validation_code}", httpController.ValidateAccount).Methods("get")

	http.Handle("/", r)
	return auth
}
