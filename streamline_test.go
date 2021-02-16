package streamline

import (
	"fmt"
	"go.uber.org/zap"
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

func Reset(ctx *StreamContext, data interface{}) error {
	d := data.(CommonInterface)
	return d.Retract()
}

func Inc(ctx *StreamContext, data interface{}) error {
	d := data.(*MyData)
	d.Counter += 1
	return nil
}

func Mult(ctx *StreamContext, data interface{}) error {
	d := data.(*MyData)
	d.Counter *= 2
	return nil
}

func TestBasic(t *testing.T) {
	l,_:=zap.NewProduction()
	s := New(StreamContext{Logger: l.Sugar()})
	//dm := MyData{Counter: 2}
	wd := WrongData{Nocounter: 222}
	//s.AddProc("1", Inc)
	//s.AddProc("mult", Mult)
	s.Add("interface", Reset)
	s.Run(&wd)
	fmt.Println(wd.Nocounter)
}
