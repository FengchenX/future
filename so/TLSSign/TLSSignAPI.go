package TLSSign
/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L $GOPATH/src/TLSSign -lsigcheck -ldl
#include <stdlib.h>
#include "sigcheck.h"
#include "multi_thread.h"
*/
import "C"

import (
	"errors"
	"unsafe"
)

// 签名设置结构体
type TLSSignConf struct {
        AccType int
        Identifier string

        AppId3rd string
        SDKAppId int

        Version int
        Expire int

		// PEM格式的密钥和公钥，如果要签名或者验证签名，就要设置其中一个
        PriKey string
		PubKey string
}

// 签名方法，返回签名的 Base64 字符串或者错误
func (conf *TLSSignConf) Sign() (string, error) {
	var sigBuf, errMsgBuf [1024]C.char

	cstr_identifier := C.CString(conf.Identifier)
	cstr_PriKey := C.CString(conf.PriKey)
	defer C.free(unsafe.Pointer(cstr_identifier))
	defer C.free(unsafe.Pointer(cstr_PriKey))

	fail := C.tls_gen_sig_ex(
		C.uint(conf.SDKAppId),
		cstr_identifier,
		(*C.char)(unsafe.Pointer(&sigBuf)),
		C.uint(len(sigBuf)),
		cstr_PriKey,
		C.uint(len(conf.PriKey)),
		(*C.char)(unsafe.Pointer(&errMsgBuf)),
		C.uint(len(errMsgBuf)))

	if fail != 0 {
		result := C.GoString(&errMsgBuf[0])
		return "", errors.New(result)
	} else {
		result := C.GoString(&sigBuf[0])
		return result, nil
	}
}

// 验证签名，返回签名的过期时间，签名时间以及是否通过验证，如果出错，那么 err != nil
func (conf *TLSSignConf) CheckSign(sign string) (Expire, signed int, valid bool, err error) {
	var ExpiredAt, signedAt C.uint
	var errMsgBuf [1024]C.char

	cstr_sign := C.CString(sign)
	cstr_PubKey := C.CString(conf.PubKey)
	cstr_identifier := C.CString(conf.Identifier)
	defer C.free(unsafe.Pointer(cstr_sign))
	defer C.free(unsafe.Pointer(cstr_PubKey))
	defer C.free(unsafe.Pointer(cstr_identifier))

	fail := C.tls_vri_sig_ex(
		cstr_sign,
		cstr_PubKey,
		C.uint(len(conf.PubKey)),
		C.uint(conf.SDKAppId),
		cstr_identifier,
		(*C.uint)(unsafe.Pointer(&ExpiredAt)),
		(*C.uint)(unsafe.Pointer(&signedAt)),
		(*C.char)(unsafe.Pointer(&errMsgBuf)),
		C.uint(len(errMsgBuf)))

	if fail != 0 {
		err := C.GoString(&errMsgBuf[0])
		return 0, 0, false, errors.New(err)
	} else {
		return int(ExpiredAt), int(signedAt), true, nil
	}
}

// 多线程情况下，使用API之前，需要先调用这个函数
func MultiThreadSetup() int {
	return int(C.multi_thread_setup())
}

// 多线程情况下，若已经不需要使用本API，就调用这个函数。
func MultiThreadCleanup() {
	C.multi_thread_cleanup()
}
