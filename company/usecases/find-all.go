package company

import company "piwi-backend-clean/company/domain"

func (c *UseCase) FindAll() (companies []company.Company, err error) {
	return c.repository.GetAll()
}
