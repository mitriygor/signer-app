package private_key

type Service interface {
	GetAllPrivateKeys(args Args) ([]*PrivateKey, error)
}

type privateKeyService struct {
	privateKeyRepo Repository
}

func NewPrivateKeyService(repo Repository) Service {
	return &privateKeyService{
		privateKeyRepo: repo,
	}
}

func (s *privateKeyService) GetAllPrivateKeys(args Args) ([]*PrivateKey, error) {
	return s.privateKeyRepo.GetAll(args)
}
