package main

import "fmt"

type Device interface {
	TurnOn()
	TurnOff()
}

type TV struct{}

func (tv *TV) TurnOn() {
	fmt.Println("Turning on TV")
}

func (tv *TV) TurnOff() {
	fmt.Println("Turning off TV")
}

type Radio struct{}

func (radio *Radio) TurnOn() {
	fmt.Println("Turning on radio")
}

func (radio *Radio) TurnOff() {
	fmt.Println("Turning off radio")
}

type RemoteControl struct {
	device Device
}

func NewRemoteControl(device Device) *RemoteControl {
	return &RemoteControl{device: device}
}

func (r *RemoteControl) TurnOn() {
	r.device.TurnOn()
}

func (r *RemoteControl) TurnOff() {
	r.device.TurnOff()
}

type AdvanceRemoteControl struct {
	RemoteControl
}

func NewAdvanceRemoteControl(device Device) AdvanceRemoteControl {
	return AdvanceRemoteControl{RemoteControl{device: device}}
}

func (r *AdvanceRemoteControl) Mute() {
	fmt.Println("Muting the device")
}

func main() {
	tv := &TV{}
	radio := &Radio{}

	tvRemoteControl := NewRemoteControl(tv)
	radioRemoteControl := NewRemoteControl(radio)

	tvRemoteControl.TurnOn()
	tvRemoteControl.TurnOff()

	radioRemoteControl.TurnOn()
	radioRemoteControl.TurnOff()

}
