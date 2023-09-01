package adapter

type Config struct {
	Title string
}

type Adapter interface {
	Send(s string) error
	Receive(i int) (int, error)
}
