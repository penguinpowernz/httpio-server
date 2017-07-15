package remoteput

import (
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

// PinMap maps the output index to the BCM Pin
var PinMap = map[int]int{
	1: 17, // WiringPi 0
	2: 18, // WiringPi 1
	3: 27, // WiringPi 2
	4: 22, // WiringPi 3
	5: 23, // WiringPi 4
	6: 24, // WiringPi 5
	7: 25, // WiringPi 6
	8: 4,  // WiringPi 7
}

type Outputs map[int]*Output

func (ot Outputs) AllOff() {
	for _, o := range ot {
		o.SetPosition(0)
	}
}

func (ot Outputs) AllOn() {
	for _, o := range ot {
		o.SetPosition(1)
		time.Sleep(100 * time.Millisecond)
	}
}

type pin interface {
	High()
	Low()
	Output()
	Read() rpio.State
}

type Output struct {
	Index    int `json:"index"`
	Position int `json:"position"`
	pin      pin
}

func NewOutputs(pinMaps ...map[int]int) Outputs {
	pinMap := PinMap
	if len(pinMaps) > 0 {
		pinMap = pinMaps[0]
	}

	o := Outputs{}

	for idx := range pinMap {
		o[idx] = NewOutput(idx, pinMap[idx])
	}

	o.AllOff()
	return o
}

func NewFakeOutputs(pinMaps ...map[int]int) Outputs {
	pinMap := PinMap
	if len(pinMaps) > 0 {
		pinMap = pinMaps[0]
	}

	o := Outputs{}

	for idx, pin := range pinMap {
		o[idx] = NewFakeOutput(idx, pin)
	}

	o.AllOff()
	return o
}

type fakePin struct {
	current rpio.State
}

func (f fakePin) High()   { f.current = rpio.High }
func (f fakePin) Low()    { f.current = rpio.Low }
func (f fakePin) Output() {}
func (f fakePin) Read() rpio.State { return f.current }

func NewFakeOutput(idx, pin int) *Output {
	p := fakePin{}
	p.Output()
	o := &Output{idx, 0, p}
	o.Position = o.GetPosition()
	return o
}


func NewOutput(idx, pin int) *Output {
	p := rpio.Pin(pin)
	p.Output()
	o := &Output{idx, 0, p}
	o.Position = o.GetPosition()
	return o
}

func (o *Output) SetPosition(pos int) {
	switch pos {
	case 1:
		o.pin.Low()
	case 0:
		o.pin.High()
	}

	o.Position = o.GetPosition()
}

func (o *Output) GetPosition() int {
	switch o.pin.Read() {
	case 1:
		return 0
	case 0:
		return 1
	}

	return -1
}
