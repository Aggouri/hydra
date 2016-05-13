package connection

import (
	"github.com/go-errors/errors"
	"github.com/ory-am/hydra/pkg"
)

type MemoryManager struct {
	Connections map[string]*Connection
}

func NewMemoryManager() *MemoryManager {
	return &MemoryManager{
		Connections: make(map[string]*Connection),
	}
}

func (m *MemoryManager) Create(c *Connection) error {
	m.Connections[c.GetID()] = c
	return nil
}

func (m *MemoryManager) Delete(id string) error {
	delete(m.Connections, id)
	return nil
}

func (m *MemoryManager) Get(id string) (*Connection, error) {
	c, ok := m.Connections[id]
	if !ok {
		return nil, errors.New(pkg.ErrNotFound)
	}
	return c, nil
}

func (m *MemoryManager) FindAllByLocalSubject(subject string) ([]*Connection, error) {
	var cs []*Connection
	for _, c := range m.Connections {
		if c.GetLocalSubject() == subject {
			cs = append(cs, c)
		}
	}
	return cs, nil
}

func (m *MemoryManager) FindByRemoteSubject(provider, subject string) (*Connection, error) {
	for _, c := range m.Connections {
		if c.GetProvider() == provider && c.GetRemoteSubject() == subject {
			return c, nil
		}
	}
	return nil, errors.New(pkg.ErrNotFound)
}