package common

type ConnectorConfig struct {
	Spec   string `json:"spec"`
	IpAddr string `json:"ip"`
	Port   uint16 `json:"port"`
}

type ConnectionFactory interface {
	Startable
	Stopable
	CreateConnection() Connection
}
