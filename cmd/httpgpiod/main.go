package main

import (
	"flag"
	
	"github.com/penguinpowernz/http-gpio-server/rest"
	"github.com/penguinpowernz/http-gpio-server/rpi"
	rpio "github.com/stianeikeland/go-rpio"
)

func main() {
	var fake bool
	var port string

	flag.BoolVar(&fake, "fake", false, "fake outputs")
	flag.StringVar(&port, "p", "80", "port to run on")
	flag.Parse()

	rpio.Open()
	defer rpio.Close()

	var outputs rpi.Outputs

	if fake {
		outputs = rpi.NewFakeOutputs()
	} else {
		outputs = rpi.NewOutputs()
	}

	outputs.AllOff()

	svr := rest.NewAPI(outputs)
	svr.Run(":" + port)
}
