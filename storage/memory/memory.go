package memory

import (
	"errors"
	"sync"
)

type MemoryStorage struct {
	data map[string]string
	mu   sync.RWMutex
}

func (m *MemoryStorage) Init() error {
	m.data = make(map[string]string)
	return nil
}

func (m *MemoryStorage) Save(url string, code string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[code] = url
	return nil
}

func (m *MemoryStorage) Load(code string) (string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, exists := m.data[code]
	if !exists {
		return "", errors.New("code not found")
	}
	return value, nil
}
