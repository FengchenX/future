//author xinbing
//time 2018/9/11 11:40
//
package utils

type Resp struct {
	Code		int
	Msg			string
	Data 		interface{}
}


func (p Resp) Success(msg string, data interface{}) *Resp {
	p.Code = 0
	p.Msg = msg
	p.Data = data
	return &p
}

func (p Resp) Failed(msg string) *Resp {
	p.Code = -1
	p.Msg = msg
	return &p
}