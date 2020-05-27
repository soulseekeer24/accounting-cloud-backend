package company

import "context"

func (c *useCase) Delete(ctx context.Context, companyID string) (err error) {
	return c.repository.Delete(ctx , companyID)
}
