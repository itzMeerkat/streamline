package streamline

import (
	"container/list"
	"errors"
)
type ProcFunc func(ctx *StreamContext, data interface{}) error

type Proc struct {
	F ProcFunc
	Name string
}

type Streamline struct{
	Name  string
	procs *list.List
	ctx   StreamContext
}


func (s *Streamline) Add(procName string, f ProcFunc) {
	s.procs.PushBack(Proc{
		F:    f,
		Name: procName,
	})
}

func (s *Streamline) Insert(target string, procName string, f ProcFunc, insertBefore bool) error {
	for e:=s.procs.Front();e!=nil;e=e.Next() {
		if e.Value.(Proc).Name == target {
			if insertBefore == true {
				s.procs.InsertBefore(Proc{
					F:    f,
					Name: procName,
				}, e)
			} else {
				s.procs.InsertAfter(Proc{
					F:    f,
					Name: procName,
				}, e)
			}
			return nil
		}
	}
	return errors.New("target process not found")
}

func (s *Streamline) Run(dataModel interface{}) error {
	for e:=s.procs.Front();e!=nil;e=e.Next() {
		v := e.Value.(Proc)
		s.ctx.Logger.Infof("Running process %v", v.Name)
		err := v.F(&s.ctx, dataModel)
		if err != nil {
			s.ctx.Logger.Errorf("Error when running %v", v.Name)
			return err
		}
	}
	return nil
}

func (s *Streamline) Describe() []string {
	var ret []string
	for e:=s.procs.Front();e!=nil;e=e.Next() {
		ret = append(ret, e.Value.(Proc).Name)
	}
	return ret
}