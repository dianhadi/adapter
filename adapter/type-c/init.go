package typec

import (
	"github.com/dianhadi/adapter/port"
)

type TypeC struct {
	title string
}

func New(config port.Config) (*TypeC, error) {
	typeC := TypeC{
		title: config.Title,
	}
	return &typeC, nil
}
