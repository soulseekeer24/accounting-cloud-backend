package gateways

import (
	"accounting/company/core/interfaces"
	company "accounting/company/core/usecases"
	"accounting/company/infra/serializer"
	"io/ioutil"
	"log"
	"net/http"
)

type CompanyHandler interface {
	GetAll(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
}

type handler struct {
	useCase company.UseCase
}

func NewHandler(useCase company.UseCase) CompanyHandler {
	return &handler{useCase: useCase}
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {

	companies, err := h.useCase.FindAll(r.Context())
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

	companyStored, err := h.useCase.Create(r.Context(), company)
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

func (h *handler) serializer(contentType string) interfaces.Serializer {
	return &serializer.JsonSerializer{}
}

func sendResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}
