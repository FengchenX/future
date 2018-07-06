package bridge

import (
	"fmt"
)

type ISchedule interface {
	SetEmp(emps ...IEmployee)
	Do()
}
type SchA struct {
	A, B IEmployee
}

func (sa *SchA) SetEmp(emps ...IEmployee) {
	sa.A = emps[0]
	sa.B = emps[1]
}

func (sa SchA) Do() {
	sa.A.Income()
	sa.B.Income()
}

type IEmployee interface {
	Name() string
	Income()
}

type Chushi struct {
}

func (Chushi) Name() string {
	return "厨师"
}
func (Chushi) Income() {
	fmt.Println("厨师收入了")
}

type Fuwuyuan struct {
}

func (Fuwuyuan) Name() string {
	return "服务员"
}

func (Fuwuyuan) Income() {
	fmt.Println("服务员收入了")
}
