package company

func (c *useCase) Delete(companyID string) (err error) {
	return c.repository.Delete(companyID)
}
