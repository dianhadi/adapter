package mac

import "github.com/dianhadi/adapter/port"

type Mac struct {
	thunderbolt port.Port
}

func New(thunderbolt port.Port) (*Mac, error) {
	mac := Mac{
		thunderbolt: thunderbolt,
	}
	return &mac, nil
}
