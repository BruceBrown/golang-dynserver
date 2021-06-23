package common

type ConnectionFactory interface {
	Startable
	Stopable
	CreateConnection() Connection
}
