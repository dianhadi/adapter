package main

import (
	"fmt"

	"github.com/dianhadi/adapter/adapter/micro"
	typeC "github.com/dianhadi/adapter/adapter/type-c"
	"github.com/dianhadi/adapter/adapter/usb"
	"github.com/dianhadi/adapter/mac"
	"github.com/dianhadi/adapter/port"
)

func main() {
	configUsb := port.Config{
		Title: "USB",
	}

	usb, err := usb.New(configUsb)
	if err != nil {
		fmt.Println(err)
	}

	configMicro := port.Config{
		Title: "Micro",
	}

	micro, err := micro.New(configMicro)
	if err != nil {
		fmt.Println(err)
	}

	configTypeC := port.Config{
		Title: "Type C",
	}

	typeC, err := typeC.New(configTypeC)
	if err != nil {
		fmt.Println(err)
	}

	mac1, err := mac.New(usb)
	if err != nil {
		fmt.Println(err)
	}

	mac2, err := mac.New(micro)
	if err != nil {
		fmt.Println(err)
	}

	mac3, err := mac.New(typeC)
	if err != nil {
		fmt.Println(err)
	}

	mac1.SendAndReceive("foo", 123)
	mac2.SendAndReceive("bar", 456)
	mac3.SendAndReceive("baz", 789)
}
