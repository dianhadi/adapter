package typec

import "fmt"

func (t TypeC) Send(s string) error {
	fmt.Println("Send with", t.title, s)
	return nil
}

func (t TypeC) Receive(i int) (int, error) {
	fmt.Println("Receive with", t.title, i)
	return i - 1, nil
}
