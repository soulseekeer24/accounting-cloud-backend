package company

type Serializer interface {
	Decode(input []byte) (company *Company,err error)
	Encode(input interface{}) (bytes []byte, err error)
}