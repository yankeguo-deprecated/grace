package gracetrack

import (
	"sync"
)

type Track struct {
	groups []*TrackGroup
	lock   sync.Locker
}

type TrackGroup struct {
	key   string
	name  string
	items []string
	lock  sync.Locker
}

func (tg *TrackGroup) SetName(name string) *TrackGroup {
	tg.name = name
	return tg
}

func (tg *TrackGroup) Add(item string) *TrackGroup {
	tg.lock.Lock()
	defer tg.lock.Unlock()
	tg.items = append(tg.items, item)
	return tg
}

func New() *Track {
	return &Track{
		lock: &sync.Mutex{},
	}
}

func (tk *Track) Group(key string) (g *TrackGroup) {
	tk.lock.Lock()
	defer tk.lock.Unlock()

	for _, g = range tk.groups {
		if g.key == key {
			return
		}
	}

	g = &TrackGroup{
		key:  key,
		lock: tk.lock,
	}
	tk.groups = append(tk.groups, g)
	return g
}

func (tk *Track) DumpPlain() (items []string) {
	tk.lock.Lock()
	defer tk.lock.Unlock()

	for _, g := range tk.groups {
		if len(g.items) == 0 {
			continue
		}
		items = append(items, g.name)
		for _, v := range g.items {
			items = append(items, "  * "+v)
		}
	}

	return
}
