package streamline

import (
	"testing"
)

type MyData struct {
	Counter int
}


type WrongData struct {
	Nocounter int
}

func (d *MyData)Retract() error {
	d.Counter = 0
	return nil
}
func (d *WrongData)Retract() error {
	d.Nocounter = 10
	return nil
}

type CommonInterface interface {
	Retract() error
}

func Reset(ctx *Context, data interface{}) error {
	d := data.(CommonInterface)
	return d.Retract()
}

func Inc(ctx *Context, data interface{}) error {
	d := data.(*MyData)
	d.Counter += 1
	return nil
}

func Mult(ctx *Context, data interface{}) error {
	d := data.(*MyData)
	d.Counter *= 2
	return nil
}

func TestBasic(t *testing.T) {
	//l,_:=zap.NewProduction()
	f := New()
	s := f.NewStreamline("a", Context{
		Logger:   nil,
		Action:   "act",
		Resource: "res",
	})

	s.Add("interface", Reset)
	//fmt.Println(wd.Nocounter)
}
