package main

import (
	"bitbucket.org/hamstar/remoteput"
	rpio "github.com/stianeikeland/go-rpio"
	"flag"
)

func main() {
	var fake bool
	var port string

	flag.BoolVar(&fake, "fake", false, "fake outputs")
	flag.StringVar(&port, "p", "80", "port to run on")
	flag.Parse()

	rpio.Open()
	defer rpio.Close()

	var outputs remoteput.Outputs

	if fake {
		outputs = remoteput.NewFakeOutputs()
	} else {
		outputs = remoteput.NewOutputs()
	}

	outputs.AllOff()

	svr := remoteput.NewAPI(outputs)
	svr.Run(":" + port)
}
