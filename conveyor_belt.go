package streamline

import (
	"context"
	"gitee.com/fat_marmota/infra/log"
)

type ConveyorBelt struct {
	DataPanel interface{}
	S *Streamline
	Ctx context.Context
	Logger log.Logger
}

func NewConveyorBelt(s *Streamline, c context.Context, dataDomainRef interface{}) *ConveyorBelt {
	return &ConveyorBelt{
		DataPanel: dataDomainRef,
		S:         s,
		Ctx:       c,
		Logger:    log.GlobalLogger,
	}
}

func (c *ConveyorBelt) Run() error {
	for e:=c.S.procs.Front();e!=nil;e=e.Next() {
		v := e.Value.(Proc)
		c.Logger.Infof("Running process %v", v.Name)
		err := v.F(c)
		if err != nil {
			c.Logger.Errorf("Error when running %v", v.Name)
			return err
		}
	}
	return nil
}