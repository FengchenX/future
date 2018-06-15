package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//utf8DecodeRune()
	//utf8DecodeLastRune()
	//utf8DecodeLastRuneInString()
	//utf8FullRune()
	//utf8FullRuneInString()
	utf8RuneCount()
}

func utf8DecodeRune() {
	b := []byte("Hello, 世界")
	for len(b) > 0 {
		r, size := utf8.DecodeRune(b)
		fmt.Printf("%c,%v\n", r, size)
		b = b[size:]
	}
}

func utf8DecodeLastRune() {
	p := []byte("Hello, 世界")
	for len(p) > 0 {
		r, size := utf8.DecodeLastRune(p)
		fmt.Printf("%c,%v\n", r, size)
		p = p[:len(p)-size]
	}
}

func utf8DecodeLastRuneInString() {
	str := "Hello, 世界"
	for len(str) > 0 {
		r, size := utf8.DecodeLastRuneInString(str)
		fmt.Printf("%c,%v\n", r, size)
		str = str[:len(str)-size]
	}
}

func utf8DecodeRuneInString() {
	str := "Hello, 世界"
	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		fmt.Printf("%c, %v", r, size)
		str = str[size:]
	}
}

func utf8FullRune() {
	var buf = []byte{228, 184, 150}
	fmt.Println(utf8.FullRune(buf))
	fmt.Println(utf8.FullRune(buf[:2]))
}

func utf8FullRuneInString() {
	str := "世"
	fmt.Println(utf8.FullRuneInString(str))
	fmt.Println(utf8.FullRuneInString(str[:2]))
}

func utf8RuneCount() {
	buf := []byte("Hello, 世界")
	fmt.Println(len(buf))
	fmt.Println(utf8.RuneCount(buf))
}

func utf8RuneCountInString() {
	str := "Hello, 世界"
	fmt.Println(len(str))
	fmt.Println(utf8.RuneCountInString(str))
}

func utf8RuneLen() {
	fmt.Println(utf8.RuneLen('a'))
	fmt.Println(utf8.RuneLen('世'))
}

func utf8RuneStart() {
	buf := []byte("a界")
	fmt.Println(utf8.RuneStart(buf[0]))
	fmt.Println(utf8.RuneStart(buf[1]))
	fmt.Println(utf8.RuneStart(buf[2]))
}

func utf8Vaild() {
	valid := []byte("Hello, 世界")
	invalid := []byte{0xff, 0xfe, 0xfd}
	fmt.Println(utf8.Valid(valid))
	fmt.Println(utf8.Valid(invalid))
}

func utf8VaildRune() {
	valid := 'a'
	invalid := rune(0xfffffff)
	fmt.Println(utf8.ValidRune(valid))
	fmt.Println(utf8.ValidRune(invalid))
}

func utf8VaildString() {
	valid := "Hello, 世界"
	invalid := string([]byte{0xff, 0xfe, 0xfd})
	fmt.Println(utf8.ValidString(valid))
	fmt.Println(utf8.ValidString(invalid))
}
