package main

import (
	"fmt"
	"time"

	"github.com/goiot/devices/led"
	"github.com/zankich/hal/grovepi"
	"golang.org/x/exp/io/i2c"
)

func main() {
	i2co, err := grovepi.New(&i2c.Devfs{Dev: "/dev/i2c-1"})
	if err != nil {
		panic(err)
	}

	l, err := led.New(i2co, "D3")
	if err != nil {
		panic(err)
	}

	fmt.Println("partially dimming the led...")
	time.Sleep(2 * time.Second)
	l.SetBrightness(64)

	fmt.Println("setting led to full brightness...")
	time.Sleep(2 * time.Second)
	l.SetBrightness(255)

	fmt.Println("partially dimming the led...")
	time.Sleep(2 * time.Second)
	l.SetBrightness(64)

	fmt.Println("turning off led...")
	time.Sleep(2 * time.Second)
	l.SetBrightness(0)

	fmt.Println("Done!")
}
