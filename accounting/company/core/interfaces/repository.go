package interfaces

import (
	company "accounting/company/core/domain"
	"context"
)

type Repository interface {
	GetAll(ctx context.Context) (companies []company.Company, err error)
	Find(ctx context.Context, id string) (company *company.Company,err error)
	Store(ctx context.Context,company *company.Company) (companyStored *company.Company,err error)
	Delete(ctx context.Context,companyID string) (err error)
}
