package cmd

import (
	"flag"
	"fmt"
	"mypass/store"
	"sort"
)

const ListName = "list"

var ListFlags = flag.NewFlagSet(ListName, flag.ExitOnError)

type List struct {
	store *store.JSONStore
}

func NewList(store *store.JSONStore) *List {
	return &List{
		store: store,
	}
}

func (l *List) Run(args []string) error {
	if len(args) > 0 {
		return l.listOne(args[0])
	}
	return l.listAll()
}

func (l *List) listOne(name string) error {
	pw, err := l.store.Get(name)
	if err != nil {
		return err
	}

	fmt.Println(pw)

	return nil
}

func (l *List) listAll() error {
	names, err := l.store.GetNames()
	if err != nil {
		return err
	}

	sort.Strings(names)
	for _, name := range names {
		fmt.Println(name)
	}

	return nil
}
