package config

const (
	ReqCount            = "listener_count"
	ErrorCount          = "listener_error_count"
	BrokerReqCount      = "broker_count"
	BrokerErrorCount    = "broker_error_count"
	LoggerReqCount      = "logger_count"
	LoggerErrorCount    = "logger_error_count"
	SignerReqCount      = "signer_count"
	SignerErrorCount    = "signer_error_count"
	KeyKeeperReqCount   = "key_keeper_count"
	KeyKeeperErrorCount = "key_keeper_error_count"
)

var Counts = []string{ReqCount, ErrorCount, BrokerReqCount, BrokerErrorCount, LoggerReqCount, LoggerErrorCount, SignerReqCount, SignerErrorCount, KeyKeeperReqCount, KeyKeeperErrorCount}
