# Decorator pattern

У нас будет структура `item`, которая нужна программе для выполнения некоторой логики. `Item` состоит из 2-х полей

- `ID` - уникальный идентификатор
- `Title` - некоторое описание



По мере выполнения программе нужно сохранять, получать `item`

- `Save` - сохранить `item` и назначить ему `id`
- `GetById` - получить `item` по `id`
- `GetAll` - получить все `item`

```go
type Repository interface {
	Save(item *Item) (int, error)
	GetById(id int) (*Item, error)
	GetAll() ([]*Item, error)
}
```
