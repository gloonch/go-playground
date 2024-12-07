package main

import "fmt"

// Target
type mobile interface {
	chargeAppleMobile()
}

// Prototype (Concrete implementation)
type apple struct {
}

func (a *apple) chargeAppleMobile() {
	fmt.Println("Charging apple mobile")
}

// Adaptee
type android struct{}

func (a *android) chargeAndroidMobile() {
	fmt.Println("Charging android mobile")
}

// Adapter
type androidadapter struct {
	android *android
}

func (receiver *androidadapter) chargeAppleMobile() {
	receiver.android.chargeAndroidMobile()
}

// Client
type client struct {
}

func (c *client) chargeMobile(mob mobile) {
	mob.chargeAppleMobile()
}

func main() {
	// First/Initial requirement
	apple := &apple{}
	client := &client{}
	client.chargeMobile(apple)

	// Extended requirement i.e Charge android mobile
	android := &android{}
	androidadapter := &androidadapter{
		android: android,
	}
	client.chargeMobile(androidadapter)
}
