package test

import (
	"testing"
)

func Test_div(t *testing.T) {
	i := Div(6, 2)
	if i != 3 {
		t.Error("未通过")
	} else {
		t.Log("通过")
	}
}

func Test_add(t *testing.T) {
	i := Add(2, 3)
	if i != 5 {
		t.Error("未通过")
	} else {
		t.Log("通过")
	}
}
