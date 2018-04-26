package main

import (
	"fmt"
	"reflect"
)

func main() {
	//TypeOf()
	//TypeValue()
	//Set()
	refOpStruct()
	//refChangeStruct()
	//sliceAppend()
}


func TypeOf() {
	var a = 10
	var i interface{}
	i = a
	t := reflect.TypeOf(i)
	switch t.Kind().String() {  //Kind是基础类型 结构体会输出struct, name打印比较具体
	case "int":
		fmt.Println("int")
	default:
		fmt.Println("other")
	}

	var _p = p{"uio",20}
	i = _p

	t1 := reflect.TypeOf(i)
	fmt.Println(t1.Name())   //p
	fmt.Println(t1.String()) //main.p
}

type p struct {
	name string
	age int
}

func(this p) say() {
	fmt.Println("Hello")
}


func TypeValue() {
	var a = 10
	var i interface{}
	i=a
	v := reflect.ValueOf(i)
	fmt.Println(v.Int())
}

func Set() {
	var a2 float64
    fv2 := reflect.ValueOf(&a2)
    fv2.Elem().SetFloat(520.00)
    fmt.Printf("%v\n", a2)    //520
}


type NotknownType struct {
	S1 string
	S2 string
	S3 string
}

func (n NotknownType) String() string {
	return n.S1 + " & " + n.S2 + " & " + n.S3
}

func(n NotknownType) Say() {
	fmt.Println("uio")
}

var secret interface{} = NotknownType{"Go", "C", "Python"}

func refOpStruct() {
	value := reflect.ValueOf(secret)
    fmt.Println(value) //Go & C & Python
    typ := reflect.TypeOf(secret)
    fmt.Println(typ) //main.NotknownType

    knd := value.Kind()
    fmt.Println(knd) // struct

    for i := 0; i < value.NumField(); i++ {
        fmt.Printf("Field %d: %v\n", i, value.Field(i))
    }

    results := value.Method(0).Call(nil)
	fmt.Println(results) // [Go & C & Python]
	fmt.Println(value.NumMethod())
}

type T struct {
    A int
    B string
}

func refChangeStruct() {
	t := T{18, "nick"}
    s := reflect.ValueOf(&t).Elem()
    typeOfT := s.Type()

    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        fmt.Printf("%d: %s %s = %v\n", i,
            typeOfT.Field(i).Name, f.Type(), f.Interface())
    }

    s.Field(0).SetInt(25)
	s.Field(1).SetString("nicky")
	s.FieldByName("B").SetString("wudi")
    fmt.Println(t)
}

type test struct {
    S1 string
    s2 string
    s3 string
}

var s interface{} = &test{
    S1: "s1",
    s2: "s2",
    s3: "s3",
}

func f1() {
	val := reflect.ValueOf(s)
    fmt.Println(val)                      //&{s1 s2 s3}
    fmt.Println(val.Elem())               //{s1 s2 s3}
    fmt.Println(val.Elem().Field(0))      //s1
	val.Elem().Field(0).SetString("hehe") //S1大写
	fmt.Println(s)
}

func f2() {
	var i interface{} = []test{
		{"uio", "kk", "oo"},
		{"dd", "ee", "qq"},
		{"ww", "cc", "zz"},
	}

	val := reflect.ValueOf(i)
	val.Index(0).Field(0).SetString("hh")
	fmt.Println(i)
}

//反射包终极操作，很多知识点在这,关于slice操作
func sliceAppend() {
	var ps []struct{
		Name string  //反射时字段需要大写
		Age int
		Grade int
	}

	SearchAll:= func(x interface{}) {
		pp:=[][]interface{}{}    //[]struct数据，，我们要想办法给x
		for i:=0;i<10;i++ {
			var dest []interface{}	
			var name = "nio"
			var age = 23
			var grade = 5
			dest = append(dest,name,age,grade)
			pp = append(pp,dest)
		}
		val := reflect.ValueOf(x) //获取value指针
		typ := reflect.TypeOf(x)  //得到指针类型
		strtyp :=typ.Elem().Elem()                //想要得到struct类型 指针类型取值是.Elem()

		for _,row := range pp {
			oneptr := reflect.New(strtyp)
			for j,elem := range row {
				oneptr.Elem().Field(j).Set(reflect.ValueOf(elem))
			}
			val2 := reflect.Append(val.Elem(),oneptr.Elem())   //
			val.Elem().Set(val2)
		}
	}

	SearchAll(&ps)     //这地方有个slice大坑，slice是按引用进行传递的，，但是传递完之后对之前已存在的内存进行修改是有效的，如果一旦调用了append，新append部分并不会对传递前slice产生效果,需要传递指针
	for _,i := range ps {
		fmt.Println(i.Name,i.Age,i.Grade)
	}
}
