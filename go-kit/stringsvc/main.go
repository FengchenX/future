package main

import (
	"strings"
	"errors"
	"context"
)

func main() {

}

type StringService interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}

type stringService struct {}
func (stringService) Uppercase(_ context.Context, s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

var ErrEmpty = errors.New("Empty string")
func (stringService) Count(_ context.Context, s string) int {
	return len(s)
} 