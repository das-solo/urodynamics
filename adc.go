package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	r := raspi.NewAdaptor()

	// default bus/address
	bpm180 := i2c.NewADS1115Driver(r)

	work := func() {
		gobot.Every(1*time.Second, func() {
			v0, _ := bpm180.Read(0, 1, 860)
			v1, _ := bpm180.Read(1, 1, 860)

			a0, _ := bpm180.AnalogRead("0")
			a1, _ := bpm180.AnalogRead("1")

			println("pin0:", v0, " ", a0)
			println("pin1:", v1, " ", a1)
		})
	}

	robot := gobot.NewRobot("adcBot",
		[]gobot.Connection{r},
		[]gobot.Device{bpm180},
		work,
	)

	robot.Start()
}
