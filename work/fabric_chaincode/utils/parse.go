package utils

import (
	"strconv"
	"strings"
)

func ParseToUintArray(in string) (out []uint, err error) {
	for _, v := range strings.Split(in, "@") {
		it, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		out = append(out, uint(it))
	}
	return
}

func ParseToStringArray(in string) (out []string, err error) {
	for _, v := range strings.Split(in, "@") {
		out = append(out, v)
	}
	return
}
