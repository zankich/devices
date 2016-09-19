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

	fmt.Println("turn on the light in 2s...")
	time.Sleep(2 * time.Second)
	if err := l.On(); err != nil {
		panic(err)
	}

	fmt.Println("turn off the light in 2s...")
	time.Sleep(2 * time.Second)
	if err := l.Off(); err != nil {
		panic(err)
	}

	fmt.Println("in 2s, toggle the light every .5s...")
	time.Sleep(2 * time.Second)
	for {
		fmt.Println("toggle")
		if err := l.Toggle(); err != nil {
			panic(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
