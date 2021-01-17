package repository

import (
	"fmt"

	"shop/models"
)

type Repository interface {
	CreateItem(item *models.Item) (*models.Item, error)
	GetItem(ID int32) (*models.Item, error)
	DeleteItem(ID int32) error
	UpdateItem(item *models.Item) (*models.Item, error)
}

type mapDB struct {
	db    map[int32]*models.Item
	maxID int32
}

func NewMapDB() Repository {
	return &mapDB{
		db:    make(map[int32]*models.Item),
		maxID: 0,
	}
}

func (m *mapDB) CreateItem(item *models.Item) (*models.Item, error) {
	m.maxID++

	newItem := &models.Item{
		ID:    m.maxID,
		Price: item.Price,
		Name:  item.Name,
	}

	m.db[newItem.ID] = newItem

	return &models.Item{
		ID:    newItem.ID,
		Name:  newItem.Name,
		Price: newItem.Price,
	}, nil
}

func (m *mapDB) GetItem(ID int32) (*models.Item, error) {
	item, ok := m.db[ID]
	if !ok {
		return nil, fmt.Errorf("Item with ID: %d is not found", ID)
	}

	return &models.Item{
		ID:    item.ID,
		Name:  item.Name,
		Price: item.Price,
	}, nil
}

func (m *mapDB) DeleteItem(ID int32) error {
	delete(m.db, ID)
	return nil
}

func (m *mapDB) UpdateItem(item *models.Item) (*models.Item, error) {
	updateItem, ok := m.db[item.ID]
	if !ok {
		return nil, fmt.Errorf("Item with ID: %d is not found", item.ID)
	}
	updateItem.Name = item.Name
	updateItem.Price = item.Price

	return &models.Item{
		ID:    updateItem.ID,
		Name:  updateItem.Name,
		Price: updateItem.Price,
	}, nil
}
