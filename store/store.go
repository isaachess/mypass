package store

import (
	"encoding/json"
	"errors"
	"io"
	"mypass/data"
	"os"
	"sync"
)

var ErrorNotFound = errors.New("Not found")

type JSONStore struct {
	path string
	file *os.File

	pwds_mu sync.RWMutex
	pwds    map[string]*data.PasswordInfo
}

func NewJSONStore(path string) *JSONStore {
	return &JSONStore{
		path: path,
		pwds: make(map[string]*data.PasswordInfo),
	}
}

func (s *JSONStore) Connect() error {
	f, err := os.Open(s.path)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := json.NewDecoder(f).Decode(&s.pwds); err != nil {
		if err == io.EOF {
			return nil
		}
		return err
	}

	return nil
}

func (s *JSONStore) Close() error {
	return s.file.Close()
}

func (s *JSONStore) Get(name string) (*data.PasswordInfo, error) {
	p, ok := s.readPwds(name)
	if !ok {
		return nil, ErrorNotFound
	}
	return p, nil
}

func (s *JSONStore) GetNames() ([]string, error) {
	s.pwds_mu.Lock()
	defer s.pwds_mu.Unlock()
	var names []string
	for name, _ := range s.pwds {
		names = append(names, name)
	}
	return names, nil
}

func (s *JSONStore) Put(name string, val *data.PasswordInfo) error {
	// Put adds the pwd info to the map and writes json to the file
	s.writePwds(name, val)
	f, err := os.Create(s.path)
	if err != nil {
		return err
	}
	return json.NewEncoder(f).Encode(s.pwds)
}

func (s *JSONStore) readPwds(name string) (*data.PasswordInfo, bool) {
	s.pwds_mu.RLock()
	defer s.pwds_mu.RUnlock()
	p, ok := s.pwds[name]
	return p, ok
}

func (s *JSONStore) writePwds(name string, val *data.PasswordInfo) {
	s.pwds_mu.Lock()
	defer s.pwds_mu.Unlock()
	s.pwds[name] = val
}
