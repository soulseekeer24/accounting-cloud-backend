package core

import (
	"context"
	"piwi-backend-clean/profiles/core/domains/profiles"
)

// Module for the users profile domian administration
type Module struct {
	profileService *profiles.Service
}

func BuildModule(profileStore profiles.Store) *Module {
	service := profiles.NewService(profileStore)
	m := Module{profileService: service}
	return &m
}

func (m *Module) CreateNewUserProfile(ctx context.Context,accountID string, p *profiles.Profile) (ID string, err error) {
	return m.profileService.CreateProfile(ctx,accountID, p)
}

func (m *Module) GetAccountProfile(ctx context.Context, accountID string) (profile *profiles.Profile, err error) {
	return m.profileService.GetProfileByAccountID(ctx, accountID)
}

func (m *Module) ValidateContact(ctx context.Context, accountID string) (sucess bool, err error) {
	return m.profileService.ValidateMainContactInfo(ctx, accountID, true)
}

func (m *Module) GetProfilebyID(ctx context.Context, profileId string) (profile *profiles.Profile, err error) {
	return m.profileService.GetProfileByID(ctx, profileId)
}

func (m *Module) UpdateProfile(ctx context.Context, ID string, update profiles.Profile) (ok bool, err error) {
	return m.profileService.UpdateProfile(ctx, ID, update)
}
