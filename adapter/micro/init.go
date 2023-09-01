package micro

import (
	"github.com/dianhadi/adapter/port"
)

type Micro struct {
	title string
}

func New(config port.Config) (*Micro, error) {
	micro := Micro{
		title: config.Title,
	}
	return &micro, nil
}
