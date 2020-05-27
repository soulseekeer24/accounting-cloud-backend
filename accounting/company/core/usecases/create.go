package company

import (
	company "accounting/company/core/domain"
	"context"
	"time"
)

func (c *useCase) Create(ctx context.Context, company *company.Company) (companyStored *company.Company, err error) {

	if _, err = company.IsValid(); err != nil {
		return nil, err
	}

	company.CreatedAt = time.Now().Unix()

	if companyStored, err = c.repository.Store(ctx, company); err != nil {
		return nil, err
	}

	return companyStored, nil
}

