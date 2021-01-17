package service

import (
	"errors"
	"shop/models"
	rep "shop/repository"
)

type Service interface {
	CreateItem(item *models.Item) (*models.Item, error)
}

type service struct {
	db rep.Repository
}

func (s *service) CreateItem(item *models.Item) (*models.Item, error) {
	if item.Name == "" {
		return nil, errors.New("item name is empty")
	}
	if item.Price <= 0 {
		return nil, errors.New("item price should be positive")
	}

	return s.db.CreateItem(item)
}

func NewService(db rep.Repository) Service {
	return &service{
		db: db,
	}
}
