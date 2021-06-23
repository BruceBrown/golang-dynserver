package common

type Runable interface {
	Startable
	Stopable
	Run() bool
}
