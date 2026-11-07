package main

type EncryptingRepository struct {
	base        Repository
	encryptFunc func(item *Item) *Item
	decryptFunc func(item *Item) *Item
	isEncrypted func(item *Item) bool
}

func NewEncryptingRepository(base Repository, encryptFunc func(item *Item) *Item, decryptFunc func(item *Item) *Item, isEncrypted func(item *Item) bool) *EncryptingRepository {
	return &EncryptingRepository{
		base:        base,
		encryptFunc: encryptFunc,
		decryptFunc: decryptFunc,
		isEncrypted: isEncrypted,
	}
}

func (r *EncryptingRepository) Save(item *Item) (int, error) {
	itemCopy := &Item{ID: item.ID, Title: item.Title}
	encryptedItem := r.encryptFunc(itemCopy)
	return r.base.Save(encryptedItem)
}

func (r *EncryptingRepository) GetById(id int) (*Item, error) {
	item, err := r.base.GetById(id)
	if err != nil {
		return nil, err
	}
	if r.isEncrypted(item) {
		item = r.decryptFunc(item)
	}
	return item, nil
}

func (r *EncryptingRepository) GetAll() ([]*Item, error) {
	items, err := r.GetAll()
	if err != nil {
		return nil, err
	}
	for i, item := range items {
		if r.isEncrypted(item) {
			items[i] = r.decryptFunc(item)
		}
	}
	return items, nil
}
