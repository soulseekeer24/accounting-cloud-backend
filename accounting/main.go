package accounting

import (
	company "accounting/company/core/usecases"
	companyInfra "accounting/company/infra/gateways"
	"accounting/company/infra/persistence"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	repo, err := persistence.NewMongoRepository("mongodb://localhost:27017", "accouting", 300)
	if err != nil {
		log.Fatal(err)
	}
	useCase := company.NewUseCase(repo)
	companyHandler := companyInfra.NewHandler(useCase)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/companies", companyHandler.GetAll)
	r.Post("/companies", companyHandler.Create)

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
