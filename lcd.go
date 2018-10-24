package lcd

import (
	"errors"
	"time"

	"golang.org/x/exp/io/i2c"
)

const (
	modeCommand = 0
	modeData    = 1

	delay = 5 * time.Microsecond
)

var (
	errOutOfRows = errors.New("Row out of range")
	errOutOfCols = errors.New("Column out of range")
)

type LCD struct {
	Bus     string
	Address int
	Rows    uint
	Cols    uint

	backlight byte
	device    *i2c.Device
}

func New() *LCD {
	return &LCD{}
}

func (l *LCD) Open() (err error) {
	l.device, err = i2c.Open(&i2c.Devfs{Dev: l.Bus}, l.Address)
	return
}

func (l *LCD) Close() {
	l.device.Close()
}

func (l *LCD) Init() error {
	l.Backlight(true)
	for _, b := range []byte{0x33, 0x32, 0x06, 0x0C, 0x28, 0x01} {
		if err := l.write(b, modeCommand); err != nil {
			return err
		}
	}
	time.Sleep(delay)
	return nil
}

func (l *LCD) Clear() error {
	return l.write(0x01, modeCommand)
}

func (l *LCD) Print(row uint, msg string) error {
	if len(msg) > int(l.Cols) {
		return errOutOfCols
	}
	if row > l.Rows {
		return errOutOfRows
	}
	var line byte
	switch row {
	case 1:
		line = 0x80
	case 2:
		line = 0xC0
	case 3:
		line = 0x94
	case 4:
		line = 0xD4
	default:
		return errOutOfRows
	}
	if err := l.write(line, modeCommand); err != nil {
		return err
	}
	for _, b := range []byte(msg) {
		if err := l.write(b, modeData); err != nil {
			return err
		}
	}
	return nil
}

func (l *LCD) Backlight(enabled bool) {
	if enabled {
		l.backlight = 0x08
	} else {
		l.backlight = 0x00
	}
}

func (l *LCD) write(data, mode byte) error {
	high := mode | (data & 0xF0) | l.backlight
	low := mode | ((data << 4) & 0xF0) | l.backlight

	if err := l.device.Write([]byte{high}); err != nil {
		return err
	}
	if err := l.toggleEnable(high); err != nil {
		return err
	}
	if err := l.device.Write([]byte{low}); err != nil {
		return err
	}
	if err := l.toggleEnable(low); err != nil {
		return err
	}

	return nil
}

func (l *LCD) toggleEnable(data byte) error {
	time.Sleep(delay)
	if err := l.device.Write([]byte{data | 0x04}); err != nil {
		return err
	}
	time.Sleep(delay)
	if err := l.device.Write([]byte{data & 0xFB}); err != nil {
		return err
	}
	time.Sleep(delay)
	return nil
}
