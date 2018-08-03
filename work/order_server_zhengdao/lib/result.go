package lib

type result struct {
}

//仿静态方法
var Result = &result{}

func (r *result) Success(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code": 0,
		"msg":  msg,
		"data": data,
	}
}

func (r *result) Fail(code int, msg string) map[string]interface{} {
	return r.FailWithData(code, msg, nil)
}

func (r *result) FailWithData(code int, msg string, data interface{}) map[string]interface{} {
	if code == 0 {
		code = -1
	}
	return map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}
