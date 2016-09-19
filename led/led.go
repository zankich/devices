package led

import (
	"github.com/goiot/exp/gpio"
	"github.com/goiot/exp/gpio/driver"
)

const (
	lo = 0
	hi = 1
)

type device interface {
	SetValue(pin string, v int) error
}

type LED struct {
	device device
	pin    string
	v      int
}

func NewLED(o driver.Opener, pin string) (*LED, error) {
	dev, err := gpio.Open(o)
	if err != nil {
		return nil, err
	}

	return &LED{
		device: dev,
		pin:    pin,
	}, nil
}

func (l *LED) On() error {
	return l.setDeviceValue(hi)
}

func (l *LED) Off() error {
	return l.setDeviceValue(lo)
}

func (l *LED) Toggle() error {
	return l.setDeviceValue(l.v ^ 1)
}

func (l *LED) setDeviceValue(val int) error {
	l.v = val
	return l.device.SetValue(l.pin, val)
}
