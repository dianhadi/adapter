package mac

import "fmt"

func (m Mac) SendAndReceive(s string, i int) error {
	err := m.thunderbolt.Send(s)
	if err != nil {
		return err
	}
	x, err := m.thunderbolt.Receive(i)
	if err != nil {
		return err
	}
	fmt.Println(i, "Processed by Mac", x)
	return nil
}
