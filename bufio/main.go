

package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"

)

func main() {
	//bufioSplit()
	bufioWriter()
}


func bufioSplit() {
	const input = "feng chen ni hao a"
	scanner:=bufio.NewScanner(strings.NewReader(input))
	split:=func (data []byte, atEOF bool) (addvace int, token []byte, err error) {
		addvace,token,err = bufio.ScanWords(data,atEOF)
		return
	}
	scanner.Split(split)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}


func bufioWriter() {
	bw:=bufio.NewWriter(os.Stdout)
	fmt.Fprintln(bw,"hello","feng")
	fmt.Fprintln(bw, "zd")
	bw.Flush()
}