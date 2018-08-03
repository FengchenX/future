package lib

import (
	"sync"
)

const MAXCOUNT = 50
const MAXCHANCOUNT = 1024

func Adder(i int, c chan int) {
	c <- i
}

type wp struct {
	cacheCh chan int
	stopCh  chan struct{}
	count   int
	lock    sync.RWMutex
	wpFunc  func(i int, c chan int)
	minChan chan int
}

func NewWp() *wp {
	w := &wp{
		cacheCh: make(chan int, MAXCHANCOUNT),
		stopCh:  make(chan struct{}),
		minChan: make(chan int, MAXCHANCOUNT),
		count:   0,
	}
	w.wpFunc = Adder
	return w
}

func (this *wp) run() {
	for {
		if this.count < MAXCOUNT-1 {
			this.lock.Lock()
			this.count++
			this.lock.Unlock()
			select {
			case num := <-this.cacheCh:
				go this.wpFunc(num, this.minChan)
			case <-this.stopCh:
				return
			}
		}
	}
}

func (this *wp) Min() {
	for {
		select {
		case <-this.minChan:
			this.lock.Lock()
			this.count--
			this.lock.Unlock()
		}
	}
}

func (this *wp) send(num int) {
	var le int
	this.lock.Lock()
	defer this.lock.Unlock()
	le = len(this.cacheCh)
	if le < MAXCHANCOUNT-1 {
		this.cacheCh <- num
	}
}
