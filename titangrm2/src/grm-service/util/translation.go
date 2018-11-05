package util

import (
	"github.com/leonelquinteros/gotext"
)

func TR(str string, vars ...interface{}) string {
	return gotext.Get(str, vars...)
}

func LoadTranslation(lib, lang, dom string) {
	gotext.Configure(lib, lang, dom)
}
