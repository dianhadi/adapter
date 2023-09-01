package typec

import (
	"github.com/dianhadi/adapter/adapter"
)

type TypeC struct {
	title string
}

func New(config adapter.Config) (*TypeC, error) {
	typeC := TypeC{
		title: config.Title,
	}
	return &typeC, nil
}
