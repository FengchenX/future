package factory

//工厂模式适合: 凡是出现了大量的产品创建
//并且具有共同的接口时, 可以通过工厂方法模式进行创建

import "fmt"

type pen interface {
	Write()
}

type pencil struct {

}

func (p *pencil) Write() {
	fmt.Println("铅笔")
}

type brushPen struct {

}

func (p *brushPen) Write() {
	fmt.Println("毛笔")
}

//PenFactory 工厂
type PenFactory struct {

}

func (pf PenFactory) Produce(typ string) pen {
	switch typ {
	case "pencil":
		return pf.ProducePencil()
	case "brush":
		return pf.ProduceBrushPen()
	default:
		return nil
	}
}

func (PenFactory) ProducePencil() pen {
	return new(pencil)
}

func (PenFactory) ProduceBrushPen() pen {
	return new(brushPen)
}

type operationer interface {
	Operation(a, b float64) float64
}

type add struct {

}

type sub struct {

}

type mul struct {

}

type div struct {

}

func (a *add) Operation(x, y float64) float64 {
	return x + y
}

func (s *sub) Operation(x, y float64) float64 {
	return x - y
}

func (m *mul) Operation(x, y float64) float64 {
	return x * y
}

func (d *div) Operation(x, y float64) float64 {
	return x / y
}

type OperationFactory struct {

}

func (of OperationFactory) Produce(typ string) operationer {
	switch typ {
	case "+":
		return new(add)
	case "-":
		return new(sub)
	case "*":
		return new(mul)
	case "/":
		return new(div)
	default:
		return nil
	}
}

