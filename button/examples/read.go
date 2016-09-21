package main

import (
	"fmt"
	"time"

	"github.com/goiot/devices/button"
	"github.com/zankich/hal/grovepi"
	"golang.org/x/exp/io/i2c"
)

func main() {
	i2co, err := grovepi.New(&i2c.Devfs{Dev: "/dev/i2c-1"})
	if err != nil {
		panic(err)
	}

	b, err := button.New(i2co, "D2")
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println(b.IsOn())
		time.Sleep(200 * time.Millisecond)
	}
}
