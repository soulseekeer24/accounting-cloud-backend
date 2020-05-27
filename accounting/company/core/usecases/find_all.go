package company

import (
	company "accounting/company/core/domain"
	"context"
	"piwi-backend-clean/accounting/company/core/domain"
)

func (c *useCase) FindAll(ctx context.Context) (companies []company.Company, err error) {
	return c.repository.GetAll(ctx)
}
