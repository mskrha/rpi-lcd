## rpi-lcd

### Description
Golang library for accessing I2C LCD screen connected to the RaspberryPi.

### Installation
`go get github.com/mskrha/rpi-lcd`

### Usage
```go
package main

import (
	"fmt"
	"time"

	"github.com/mskrha/rpi-lcd"
)

func main() {
	l := lcd.New(lcd.LCD{Bus: "/dev/i2c-1", Address: 0x27, Rows: 2, Cols: 16, Backlight: true})

	if err := l.Init(); err != nil {
		panic(err)
	}

	for {
		if err := l.Print(1, 0, time.Now().Format("02.01.")); err != nil {
			fmt.Println(err)
			return
		}
		if err := l.Print(1, 8, time.Now().Format("15:04:05")); err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(time.Second)
	}
}
```
