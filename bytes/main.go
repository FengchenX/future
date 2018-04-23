

package main


import (
	"sort"
	"os"
	"fmt"
	"bytes"

)

func main() {
	//bytesBuffer()
	//bytesReader()
	//bytesCompare()
	//bytesCompareSearch()
	bytesTrimPrefix()
}

func bytesBuffer(){
	var b bytes.Buffer
	b.Write([]byte("Hello"))
	fmt.Fprintf(&b,"world!")
	b.WriteTo(os.Stdout)
}

func bytesReader() {
	b:=bytes.NewBufferString("fengchen")
	b.WriteTo(os.Stdout)
}

func bytesCompare() {
	var a = []byte{10,20,30}
	var b = []byte{20,20,30}
	if bytes.Compare(a,b)<0 {
		fmt.Println("a < b")
	}
}

func bytesCompareSearch() {
	var needle  = []byte{10,20}
	var haystack = [][]byte{
		{9,20},
		{10,20},
	} // assume sorted
	i:=sort.Search(len(haystack),func(i int) bool {
		return bytes.Compare(haystack[i],needle)>=0
	})

	if i <len(haystack) && bytes.Equal(haystack[i],needle) {
		fmt.Println("find",i)
	}
}

func bytesTrimPrefix() {
	var b = []byte("Goodbye,World")
	b = bytes.TrimPrefix(b,[]byte("Goodbye"))
	fmt.Println("Hello",string(b))
}