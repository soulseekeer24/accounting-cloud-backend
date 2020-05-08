package company

type Serializer interface {
	Decode(input []byte) (company *Company,err error)
	Encode(input *Company) (companyBytes *[]byte, err error)
}