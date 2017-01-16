# Pine 64 - Temperature / Humidity Reader

This code reads the temperature and relative humidity from a Pine64 I2C PMSD001 sensor.

It's wired up via the Pi 2 bus as follows:

- Pin 3 - SDA
- Pin 4 - DC 5V
- Pin 5 - SCL
- Pin 6 - GND

![](wiring.jpg)

## Usage

```
$ go build -o reader main.go
$ ./reader
temp=32.32 humidity=73
```
