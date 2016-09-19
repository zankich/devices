package led

import (
	"errors"
	"testing"
)

func TestOff(t *testing.T) {
	fakeDevice := &fakeDevice{}
	led := LED{device: fakeDevice, pin: "D5", v: 1}

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
	led := LED{device: fakeErrorDevice}

	err := led.Off()
	if err == nil || err != expectedErr {
		t.Fatalf("want: %q got: %q", expectedErr, err)
	}
}
