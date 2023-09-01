package port

type Config struct {
	Title string
}

type Port interface {
	Send(s string) error
	Receive(i int) (int, error)
}
