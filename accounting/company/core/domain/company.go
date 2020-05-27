package company

type Company struct {
	ID                      string `json:"id" bson:"_id"`
	Name                    string `json:"name" bson:"name"`
	TaxIdentificationNumber string `json:"tax_identification_number" bson:"tax_identification_number"`
	CreatedAt               int64  `json:"created_at" bson:"created_at"`
}

func (c *Company) IsValid() (bool, error) {



	return true, nil
}
