package main

import "errors"

var ErrValidationFailed = errors.New("item validation failed")

type FilteringRepository struct {
	base            Repository
	onSavePredicate func(item *Item) bool
}

func (r *FilteringRepository) Save(item *Item) (int, error) {
	if !r.onSavePredicate(item) {
		return 0, ErrValidationFailed
	}
	return r.Save(item)
}

func (r *FilteringRepository) GetById(id int) (*Item, error) {
	return r.base.GetById(id)

}

func (r *FilteringRepository) GetAll() ([]*Item, error) {
	return r.base.GetAll()
}
