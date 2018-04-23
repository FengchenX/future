
package main

import (
	"fmt"
	"regexp"

)

/**
. 任意字符
[xyz] 字符族
[a-z] a到z之间的所有字符
[a-zA-Z0-9]*  a-z,A-Z,0-9之间的所有字符组成的字符串
xy 匹配x后接着匹配y
x|y 匹配x或y(优先匹配x)
x*             重复>=0次匹配x，越多越好（优先重复匹配x）
x+             重复>=1次匹配x，越多越好（优先重复匹配x）
x?             0或1次匹配x，优先1次
x{n,m}         n到m次匹配x，越多越好（优先重复匹配x）
x{n,}          重复>=n次匹配x，越多越好（优先重复匹配x）
x{n}           重复n次匹配x
x*?            重复>=0次匹配x，越少越好（优先跳出重复）
x+?            重复>=1次匹配x，越少越好（优先跳出重复）
x??            0或1次匹配x，优先0次
x{n,m}?        n到m次匹配x，越少越好（优先跳出重复）
x{n,}?         重复>=n次匹配x，越少越好（优先跳出重复）
x{n}?          重复n次匹配x
^              匹配文本开始，标志m为真时，还匹配行首
$              匹配文本结尾，标志m为真时，还匹配行尾
\A             匹配文本开始
\b             单词边界（一边字符属于\w，另一边为文首、文尾、行首、行尾或属于\W）
\B             非单词边界
\z             匹配文本结尾
*/


func main() {
/*
	var validID = regexp.MustCompile(`^[a-z]+\[[0-9]+\]$`)
	fmt.Println(validID.MatchString("adam[23]"))
	fmt.Println(validID.MatchString("eve[7]"))
	fmt.Println(validID.MatchString("Job[48]"))
	fmt.Println(validID.MatchString("snakey"))
*/
	//regexpMatchString()
	//regexpFindAllString()
	//regexpFindAllStringSubmatch()
	regexpTest()

}

func regexpMatchString() {
	matched, err := regexp.MatchString("foo.*","seafood")
	fmt.Println(matched,err)
	matched, err = regexp.MatchString("bar.*","seafood")
	fmt.Println(matched, err)
	matched, err = regexp.MatchString("a(b","seafood")
	fmt.Println(matched,err)
}
func regexpFindAllString() {
	re := regexp.MustCompile("a.")
	fmt.Println(re.FindAllString("paranormal", -1))
	fmt.Println(re.FindAllString("paranormal",2))
	fmt.Println(re.FindAllString("graal", -1))
	fmt.Println(re.FindAllString("none", -1))
}
func regexpFindAllStringSubmatch() {
	re := regexp.MustCompile("a(x*)b")
	fmt.Printf("%q\n",re.FindAllStringSubmatch("-ab-",-1))
	fmt.Printf("%q\n",re.FindAllStringSubmatch("-axxb-",-1))
	fmt.Printf("%q\n",re.FindAllStringSubmatch("-ab-axb-",-1))
	fmt.Printf("%q\n",re.FindAllStringSubmatch("-axxb-ab-",-1))
}
func regexpFindString() {
	re := regexp.MustCompile("fo.?")
	fmt.Printf("%q\n",re.FindString("seafood"))
	fmt.Printf("%q\n",re.FindString("meat"))
}


func regexpTest() {
/*
下面是抓取url正则表达式，第二个测试很好用
([/w-]+/.)+[/w-]+.([^a-z])(/[/w-: ./?%&=]*)?|[a-zA-Z/-/.][/w-]+.([^a-z])(/[/w-: ./?%&=]*)?
(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]
*/

	var re = regexp.MustCompile(`(https?|ftp|file)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]`)
	var str = "http://www.runoob.com:80/html/html-tutorial.html****http://www.runoob.com:80/html/html-tutorial.html";
	fmt.Println(re.FindAllStringSubmatch(str, -1))

	var re2 = regexp.MustCompile("[a-zA-Z0-9]*&[a-zA-Z0-9]*/[a-z]*")
	var str2 = "abcdABCD&99sda/html"
	fmt.Println(re2.FindAllString(str2, -1))

}