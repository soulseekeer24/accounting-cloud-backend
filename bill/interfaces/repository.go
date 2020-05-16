package interfaces

import (
	"context"
	bill "piwi-backend-clean/bill/domain"
)

type Repository interface {
	GetAll(ctx context.Context) (bills []bill.Bill, err error)
	Find(ctx context.Context, id string) (bill *bill.Bill,err error)
	Store(ctx context.Context,bill *bill.Bill) (billStored *bill.Bill,err error)
	Delete(ctx context.Context,billID string) (err error)
}
