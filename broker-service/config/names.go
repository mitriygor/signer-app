package config

const (
	ReqCount            = "broker_count"
	ErrorCount          = "broker_error_count"
	LoggerReqCount      = "logger_count"
	LoggerErrorCount    = "logger_error_count"
	ListenerReqCount    = "listener_count"
	ListenerErrorCount  = "listener_error_count"
	SignerReqCount      = "signer_count"
	SignerErrorCount    = "signer_error_count"
	KeyKeeperReqCount   = "key_keeper_count"
	KeyKeeperErrorCount = "key_keeper_error_count"
)

var Counts = []string{ReqCount, ErrorCount, ListenerReqCount, ListenerErrorCount, LoggerReqCount, LoggerErrorCount, SignerReqCount, SignerErrorCount, KeyKeeperReqCount, KeyKeeperErrorCount}
