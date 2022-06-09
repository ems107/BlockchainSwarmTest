package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ev3go/ev3dev"
)

func main() {
	direction := os.Args[1]

	outA, err := ev3dev.TachoMotorFor("ev3-ports:outA", "lego-ev3-l-motor")
	if err != nil {
		panic(err)
	}
	maxMedium := outA.MaxSpeed()
	outB, err := ev3dev.TachoMotorFor("ev3-ports:outB", "lego-ev3-l-motor")
	if err != nil {
		panic(err)
	}

	if direction == "Right" {
		outA.SetSpeedSetpoint(30 * maxMedium / 100).Command("run-forever")
		checkErrors(outB)
		time.Sleep(time.Second / 2)
		outA.Command("stop")
		checkErrors(outA, outB)
	} else if direction == "Left" {
		outB.SetSpeedSetpoint(30 * maxMedium / 100).Command("run-forever")
		checkErrors(outB)
		time.Sleep(time.Second / 2)
		outB.Command("stop")
		checkErrors(outA, outB)
	}
}

func checkErrors(devs ...ev3dev.Device) {
	for _, d := range devs {
		err := d.(*ev3dev.TachoMotor).Err()
		if err != nil {
			drv, dErr := ev3dev.DriverFor(d)
			if dErr != nil {
				drv = fmt.Sprintf("(missing driver name: %v)", dErr)
			}
			addr, aErr := ev3dev.AddressOf(d)
			if aErr != nil {
				drv = fmt.Sprintf("(missing port address: %v)", aErr)
			}
			log.Fatalf("motor error for %s:%s on port %s: %v", d, drv, addr, err)
		}
	}
}
