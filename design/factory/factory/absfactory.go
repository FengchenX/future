package factory

import (
	"fmt"
)

//抽象工厂
type AbstractFactory interface {
	Produce() pen  //生产笔
}

type PencilFactory struct {

}

func (PencilFactory) Produce() pen {
	return new(pencil)
}

type BrushPen struct {

}

func (BrushPen) Produce() pen {
	return new(brushPen)
}

type User struct {
	Name string
	Age int
}

type IUser interface {
	Insert(user User)
	Get(name string) User
}

type Mysql struct {

}

func (m *Mysql) Insert(user User) {
	fmt.Println("Mysql insert", user)
}

func (m *Mysql) Get(name string) User {
	fmt.Println("Mysql get user", name)
	return User{}
}

type Access struct {

}

func (a *Access) Insert(user User) {
	fmt.Println("Access insert", user)
}

func (a *Access) Get(name string) User {
	fmt.Println("Access get user", name)
	return User{}
}

type DBAbsFactory interface {
	Produce() IUser
}

type MysqlFactory struct {

}

func (MysqlFactory) Produce() IUser {
	return new(Mysql)
}

type AccessFactory struct {

}

func (AccessFactory) Produce() IUser {
	return new(Access)
}



