package store

import "mypass/data"

type Store interface {
	Connect() error
	Get(name string) (*data.PasswordInfo, error)
	GetNames() ([]string, error)
	MatchNames(partial_name string) []string
	Put(name string, val *data.PasswordInfo) error
}
