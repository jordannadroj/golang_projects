package db

type TodoDatabase interface {
	CloseDB()
	ListItems() ([]Todo, error)
	AddItem(item string) error
	UpdateItem(oldItem, newItem string) error
	DeleteItem(itemID string) error
}
