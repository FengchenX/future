package main

import (
	"time"
	"log"
	"os"
	"fmt"
	"flag"
	"runtime/pprof"
)

var (
	cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")
)

func main() {
	fmt.Println("begin")
	flag.Parse()
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		
		go func() {
			pprof.StartCPUProfile(f)
			time.Sleep(2*time.Second)
			defer pprof.StopCPUProfile()
		}()
		
	}

	for i := 0; i < 30; i++ {

        nums := fibonacci(i)

		fmt.Println(nums)
		for {

		}

    }
}

//递归实现的斐波纳契数列

func fibonacci(num int) int {

    if num < 2 {

        return 1

    }

    return fibonacci(num-1) + fibonacci(num-2)

}
