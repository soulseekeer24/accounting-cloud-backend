package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	company "piwi-backend-clean/company/cmd/usecases"
	handler "piwi-backend-clean/company/infra/gateways"
	"piwi-backend-clean/company/infra/persistence"
	"syscall"
)

func main() {

	repo, err := persistence.NewMongoRepository("mongodb://localhost:27017", "accouting", 300)
	if err != nil {
		log.Fatal(err)
	}
	useCase := company.NewUseCase(repo)
	handler := handler.NewHandler(useCase)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/companies", handler.GetAll)
	r.Post("/companies", handler.Create)

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
