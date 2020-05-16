package interfaces

import (
	bill "piwi-backend-clean/bill/domain"
)

type Serializer interface {
	Decode(input []byte) (bill *bill.Bill, err error)
	Encode(input interface{}) (bytes []byte, err error)
}
