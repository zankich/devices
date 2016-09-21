package led

import (
	"errors"
	"testing"
)

type fakeDevice struct {
	pin string
	v   int
	err error
}

func (d *fakeDevice) SetValue(pin string, v int) error {
	if d.err != nil {
		return d.err
	}

	d.pin = pin
	d.v = v

	return nil
}

func TestOn(t *testing.T) {
	fakeDevice := &fakeDevice{}
	led := LED{Device: fakeDevice, pin: "D1"}

	err := led.On()
	if err != nil {
		t.Fatal("want: nil got:", err)
	}

	if fakeDevice.pin != "D1" {
		t.Fatal("want: D1 got:", fakeDevice.pin)
	}

	if fakeDevice.v != 1 {
		t.Fatal("want: 1 got:", fakeDevice.v)
	}
}

func TestOnReturnsError(t *testing.T) {
	expectedErr := errors.New("On() failed")
	fakeErrorDevice := &fakeDevice{err: expectedErr}
	led := LED{Device: fakeErrorDevice}

	err := led.On()
	if err == nil || err != expectedErr {
		t.Fatalf("want: %q got: %q", expectedErr, err)
	}
}

func TestOff(t *testing.T) {
	fakeDevice := &fakeDevice{}
	led := LED{Device: fakeDevice, pin: "D5", v: 1}

	err := led.Off()
	if err != nil {
		t.Fatal("want: nil got:", err)
	}

	if fakeDevice.pin != "D5" {
		t.Fatal("want: D1 got:", fakeDevice.pin)
	}

	if fakeDevice.v != 0 {
		t.Fatal("want: 0 got:", fakeDevice.v)
	}
}

func TestOffReturnsError(t *testing.T) {
	expectedErr := errors.New("Off() failed")
	fakeErrorDevice := &fakeDevice{err: expectedErr}
	led := LED{Device: fakeErrorDevice}

	err := led.Off()
	if err == nil || err != expectedErr {
		t.Fatalf("want: %q got: %q", expectedErr, err)
	}
}

func TestToggleLOtoHI(t *testing.T) {
	fakeDevice := &fakeDevice{}
	led := LED{Device: fakeDevice, pin: "D1", v: 0}

	err := led.Toggle()
	if err != nil {
		t.Fatal("want: nil got:", err)
	}

	if fakeDevice.pin != "D1" {
		t.Fatal("want: D1 got:", fakeDevice.pin)
	}

	if fakeDevice.v != 1 {
		t.Fatal("want: 1 got:", fakeDevice.v)
	}
}

func TestToggleHIToLO(t *testing.T) {
	fakeDevice := &fakeDevice{}
	led := LED{Device: fakeDevice, pin: "D4", v: 1}

	err := led.Toggle()
	if err != nil {
		t.Fatal("want: nil got:", err)
	}

	if fakeDevice.pin != "D4" {
		t.Fatal("want: D4 got:", fakeDevice.pin)
	}

	if fakeDevice.v != 0 {
		t.Fatal("want: 0 got:", fakeDevice.v)
	}
}

func TestToggleReturnsError(t *testing.T) {
	expectedErr := errors.New("toggle failed")
	fakeErrorDevice := &fakeDevice{err: expectedErr}
	led := LED{Device: fakeErrorDevice}

	err := led.Toggle()
	if err == nil || err != expectedErr {
		t.Fatalf("want: %q got: %q", expectedErr, err)
	}
}
