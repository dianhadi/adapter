package usb

import "fmt"

func (u Usb) Send(s string) error {
	fmt.Println("Send with", u.title, s)
	return nil
}

func (u Usb) Receive(i int) (int, error) {
	fmt.Println("Receive with", u.title, i)
	return i + 1, nil
}
