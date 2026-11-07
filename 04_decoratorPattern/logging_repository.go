package main

import (
	"fmt"
	"io"
)

type LoggingRepository struct {
	base Repository
	to   io.Writer
}

func NewLoggingRepository(base Repository, to io.Writer) *LoggingRepository {
	return &LoggingRepository{
		base: base,
		to:   to,
	}
}

func (r *LoggingRepository) Save(item *Item) (int, error) {
	id, err := r.base.Save(item)
	log := fmt.Sprintf("Save: item=%v -> id=%v, err=%v\n", item, id, err)
	r.to.Write([]byte(log))
	return id, err
}
func (r *LoggingRepository) GetById(id int) (*Item, error) {
	item, err := r.base.GetById(id)
	log := fmt.Sprintf("GetById: id=%v -> item=%v, err=%v\n", id, item, err)
	r.to.Write([]byte(log))
	return item, err
}

func (r *LoggingRepository) GetAll() ([]*Item, error) {
	items, err := r.base.GetAll()
	log := fmt.Sprintf("GetById: -> items=%v, err=%v\n", items, err)
	r.to.Write([]byte(log))
	return items, err
}
