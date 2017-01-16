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

## Useful Resources

- http://files.pine64.org/doc/Pine%20A64%20Schematic/Pine%20A64%20Pin%20Assignment%20160119.pdf
- http://wiki.pine64.org/index.php/POT
- https://dave.cheney.net/2014/08/03/tinyterm-a-silly-terminal-emulator-written-in-go
