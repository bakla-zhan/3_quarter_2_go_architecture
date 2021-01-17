package repository

import "shop/models"

type RepositoryMock struct {
}

func (m *RepositoryMock) CreateItem(item *models.Item) (*models.Item, error) {
	return nil, nil
}

func (m *RepositoryMock) GetItem(ID int32) (*models.Item, error) {
	return nil, nil
}

func (m *RepositoryMock) DeleteItem(ID int32) error {
	return nil
}

func (m *RepositoryMock) UpdateItem(item *models.Item) (*models.Item, error) {
	return nil, nil
}
