package company

import company "piwi-backend-clean/company/domain"

func (c *useCase) FindAll() (companies []company.Company, err error) {
	return c.repository.GetAll()
}
