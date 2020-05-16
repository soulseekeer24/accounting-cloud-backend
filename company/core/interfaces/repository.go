package interfaces

import (
	"context"
	"piwi-backend-clean/company/core/domain"
)

type Repository interface {
	GetAll(ctx context.Context) (companies []company.Company, err error)
	Find(ctx context.Context, id string) (company *company.Company,err error)
	Store(ctx context.Context,company *company.Company) (companyStored *company.Company,err error)
	Delete(ctx context.Context,companyID string) (err error)
}
