package log_entry

import "go.mongodb.org/mongo-driver/mongo"

type Service interface {
	InsertLogEntry(logEntry LogEntry) error
	GetAllLogEntries() ([]*LogEntry, error)
	GetLogEntry(id string) (*LogEntry, error)
	UpdateLogEntry(logEntry LogEntry) (*mongo.UpdateResult, error)
	GetCountService(countName string) int
}

type LogEntryService struct {
	logEntryRepo Repository
}

func NewLogEntryService(repo Repository) Service {
	return &LogEntryService{
		logEntryRepo: repo,
	}
}

func (ls *LogEntryService) InsertLogEntry(logEntry LogEntry) error {
	return ls.logEntryRepo.Insert(logEntry)
}

func (ls *LogEntryService) GetAllLogEntries() ([]*LogEntry, error) {
	return ls.logEntryRepo.All()
}

func (ls *LogEntryService) GetLogEntry(id string) (*LogEntry, error) {
	return ls.logEntryRepo.GetOne(id)
}

func (ls *LogEntryService) UpdateLogEntry(logEntry LogEntry) (*mongo.UpdateResult, error) {
	return ls.logEntryRepo.Update(logEntry)
}

func (ls *LogEntryService) GetCountService(countName string) int {
	return ls.logEntryRepo.GetCount(countName)
}
