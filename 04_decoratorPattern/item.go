package main

import "fmt"

const NoID int = 0

type Item struct {
	ID    int
	Title string
}

func NewItem(title string) *Item {
	return &Item{
		ID:    NoID,
		Title: title,
	}
}

func (i *Item) String() string {
	return fmt.Sprintf("id: %d, title: %s", i.ID, i.Title)
}
