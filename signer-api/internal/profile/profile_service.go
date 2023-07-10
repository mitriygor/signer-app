package profile

import "fmt"

type Service interface {
	GetAllProfiles(args Args) ([]*Profile, error)
	SignAllProfiles()
}

type profileService struct {
	profileRepo Repository
}

func NewProfileService(repo Repository) Service {
	return &profileService{
		profileRepo: repo,
	}
}

func (ps *profileService) GetAllProfiles(args Args) ([]*Profile, error) {
	return ps.profileRepo.GetAll(args)
}

func (ps *profileService) SignAllProfiles() {
	fmt.Printf("\nSignAllProfiles\n")
	ps.profileRepo.SignAll()
}
