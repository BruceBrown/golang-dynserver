package interfaces

type Runable interface {
	Startable
	Stopable
	Run() bool
}
