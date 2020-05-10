package tests

import (
	"piwi-backend-clean/profiles/core/domains/profiles"
	"testing"


	"golang.org/x/net/context"
)

var store = &MuckProfileStore{}
var ctx = context.TODO()

type ModuleTestFunction func(s *users.Module, t *testing.T)

type TestCase struct {
	Name     string
	Callback ModuleTestFunction
}

func ModuleSuite(s *users.Module, t *testing.T) {
	tc := []TestCase{
		{Name: "Attemp create with no accountID", Callback: AttempToCreateWithMissingAccount},
		{Name: "Create profile succed", Callback: CreateProfileSucceded},
		{Name: "Create profile without contact info", Callback: CreateProfileWithoutContactInfo},
	}

	for _, tCase := range tc {
		t.Run(tCase.Name, func(t *testing.T) {
			tCase.Callback(s, t)
		})
	}
}

func AttempToCreateWithMissingAccount(s *users.Module, t *testing.T) {
	badProfile := profiles.Profile{}

	_, err := s.CreateNewUserProfile(ctx, &badProfile)
	if err == nil {
		t.Error("It should have returned a error")
	}

	switch te := err.(type) {
	case profiles.MissingAccountIDError:
		return
	default:
		t.Errorf("expected [%v] to be [MissingAccountIDError]", te)
	}
}

func CreateProfileSucceded(s *users.Module, t *testing.T) {
	profile := profiles.Profile{
		AccountID: "1",
		Contacts: []profiles.ContactInfo{
			{Value: "migue@test.com", Channel: profiles.Email},
		},
		FirstName: "Miguel",
		LastName:  "Olivarez"}

	_, err := s.CreateNewUserProfile(ctx, &profile)
	if err != nil {
		t.Errorf("It should have returned a no error got %v", err)
	}

}

func CreateProfileWithoutContactInfo(s *users.Module, t *testing.T) {
	profile := profiles.Profile{
		AccountID: "1",

		FirstName: "Miguel",
		LastName:  "Olivarez"}

	_, err := s.CreateNewUserProfile(ctx, &profile)
	if err == nil {
		t.Error("It should have returned a missing conctact error  ")
	}

	switch te := err.(type) {
	case profiles.NoContactsOnProfileError:
		return
	default:
		t.Errorf("expected [%v] to be [MissingAccountIDError]", te)
	}

}

func TestModule(t *testing.T) {

	module := users.BuildModule(store)
	// module.ConnectToProfiles(profileMuck)
	ModuleSuite(module, t)
}
