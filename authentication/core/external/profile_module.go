package external

import "piwi-backend-clean/authentication/core/dto"

type ProfileModule interface {
	GetProfileByAccountID(accID string) (profile *dto.Profile, err error)

	CreateProfile(profile *dto.Profile) (success bool, err error)

	ValidateEmail(accountId string) (success bool, err error)
}
