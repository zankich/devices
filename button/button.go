package button

import (
	"github.com/goiot/exp/gpio"
	"github.com/goiot/exp/gpio/driver"
)

type Device interface {
	GetValue(pin string) (int, error)
}

type Button struct {
	Device Device
	Pin    string
}

func New(o driver.Opener, pin string) (*Button, error) {
	dev, err := gpio.Open(o)
	if err != nil {
		return nil, err
	}

	return &Button{
		Device: dev,
		Pin:    pin,
	}, nil
}

func (b *Button) IsOn() (bool, error) {
	value, err := b.Device.GetValue(b.Pin)
	if err != nil {
		return false, err
	}

	return value == 1, nil
}
