package streamline

import (
	"context"
	"errors"
	"fmt"
	"gitee.com/fat_marmota/infra/log"
	"net/http"
)

type ConveyorBelt struct {
	DataDomain interface{}
	S          *Streamline
	Ctx        context.Context
	Logger     log.Logger
	LogInfoGen func(*ConveyorBelt) string
}

func NewConveyorBelt(s *Streamline, c context.Context, dataDomainRef interface{}, f func(*ConveyorBelt)string) *ConveyorBelt {
	if f == nil {
		f = func(belt *ConveyorBelt) string {
			return belt.S.Name
		}
	}
	return &ConveyorBelt{
		DataDomain: dataDomainRef,
		S:          s,
		Ctx:        c,
		Logger:     log.GlobalLogger,
		LogInfoGen: f,
	}
}

func (c *ConveyorBelt) Run() (int, error) {
	var code int
	for e:=c.S.procs.Front();e!=nil;e=e.Next() {
		v := e.Value.(Proc)
		c.Debugw("Executing", v.Name)
		code = v.F(c)
		if code != http.StatusOK {
			return code, errors.New(fmt.Sprintf("Error when running %v", v.Name))
		}
	}
	return code, nil
}

func (c *ConveyorBelt) Infow(args ...interface{}) {
	c.Logger.Infow(c.LogInfoGen(c), args...)
}
func (c *ConveyorBelt) Warnw(args ...interface{}) {
	c.Logger.Warnw(c.LogInfoGen(c), args...)
}
func (c *ConveyorBelt) Errorw(args ...interface{}) {
	c.Logger.Errorw(c.LogInfoGen(c), args...)
}
func (c *ConveyorBelt) Debugw(args ...interface{}) {
	c.Logger.Debugw(c.LogInfoGen(c), args...)
}