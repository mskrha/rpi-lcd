## rpi-lcd

### Description
Golang library for accessing I2C LCD screen connected to the RaspberryPi.

### Tested LCDs:
* [Geekcreit® IIC/I2C 1602 Yellow Green Backlight LCD Display Module For Arduino](https://www.banggood.com/IICI2C-1602-Yellow-Green-Backlight-LCD-Display-Module-For-Arduino-p-950728.html)
* [Geekcreit® IIC I2C 2004 204 20 x 4 Character LCD Display Screen Module Blue For Arduino](https://www.banggood.com/IIC-I2C-2004-204-20-x-4-Character-LCD-Display-Module-Blue-p-908616.html)

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

### Notes
Inspired by the [Adafruit Liquid Crystal library](https://github.com/adafruit/Adafruit_LiquidCrystal).
