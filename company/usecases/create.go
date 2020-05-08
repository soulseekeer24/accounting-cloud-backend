package company

import (
	"piwi-backend-clean/company/domain"
	"time"
)

func (c *useCase) Create(company *company.Company) (companyStored *company.Company, err error) {

	if _, err = company.IsValid(); err != nil {
		return nil, err
	}

	company.CreatedAt = time.Now().Unix()

	if companyStored, err = c.repository.Store(company); err != nil {
		return nil, err
	}

	return companyStored, nil
}
