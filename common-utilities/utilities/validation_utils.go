//author xinbing
//time 2018/9/5 13:52
//正则表达式验证集合
package utilities

import (
	"regexp"
)

// 校验国内手机号码
var phoneReg = regexp.MustCompile("^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0-3,5-8])|(18[0-9])|166|198|199)\\d{8}$")
func ValidPhone(phone string) bool{
	if len(phone) == 0 {
		return false
	}
	return phoneReg.MatchString(phone)
}

// 校验邮箱
var emailReg = regexp.MustCompile("^[A-Za-z0-9]+([-_.][A-Za-z0-9]+)*@([A-Za-z0-9]+[-.])+[A-Za-z0-9]{2,4}$")
func ValidEmail(email string) bool {
	if len(email) == 0 {
		return false
	}
	return emailReg.MatchString(email)
}