package test

import (
	"testing"
	"fmt"
	"strings"
)

func TestJoin(t *testing.T) {
	slic := []string {
		"1","2","3","4","5",
	}
	fmt.Println(strings.Join(slic,","))
	fmt.Println("('" + strings.Join(slic,"','") + "')")
	slic2 := make([]string,0)
	slic = append(slic,slic2...)
	fmt.Println(len(slic))
	fmt.Println("==================")
}
