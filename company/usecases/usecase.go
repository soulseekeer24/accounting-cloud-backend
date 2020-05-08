package company

import company "piwi-backend-clean/company/domain"

type useCase struct {
	repository company.Repository
}

type UseCase interface {
	Create(company *company.Company) (companyStored *company.Company, err error)
	Delete(companyID string) (err error)
	FindAll() (companies []company.Company, err error)
}

func NewUseCase(repository company.Repository) UseCase {
	return &useCase{repository: repository}
}
