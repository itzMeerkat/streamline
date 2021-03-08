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
}

func NewConveyorBelt(s *Streamline, c context.Context, dataDomainRef interface{}) *ConveyorBelt {
	return &ConveyorBelt{
		DataDomain: dataDomainRef,
		S:          s,
		Ctx:        c,
		Logger:     log.GlobalLogger,
	}
}

func (c *ConveyorBelt) Run() (int, error) {
	var code int
	for e:=c.S.procs.Front();e!=nil;e=e.Next() {
		v := e.Value.(Proc)
		c.Logger.Debugf("Running process %v", v.Name)
		code = v.F(c)
		if code != http.StatusOK {
			return code, errors.New(fmt.Sprintf("Error when running %v", v.Name))
		}
	}
	return code, nil
}
