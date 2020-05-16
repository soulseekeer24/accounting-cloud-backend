package company

import "context"

type Repository interface {
	GetAll(ctx context.Context) (companies []Company, err error)
	Find(ctx context.Context, id string) (company *Company,err error)
	Store(ctx context.Context,company *Company) (companyStored *Company,err error)
	Delete(ctx context.Context,companyID string) (err error)
}
