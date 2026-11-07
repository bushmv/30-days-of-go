package main

type CachingRepository struct {
	base  Repository
	cache map[int]*Item
}

func (r *CachingRepository) Save(item *Item) (int, error) {
	id, err := r.base.Save(item)
	if err != nil {
		return 0, err
	}
	r.cache[id] = item
	return id, nil
}

func (r *CachingRepository) GetById(id int) (*Item, error) {
	if item, ok := r.cache[id]; ok {
		return item, nil
	}
	item, err := r.base.GetById(id)
	if err != nil {
		return nil, err
	}
	r.cache[id] = item
	return item, nil
}

// don't cache all items
func (r *CachingRepository) GetAll() ([]*Item, error) {
	return r.base.GetAll()
}
