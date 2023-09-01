package micro

import (
	"github.com/dianhadi/adapter/adapter"
)

type Micro struct {
	title string
}

func New(config adapter.Config) (*Micro, error) {
	micro := Micro{
		title: config.Title,
	}
	return &micro, nil
}
