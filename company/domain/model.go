package company

import "time"

type Company struct {
	ID                      string `json:"id" bson:"_id"`
	Name                    string `json:"name" bson:"name"`
	TaxIdentificationNumber string `json:"tax_identification_number" bson:"tax_identification_number"`
	CreatedAt               int64  `json:"created_at" bson:"created_at"`
}

func (c *Company) IsValid() (bool, error) {

	c.



	c.CreatedAt = time.Now().Unix()
	return true, nil
}
