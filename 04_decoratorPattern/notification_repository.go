package main

import "fmt"

type Notifier interface {
	Notify(err error)
}

type ConsoleNotifier struct{}

func NewConsoleNotifier() *ConsoleNotifier {
	return &ConsoleNotifier{}
}

func (n *ConsoleNotifier) Notify(err error) {
	fmt.Printf("notification error: %v\n", err)
}

type NotificationRepository struct {
	base     Repository
	notifier Notifier
}

func NewNotificationRepository(base Repository, notifier Notifier) *NotificationRepository {
	return &NotificationRepository{
		base:     base,
		notifier: notifier,
	}
}

func (r *NotificationRepository) Save(item *Item) (int, error) {
	id, err := r.base.Save(item)
	if err != nil {
		r.notifier.Notify(err)
	}
	return id, err
}

func (r *NotificationRepository) GetById(id int) (*Item, error) {
	item, err := r.base.GetById(id)
	if err != nil {
		r.notifier.Notify(err)
	}
	return item, err
}

func (r *NotificationRepository) GetAll() ([]*Item, error) {
	items, err := r.base.GetAll()
	if err != nil {
		r.notifier.Notify(err)
	}
	return items, err
}
