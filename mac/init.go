package mac

import "github.com/dianhadi/adapter/adapter"

type Mac struct {
	thunderbolt adapter.Adapter
}

func New(thunderbolt adapter.Adapter) (*Mac, error) {
	mac := Mac{
		thunderbolt: thunderbolt,
	}
	return &mac, nil
}
