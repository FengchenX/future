package utils

type ResultObject struct {
	Code int
	Msg string
	Data interface{}
}
//仿静态方法
var Result = &ResultObject{}
func (r *ResultObject) Success(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"Code": 0,
		"Msg":  msg,
		"Data": data,
	}
}

func (r *ResultObject) Fail(msg string) map[string]interface{} {
	return r.FailWithData(-1, msg, nil)
}

func (r *ResultObject) FailWithData(code int, msg string, data interface{}) map[string]interface{} {
	if code == 0 {
		code = -1
	}
	return map[string]interface{}{
		"Code": code,
		"Msg":  msg,
		"Data": data,
	}
}

func (r *ResultObject) New(code int, msg string, data interface{}) ResultObject{
	return ResultObject{
		Code:code,
		Msg:msg,
		Data:data,
	}
}
