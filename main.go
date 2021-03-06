package main

import (
	"fmt"
	"golang.org/x/exp/io/i2c"
	"time"
)

const (
	TEMP_NO_HOLD = byte(0xF3)
	HUMIDITY     = byte(0xF5)
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func waitForRead() {
	time.Sleep(500 * time.Millisecond)
}

func getTemp(d i2c.Device) (float64, float64) {
	err := d.Write([]byte{TEMP_NO_HOLD})
	check(err)

	waitForRead()

	response := make([]byte, 2)
	check(d.Read(response))

	temp := (int64(response[0])*256 + int64(response[1])) & 0xFFFC
	tempC := -46.85 + (175.72 * float64(temp) / 65536.0)
	tempF := tempC*1.8 + 32

	// fmt.Printf("%.2fºF\n", tempF)
	return tempF, tempC
}

func getRelativeHumidity(d i2c.Device) int64 {
	err := d.Write([]byte{HUMIDITY})
	check(err)

	waitForRead()

	response := make([]byte, 2)
	check(d.Read(response))

	data := (int64(response[0])*256 + int64(response[1])) & 0xFFFC
	humidity := ((125 * data) / 65536) - 6

	return humidity
}

func main() {
	d, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, 0x40)
	check(err)

	tempF, _ := getTemp(*d)
	humidity := getRelativeHumidity(*d)

	fmt.Printf("temp=%.2f humidity=%d\n", tempF, humidity)
}
