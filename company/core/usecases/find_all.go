package company

import (
	"context"
	"piwi-backend-clean/company/core/domain"
)

func (c *useCase) FindAll(ctx context.Context) (companies []company.Company, err error) {
	return c.repository.GetAll(ctx)
}
