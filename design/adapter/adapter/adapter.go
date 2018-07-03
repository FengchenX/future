package adapter

import "fmt"



type INonBattery interface {
	Use()
}

type IReBattery interface {
	Use()
	Charge()
}

type NonA struct {

}

func (NonA) Use() {
	fmt.Println("NonA using")
}

//适配可充电电池使用接口
type AdapterNonToYes struct {
	INonBattery
}

func (AdapterNonToYes) Charge() {
	fmt.Println("AdapterNonToYes Charging")
}

//接口的适配器模式

type ReBatteryAbstract struct {

}

func (ReBatteryAbstract) Use() {
	fmt.Println("ReBatteryAbstract using")
}

func (ReBatteryAbstract) Charge() {
	fmt.Println("ReBatteryAbstract Charging")
}

type NonReB struct {
	ReBatteryAbstract
}

func (NonReB) Use() {
	fmt.Println("NonReB using")
}

//test

func AdapterTest() {
	var battery IReBattery

	battery = AdapterNonToYes{NonA{}}
	battery.Use()
	battery.Charge()

	battery = NonReB{}
	battery.Use()
	battery.Charge()
}

