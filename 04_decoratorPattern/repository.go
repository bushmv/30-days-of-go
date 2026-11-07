package main

import "errors"

var ErrItemNotFound = errors.New("item not found")

type Repository interface {
	Save(item *Item) (int, error)
	GetById(id int) (*Item, error)
	GetAll() ([]*Item, error)
}
