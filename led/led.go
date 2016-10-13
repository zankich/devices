package led

import (
	"github.com/goiot/exp/gpio"
	"github.com/goiot/exp/gpio/driver"
)

const (
	lo = 0
	hi = 1
)

type Device interface {
	SetValue(pin string, v int) error
	SetPWMValue(ping string, v int) error
}

type LED struct {
	Device Device
	pin    string
	v      int
}

func New(o driver.Opener, pin string) (*LED, error) {
	dev, err := gpio.Open(o)
	if err != nil {
		return nil, err
	}

	return &LED{
		Device: dev,
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

func (l *LED) SetBrightness(v int) error {
	return l.Device.SetPWMValue(l.pin, v)
}

func (l *LED) setDeviceValue(val int) error {
	l.v = val
	return l.Device.SetValue(l.pin, val)
}
