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

func Reset(c *ConveyorBelt) error {
	d := c.DataPanel.(CommonInterface)
	return d.Retract()
}

func Inc(c *ConveyorBelt) error {
	d := c.DataPanel.(*MyData)
	d.Counter += 1
	return nil
}

func Mult(c *ConveyorBelt) error {
	d := c.DataPanel.(*MyData)
	d.Counter *= 2
	return nil
}

func Print(c *ConveyorBelt) error {
	d := c.DataPanel.(*MyData)
	fmt.Println(d.Counter)
	return nil
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
