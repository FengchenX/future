//author xinbing
//time 2018/9/5 13:54
package utilities

import (
	"fmt"
	"testing"
)

func TestValidPhone(t *testing.T) {
	fmt.Println(ValidPhone("17417771777"))
}

func TestValidEmail(t *testing.T) {
	fmt.Println(ValidEmail("bin-g.xin@cnlaunch.com-a.cn.bb"))
}

func TestPwdGrade(t *testing.T) {
}