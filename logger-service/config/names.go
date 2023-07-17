package config

const (
	ReqCount            = "logger_count"
	ErrorCount          = "logger_error_count"
	ListenerReqCount    = "listener_count"
	ListenerErrorCount  = "listener_error_count"
	BrokerReqCount      = "broker_count"
	BrokerErrorCount    = "broker_error_count"
	SignerReqCount      = "signer_count"
	SignerErrorCount    = "signer_error_count"
	KeyKeeperReqCount   = "key_keeper_count"
	KeyKeeperErrorCount = "key_keeper_error_count"
)

var Counts = []string{ReqCount, ErrorCount, ListenerReqCount, ListenerErrorCount, BrokerReqCount, BrokerErrorCount, SignerReqCount, SignerErrorCount, KeyKeeperReqCount, KeyKeeperErrorCount}
