package depmanager

import (
	"github.com/denisdubovitskiy/pbtool/internal/depmanager/datastruct/queue"
	"github.com/denisdubovitskiy/pbtool/internal/depmanager/datastruct/stringset"
)

type Manager struct {
	deps    *queue.Queue
	visited *stringset.Set
}

func NewDeduplicateDependencyManager() *Manager {
	return &Manager{
		deps:    queue.New(),
		visited: stringset.New(),
	}
}

func (m *Manager) EnqueueDependencies(deps ...string) {
	for _, dep := range deps {
		if !m.visited.Has(dep) {
			m.deps.Enqueue(dep)
			m.visited.Add(dep)
		}
	}
}

func (m *Manager) IsEmpty() bool {
	return m.deps.Len() == 0
}

func (m *Manager) Enqueue(dep string) {
	m.deps.Enqueue(dep)
}

func (m *Manager) Dequeue() string {
	return m.deps.Dequeue()
}
