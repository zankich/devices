package led

import (
	"golang.org/x/exp/io/gpio"
	"golang.org/x/exp/io/gpio/driver"
)

type Led struct {
	Device *gpio.Device
	on     bool
	pin    string
}

func NewLED(o driver.Opener, pin string) (*Led, error) {
	dev, err := gpio.Open(o)
	if err != nil {
		return nil, err
	}

	return &Led{
		Device: dev,
		pin:    pin,
	}, nil
}

func (l *Led) Toggle() error {
	if l.on {
		if err := l.Device.SetValue(l.pin, 0); err != nil {
			return err
		}
		l.on = false
	} else {
		if err := l.Device.SetValue(l.pin, 1); err != nil {
			return err
		}
		l.on = true
	}
	return nil
}
