package adapter

import (
	"fmt"
)

type IBike interface {
	Drive()
}

type MountainBike struct {
}

func (MountainBike) Drive() {
	fmt.Println("Moun driving")
}

type NorBike struct {
}

func (NorBike) Drive() {
	fmt.Println("Nor driving")
}

type IGoxiangBike interface {
	UnLock()
	Drive()
}

type Mobai struct {
}

func (Mobai) UnLock() {
	fmt.Println("mobai unlock")
}
func (Mobai) Drive() {
	fmt.Println("mobai driving")
}

type AdapterBikeToGoxiang struct {
	IBike
}

func (AdapterBikeToGoxiang) UnLock() {
	fmt.Println("adapter unlock")
}
