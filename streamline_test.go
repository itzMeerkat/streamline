package streamline

import (
	"context"
	"fmt"
	"gitee.com/fat_marmota/infra/log"
	"testing"
)

type MyData struct {
	Counter int
}


type WrongData struct {
	Nocounter int
}

func (d *MyData)Retract() int {
	d.Counter = 0
	return 200
}
func (d *WrongData)Retract() int {
	d.Nocounter = 10
	return 200
}

type CommonInterface interface {
	Retract() int
}

func Reset(c *ConveyorBelt) int {
	d := c.DataDomain.(CommonInterface)
	return d.Retract()
}

func Inc(c *ConveyorBelt) int {
	d := c.DataDomain.(*MyData)
	d.Counter += 1
	return 200
}

func Mult(c *ConveyorBelt) int {
	d := c.DataDomain.(*MyData)
	d.Counter *= 2
	return 200
}

func Print(c *ConveyorBelt) int {
	d := c.DataDomain.(*MyData)
	fmt.Println(d.Counter)
	return 200
}

func TestBasic(t *testing.T) {
	log.InitZapSugared(true, false, 1)
	f := New()
	s := f.NewStreamline("a","b","c")

	s.Add("interface", Reset)
	s.Add("add",Inc)
	s.Add("add",Inc)
	s.Add("p", Print)
	c := NewConveyorBelt(s, context.Background(), &MyData{})
	c.Run()
	//fmt.Println(wd.Nocounter)
}
