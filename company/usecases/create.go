package company

import (
	company "piwi-backend-clean/company/domain"
)


func (c *UseCase) Create(company *company.Company) (companyStored *company.Company, err error) {

	if _, err = company.IsValid(); err != nil {
		return nil, err
	}

	if companyStored, err = c.repository.Store(company); err != nil {
		return nil, err
	}

	return companyStored, nil
}
