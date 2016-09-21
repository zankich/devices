package led

import (
	"errors"
	"testing"
)

func TestOn(t *testing.T) {
	fakeDevice := &fakeDevice{}
	led := LED{device: fakeDevice, pin: "D1"}

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
	led := LED{device: fakeErrorDevice}

	err := led.On()
	if err == nil || err != expectedErr {
		t.Fatalf("want: %q got: %q", expectedErr, err)
	}
}
