package led

import (
	"testing"

	"github.com/goiot/exp/gpio/driver"
)

type conn struct {
	pin string
	v   int
}

type opener struct{}

func (o opener) Open() (driver.Conn, error) {
	return &conn{}, nil
}

func (c *conn) Value(pin string) (int, error)                       { panic("not implemented") }
func (c *conn) SetDirection(pin string, dir driver.Direction) error { panic("not implemented") }
func (c *conn) Map(virtual string, physical int)                    { panic("not implemented") }
func (c *conn) Close() error                                        { panic("not implemented") }
func (c *conn) SetValue(pin string, v int) error {
	c.pin = pin
	c.v = v

	return nil
}

func TestToggle(t *testing.T) {
	fakeDevice := &conn{}
	led := Led{Device: fakeDevice, Pin: "D1"}

	led.Toggle()
	if fakeDevice.pin != "D1" {
		t.Fatal("want: D1 got:", fakeDevice.pin)
	}

	if fakeDevice.v != 1 {
		t.Fatal("want: 1 got:", fakeDevice.v)
	}

	led.Toggle()
	if fakeDevice.v != 0 {
		t.Fatal("want: 0 got:", fakeDevice.v)
	}
}
