package streamline

import (
	"container/list"
	"strings"
)

type Factory struct {
	streamlines []*Streamline
}

func (f *Factory) New(meta Context) *Streamline {
	sl := &Streamline{
		procs: list.New(),
		ctx:   meta,
	}

	f.streamlines = append(f.streamlines, sl)

	return sl
}

func (f *Factory) Plot() string {
	res := strings.Builder{}
	for _,v := range f.streamlines {
		res.WriteString(strings.Join(v.Describe(),","))
		res.WriteRune('\n')
	}
	return res.String()
}