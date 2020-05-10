package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
	"os/signal"
	"piwi-backend-clean/authentication"
	"piwi-backend-clean/middlewares"
	"piwi-backend-clean/profiles"
	"syscall"
	"time"
)

const DB_URI = "mongodb://localhost:27017"

func main() {
	client, cancel := ConnectMongoDB(DB_URI)
	defer cancel()
	r := chi.NewRouter()
	//AuthenticationModule
	authModule := authentication.BuildAuthModule(client, r)
	profilesModules := profiles.BuildModule(client, r)

	// Mount middlewares dependencies
	middlewares.SetAuthModule(authModule)
	middlewares.SetProfilesModule(profilesModules)

	errs := make(chan error, 2)

	go func() {
		fmt.Println("Listening on port :8000")
		errs <- http.ListenAndServe(":8000", r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
}

func ConnectMongoDB(uri string) (client *mongo.Client, cancel context.CancelFunc) {

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client, cancel
}
