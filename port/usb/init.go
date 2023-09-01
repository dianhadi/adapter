package usb

import (
	"github.com/dianhadi/adapter/adapter"
)

type Usb struct {
	title string
}

func New(config adapter.Config) (*Usb, error) {
	usb := Usb{
		title: config.Title,
	}
	return &usb, nil
}
