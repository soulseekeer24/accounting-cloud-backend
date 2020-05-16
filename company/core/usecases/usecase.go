package company

import (
	"context"
	company "piwi-backend-clean/company/core/domain"
	"piwi-backend-clean/company/core/interfaces"
)

type useCase struct {
	repository interfaces.Repository
}

type UseCase interface {
	Create(ctx context.Context, company *company.Company) (companyStored *company.Company, err error)
	Delete(ctx context.Context,companyID string) (err error)
	FindAll(ctx context.Context) (companies []company.Company, err error)
}

func NewUseCase(repository interfaces.Repository) UseCase {
	return &useCase{repository: repository}
}
