//author xinbing
//time 2018/9/11 10:23
//
package db

import (
	"fmt"
	"testing"
)

type a struct {
	name *string
}

func (p *a) tt() {
	nn := "张无忌"
	na := &a{
		name: &nn,
	}
	p.name = na.name
	na.name = nil
}
func TestInitGormDB(t *testing.T) {
	aa := a{
	}
	fmt.Println(aa)
	aa.tt()
	fmt.Println(aa)
}


