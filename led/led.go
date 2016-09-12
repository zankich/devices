package led

import (
	"github.com/goiot/exp/gpio"
	"github.com/goiot/exp/gpio/driver"
)

type Device interface {
	SetValue(pin string, v int) error
}

type Led struct {
	Device Device
	on     bool
	Pin    string
}

func NewLED(o driver.Opener, pin string) (*Led, error) {
	dev, err := gpio.Open(o)
	if err != nil {
		return nil, err
	}

	return &Led{
		Device: dev,
		Pin:    pin,
	}, nil
}

func (l *Led) Toggle() error {
	if l.on {
		if err := l.Device.SetValue(l.Pin, 0); err != nil {
			return err
		}
		l.on = false
	} else {
		if err := l.Device.SetValue(l.Pin, 1); err != nil {
			return err
		}
		l.on = true
	}
	return nil
}
