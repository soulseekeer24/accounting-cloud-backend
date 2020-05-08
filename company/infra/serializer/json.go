package serializer

import (
	"encoding/json"
	company "piwi-backend-clean/company/domain"
)

type JsonSerializer struct{}

func (s *JsonSerializer) Decode(input []byte) (result *company.Company, err error) {
	result = &company.Company{}

	if err := json.Unmarshal(input, result); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *JsonSerializer) Encode(input interface{}) (bytes []byte, err error) {
	bytes, err = json.Marshal(input)

	if err != nil {
		return nil, err
	}

	return bytes, nil
}
