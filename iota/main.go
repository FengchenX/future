
package main
import (
	"fmt"

)
func main() {
	//移位操作
	fmt.Println(flag_8)
	fmt.Println(one_2)  //1
	fmt.Println(one_4)  //10
	fmt.Println(one_5)  //10
	fmt.Println(one_6)  //5
	fmt.Println(one_7)  //6
}
//完整示例
type FileMode uint32
const (
	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
	ModeAppend                                     // a: append-only
	ModeExclusive                                  // l: exclusive use
	ModeTemporary                                  // T: temporary file; Plan 9 only
	ModeSymlink                                    // L: symbolic link
	ModeDevice                                     // D: device file
	ModeNamedPipe                                  // p: named pipe (FIFO)
	ModeSocket                                     // S: Unix domain socket
	ModeSetuid                                     // u: setuid
	ModeSetgid                                     // g: setgid
	ModeCharDevice                                 // c: Unix character device, when ModeDevice is set
	ModeSticky                                     // t: sticky

	// Mask for the type bits. For regular files, none will be set.
	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice

	ModePerm FileMode = 0777 // Unix permission bits
)
type myint int8
const (
	_ = iota
	flag_2 myint = 1 << iota   // 00000010
	flag_4               // 00000100
	flag_8               // 00001000
)
const (
	one_1 int = iota  //0
	one_2             //1
	one_3             //2
	one_4 =10         //10
	one_5             //10
	one_6 int = iota  //5
	one_7             //6
)