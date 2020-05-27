package interfaces

import (
	company "accounting/company/core/domain"
)

type Serializer interface {
	Decode(input []byte) (company *company.Company,err error)
	Encode(input interface{}) (bytes []byte, err error)
}