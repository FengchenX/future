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




type IPlayer interface {
	Player(name string)
	Attack()
	Defense()
}

type Fowards struct {
	name string
}

func (f *Fowards) Player(name string) {
	f.name = name
}

func (f *Fowards) Attack() {
	fmt.Printf("前锋%s进攻\n", f.name)
}

func (f *Fowards) Defense() {
	fmt.Printf("前锋%s防守\n", f.name)
}

type Center struct {
	name string
}

func (f *Center) Player(name string) {
	f.name = name
}

func (f *Center) Attack() {
	fmt.Printf("中锋%s进攻\n", f.name)
}

func (f *Center) Defense() {
	fmt.Printf("中锋%s防守\n", f.name)
}

type Guards struct {
	name string
}

func (f *Guards) Player(name string) {
	f.name = name
}

func (f *Guards) Attack() {
	fmt.Printf("后卫%s进攻\n", f.name)
}

func (f *Guards) Defense() {
	fmt.Printf("后卫%s防守\n", f.name)
}

type ForCenter struct {
	name string
}

func (fc *ForCenter) Player(name string) {
	fc.name = name
}

func (fc *ForCenter) 进攻() {
	fmt.Printf("外籍中锋%s进攻\n", fc.name)
}

func (fc *ForCenter) 防守() {
	fmt.Printf("外籍中锋%s防守\n", fc.name)
}

type Translator struct{
	ForCenter
}

func (t *Translator) Player(name string) {
	t.ForCenter.Player(name)
}

func (t *Translator) Attack() {
	t.进攻()
}

func (t *Translator) Defense() {
	t.防守()
}
