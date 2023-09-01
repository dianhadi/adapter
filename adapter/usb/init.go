package usb

import (
	"github.com/dianhadi/adapter/port"
)

type Usb struct {
	title string
}

func New(config port.Config) (*Usb, error) {
	usb := Usb{
		title: config.Title,
	}
	return &usb, nil
}
