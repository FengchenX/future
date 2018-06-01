
package main

import (
	"fmt"
	"errors"

)

func main() {
	err := f()
	fmt.Println(err)
}

var errFoo = errors.New("this is a new error")

func f() error {
	return errFoo
}