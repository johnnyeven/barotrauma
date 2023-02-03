package core

import (
	"time"
)

type World struct {
	updateInterval int64
	lastUpdate     int64
	currentTime    int64
	components     map[string]IComponent
}

func NewWorld(interval int64) *World {
	return &World{
		updateInterval: interval,
		lastUpdate:     0,
		currentTime:    0,
		components:     make(map[string]IComponent),
	}
}

func (w *World) Update() {
	w.currentTime = time.Now().UnixMilli()
	if w.currentTime-w.lastUpdate >= w.updateInterval {
		for _, component := range w.components {
			component.Update(w.currentTime)
		}
		w.lastUpdate = w.currentTime
	}
}

func (w *World) AttachComponent(c IComponent) {
	w.components[c.Id()] = c
}

func (w *World) RemoveComponent(id string) IComponent {
	if c, ok := w.components[id]; ok {
		return c
	}
	return nil
}
