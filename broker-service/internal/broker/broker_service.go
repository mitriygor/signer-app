package broker

type Service interface {
	IncrCountService()
	GetCountService() int
	LogEventViaRabbitService(l LogPayload)
}

type brokerService struct {
	brokerRepo Repository
}

func NewBrokerService(repo Repository) Service {
	return &brokerService{
		brokerRepo: repo,
	}
}

func (s *brokerService) IncrCountService() {
	s.brokerRepo.IncrCount()
}

func (s *brokerService) GetCountService() int {
	return s.brokerRepo.GetCount()
}

func (s *brokerService) LogEventViaRabbitService(l LogPayload) {
	s.IncrCountService()
	s.brokerRepo.LogEventViaRabbit(l)
}
