package gateways

import (
	"io/ioutil"
	"log"
	"net/http"
	domain "piwi-backend-clean/company/domain"
	js "piwi-backend-clean/company/infra/serializer"
	useCase "piwi-backend-clean/company/usecases"
)

type CompanyHandler interface {
	GetAll(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
}

type handler struct {
	useCase useCase.UseCase
}

func NewHandler(useCase useCase.UseCase) CompanyHandler {
	return &handler{useCase: useCase}
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	companies, err := h.useCase.FindAll()
	contentType := r.Header.Get("Content-Type")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	companiesBytes, err := h.serializer(contentType).Encode(&companies)

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	sendResponse(w, contentType, companiesBytes, http.StatusOK)
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")

	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	company, err := h.serializer(contentType).Decode(requestBody)

	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	companyStored, err := h.useCase.Create(company)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	companyBytes, err := h.serializer(contentType).Encode(&companyStored)

	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	sendResponse(w, contentType, companyBytes, http.StatusOK)
}

func (h *handler) serializer(contentType string) domain.Serializer {
	return &js.JsonSerializer{}
}

func sendResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}
