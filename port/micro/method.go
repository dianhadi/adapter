package micro

import "fmt"

func (m Micro) Send(s string) error {
	fmt.Println("Send with", m.title, s)
	return nil
}

func (m Micro) Receive(i int) (int, error) {
	fmt.Println("Receive with", m.title, i)
	return i, nil
}
