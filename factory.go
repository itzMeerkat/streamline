package streamline

import (
	"container/list"
	"strings"
)

type Factory struct {
	streamlines map[string]*Streamline
}

func New() *Factory {
	return &Factory{streamlines: make(map[string]*Streamline)}
}

func (f *Factory) NewStreamline(name, action, resource string) *Streamline {
	sl := &Streamline{
		procs: list.New(),
		Action: action,
		Resource: resource,
	}

	f.streamlines[name] = sl
	return sl
}

func (f *Factory) Get(name string) *Streamline {
	return f.streamlines[name]
}

func (f *Factory) Plot() string {
	res := strings.Builder{}
	for _,v := range f.streamlines {
		res.WriteString(strings.Join(v.Describe(),","))
		res.WriteRune('\n')
	}
	return res.String()
}

func (f *Factory) GetAllStreamlines() map[string]*Streamline {
	return f.streamlines
}