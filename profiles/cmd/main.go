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
	"piwi-backend-clean/profiles"
	"syscall"
	"time"
)

const dbUrl = "mongodb://localhost:27017"

func main() {
	client, cancel := ConnectMongoDB(dbUrl)
	defer cancel()
	r := chi.NewRouter()
	_ = profiles.BuildModule(client, r)
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
