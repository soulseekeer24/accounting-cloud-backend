package company

func (c *UseCase) Delete(companyID string) (err error) {
	return c.repository.Delete(companyID)
}
