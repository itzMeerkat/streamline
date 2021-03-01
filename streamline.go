package streamline

import (
	"container/list"
	"errors"
)
type ProcFunc func(*ConveyorBelt) error

type Proc struct {
	F ProcFunc
	Name string
}

type Streamline struct{
	Name  string
	procs *list.List
	Tags []string
	// These are for RBAC authentication
	// If any of them is nil, it means no authentication is enabled
	Action string
	Resource string
}

func (s *Streamline) AddTag(tag string) {
	s.Tags = append(s.Tags, tag)
}

func (s *Streamline) Add(procName string, f ProcFunc) {
	s.procs.PushBack(Proc{
		F:    f,
		Name: procName,
	})
}

func (s *Streamline) Insert(target string, procName string, f ProcFunc, insertBefore bool) error {
	newProc := Proc{
		F:    f,
		Name: procName,
	}

	for e:=s.procs.Front();e!=nil;e=e.Next() {
		if e.Value.(Proc).Name == target {
			if insertBefore == true {
				s.procs.InsertBefore(newProc, e)
			} else {
				s.procs.InsertAfter(newProc, e)
			}
			return nil
		}
	}
	return errors.New("target process not found")
}

func (s *Streamline) Describe() []string {
	var ret []string
	for e:=s.procs.Front();e!=nil;e=e.Next() {
		ret = append(ret, e.Value.(Proc).Name)
	}
	return ret
}

func (s *Streamline) InsertFront(procName string, f ProcFunc) error {
	s.procs.PushFront(Proc{
		F:    f,
		Name: procName,
	})
	return nil
}