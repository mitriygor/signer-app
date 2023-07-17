package broker

type Service interface {
	IncrCountService(countName string)
	GetCountService(countName string) int
	LogEventViaRabbitService(l RequestPayload)
}

type brokerService struct {
	brokerRepo Repository
}

func NewBrokerService(repo Repository) Service {
	return &brokerService{
		brokerRepo: repo,
	}
}

func (s *brokerService) LogEventViaRabbitService(l RequestPayload) {
	s.brokerRepo.LogEventViaRabbit(l)
}

func (s *brokerService) IncrCountService(countName string) {
	s.brokerRepo.IncrCount(countName)
}

func (s *brokerService) GetCountService(countName string) int {
	return s.brokerRepo.GetCount(countName)
}
