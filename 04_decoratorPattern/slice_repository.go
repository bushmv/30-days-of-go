package main

type SliceRepository struct {
	items []*Item
}

func NewSliceRepository() *SliceRepository {
	return &SliceRepository{
		items: make([]*Item, 0),
	}
}

func (r *SliceRepository) Save(item *Item) (int, error) {
	for i := range r.items {
		if r.items[i].ID == item.ID {
			r.items[i] = item
			return r.items[i].ID, nil
		}
	}
	item.ID = len(r.items) + 1
	r.items = append(r.items, item)
	return item.ID, nil
}

func (r *SliceRepository) GetById(id int) (*Item, error) {
	for i := range r.items {
		if r.items[i].ID == id {
			return r.items[i], nil
		}
	}
	return nil, ErrItemNotFound
}

func (r *SliceRepository) GetAll() ([]*Item, error) {
	copied := make([]*Item, len(r.items))
	copy(copied, r.items)
	return copied, nil
}
