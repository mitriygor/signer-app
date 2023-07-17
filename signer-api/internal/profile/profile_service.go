package profile

type Service interface {
	GetAllProfiles(args Args) ([]*Profile, error)
	SignAllProfilesWithParams(signPayload SignPayload)
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
	ps.profileRepo.SignAll()
}

func (ps *profileService) SignAllProfilesWithParams(signPayload SignPayload) {
	ps.profileRepo.SignAllWithParams(signPayload)
}
