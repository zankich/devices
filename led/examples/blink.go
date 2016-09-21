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

	l, err := led.NewLED(i2co, "D2")
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println("toggle")
		if err := l.Toggle(); err != nil {
			panic(err)
		}
		time.Sleep(200 * time.Millisecond)
	}
}
