package controllers

import (
	"github.com/gin-gonic/gin"
	"os"
	"runtime/pprof"
	"log"
	"time"
)

type RespondTemp struct {
	Errorcode    int
	Errormessage string
	Ok           bool
	Data         interface{}
}

func Profile(controller *gin.Context) {
	respondTemp := RespondTemp{}
	f, err := os.Create("mem.mprof")
	if err != nil {
		log.Println("create memprofile err:", err.Error())
		respondTemp.Errormessage = err.Error()
		respondTemp.Ok = false
		controller.JSON(200, respondTemp)
		return
	}
	pprof.WriteHeapProfile(f)
	f.Close()

	fc, cerr := os.OpenFile("cpu.prof", os.O_RDWR|os.O_CREATE, 0644)
	if cerr != nil {
		log.Println("create memprofile err:", err.Error())
		respondTemp.Errormessage = err.Error()
		respondTemp.Ok = false
		controller.JSON(200, respondTemp)
		return
	}
	pprof.StartCPUProfile(fc)
	time.Sleep(2 * time.Second)
	pprof.StopCPUProfile()
	fc.Close()

	respondTemp.Ok = true
	controller.JSON(200, respondTemp)
	return
}
