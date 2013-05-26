package limitlessled

import (
	"fmt"
	"net"
	"time"
)

/*
 * Control codes for the LimitlessLED bulbs, taken from Taken from http://forum.micasaverde.com/index.php?topic=14471.0
 * 35 00 55 - All On
 * 39 00 55 - All Off
 * 3C 00 55 - Brightness Up
 * 34 00 55 - Brightness Down  (There are ten steps between min and max)
 * 3E 00 55 - Warmer
 * 3F 00 55 - Cooler  (There are ten steps between warmest and coolest)
 * 38 00 55 - Zone 1 On
 * 3B 00 55 - Zone 1 Off
 * 3D 00 55 - Zone 2 On
 * 33 00 55 - Zone 2 Off
 * 37 00 55 - Zone 3 On
 * 3A 00 55 - Zone 3 Off
 * 32 00 55 - Zone 4 On
 * 36 00 55 - Zone 4 Off
 * B5 00 55 - All On Full    (Send >=100ms after All On)
 * B8 00 55 - Zone 1 Full  (Send >=100ms after Zone 1 On)
 * BD 00 55 - Zone 2 Full  (Send >=100ms after Zone 2 On)
 * B7 00 55 - Zone 3 Full  (Send >=100ms after Zone 3 On)
 * B2 00 55 - Zone 4 Full  (Send >=100ms after Zone 4 On)
 * B9 00 55 - All Nightlight         (Send >=100ms after All Off)
 * BB 00 55 - Zone 1 Nightlight  (Send >=100ms after Zone 1 Off)
 * B3 00 55 - Zone 2 Nightlight  (Send >=100ms after Zone 2 Off)
 * BA 00 55 - Zone 3 Nightlight  (Send >=100ms after Zone 3 Off)
 * B6 00 55 - Zone 4 Nightlight  (Send >=100ms after Zone 4 Off)
 */

var (
	ALL_ON          = []byte{0x35, 0x00, 0x55}
	ALL_OFF         = []byte{0x39, 0x00, 0x55}
	BRIGHTNESS_UP   = []byte{0x3C, 0x00, 0x55}
	BRIGHTNESS_DOWN = []byte{0x34, 0x00, 0x55}
	WARMER          = []byte{0x3E, 0x00, 0x55}
	COOLER          = []byte{0x3F, 0x00, 0x55}
	ALL_ON_FULL     = []byte{0xB5, 0x00, 0x55}
	ALL_NIGHTLIGHT  = []byte{0xB9, 0x00, 0x55}
)

type Bridge struct {
	*net.UDPConn
}

type Bulb struct {
	Brightness  int
	Temperature int
	IsOn        bool
}

func (bulb Bulb) String() string {
	return fmt.Sprintf("Bulb{Brightness: %d, Temperature:%d, IsOn: %s", bulb.Brightness, bulb.Temperature, bulb.IsOn)
}

func Dial(host string) (*Bridge, error) {
	addr, err := net.ResolveUDPAddr("udp4", host)

	if err != nil {
		return nil, err
	}

	s, err := net.DialUDP("udp4", nil, addr)
	return &Bridge{s}, err
}

func (bridge *Bridge) SendCommand(command []byte) {
	bridge.Write(command)
	time.Sleep(time.Millisecond * 50)
}

func (bridge *Bridge) Set(bulb, other Bulb) Bulb {

	if bulb.IsOn && !other.IsOn {
		bridge.SendCommand(ALL_OFF)
		return Bulb{
			Brightness:  bulb.Brightness,
			Temperature: bulb.Temperature,
			IsOn:        false}
	}

	if !bulb.IsOn && other.IsOn {
		bridge.SendCommand(ALL_ON)
	}

	brightnessSteps := other.Brightness - bulb.Brightness
	temperatureSteps := other.Temperature - bulb.Temperature

	if brightnessSteps < 0 {
		for i := 0; i > brightnessSteps; i-- {
			bridge.SendCommand(BRIGHTNESS_DOWN)
		}
	} else {
		for i := 0; i < brightnessSteps; i++ {
			bridge.SendCommand(BRIGHTNESS_UP)
		}
	}

	if temperatureSteps < 0 {
		for i := 0; i > temperatureSteps; i-- {
			bridge.SendCommand(WARMER)
		}
	} else {
		for i := 0; i < temperatureSteps; i++ {
			bridge.SendCommand(COOLER)
		}
	}

	return other
}
