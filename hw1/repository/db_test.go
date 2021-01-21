package repository

import (
	"errors"
	"reflect"
	"shop/models"
	"testing"
)

func TestMapDBCreateItem(t *testing.T) {
	mDB := mapDB{
		db:    make(map[int32]*models.Item, 5),
		maxID: 0,
	}

	currentID := int32(1)
	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_1",
		Price: 10.0,
	}
	currentID++

	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_2",
		Price: 15.0,
	}
	currentID++

	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_3",
		Price: 20.0,
	}
	// TEST BEGINS HERE

	mDB.maxID = currentID

	newItem := &models.Item{
		Name:  "TestName_4",
		Price: 25.0,
	}

	createdItem, err := mDB.CreateItem(newItem)
	if err != nil {
		t.Fatal(err)
	}
	currentID++

	if createdItem.ID != currentID {
		t.Errorf("expected id: %v, got id: %v", currentID, createdItem.ID)
	}
	if createdItem.Name != newItem.Name {
		t.Errorf("expected name: %v, got name: %v", newItem.Name, createdItem.Name)
	}
	if createdItem.Price != newItem.Price {
		t.Errorf("expected price: %v, got price: %v", newItem.Price, createdItem.Price)
	}

	existingItem := mDB.db[currentID]
	if existingItem == nil {
		t.Error("got nil item")
	}

	if existingItem.ID != currentID {
		t.Errorf("expected id: %v, got id: %v", currentID, existingItem.ID)
	}
	if existingItem.Name != newItem.Name {
		t.Errorf("expected name: %v, got name: %v", newItem.Name, existingItem.Name)
	}
	if existingItem.Price != newItem.Price {
		t.Errorf("expected price: %v, got price: %v", newItem.Price, existingItem.Price)
	}
}

func TestMapDBGetItem(t *testing.T) {
	// inserting test data in DB
	mDB := mapDB{
		db:    make(map[int32]*models.Item, 5),
		maxID: 0,
	}

	currentID := int32(1)
	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_1",
		Price: 10.0,
	}
	currentID++

	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_2",
		Price: 15.0,
	}
	currentID++

	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_3",
		Price: 20.0,
	}
	// TEST BEGINS HERE

	testItems := []*models.Item{
		{
			ID:    1,
			Name:  "TestName_1",
			Price: 10.0,
		},
		{
			ID:    2,
			Name:  "TestName_2",
			Price: 15.0,
		},
		{
			ID:    3,
			Name:  "TestName_3",
			Price: 20.0,
		},
	}

	for idx, item := range testItems {
		gottedItem, _ := mDB.GetItem(int32(idx + 1))
		if !reflect.DeepEqual(gottedItem, item) {
			t.Errorf("got %v, expected %v", gottedItem, item)
		}
	}

	var ErrFound = errors.New("Item with ID: 4 is not found")

	_, err := mDB.GetItem(4)
	if !reflect.DeepEqual(err, ErrFound) {
		t.Errorf("gotted unexpected error: %v", err)
	}
}

func TestMapDBDeleteItem(t *testing.T) {
	// inserting test data in DB
	mDB := mapDB{
		db:    make(map[int32]*models.Item, 5),
		maxID: 0,
	}

	currentID := int32(1)
	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_1",
		Price: 10.0,
	}
	currentID++

	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_2",
		Price: 15.0,
	}
	currentID++

	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_3",
		Price: 20.0,
	}
	// TEST BEGINS HERE

	err := mDB.DeleteItem(2)
	if err != nil {
		t.Error("delete fail")
	}

	var ErrFound = errors.New("Item with ID: 2 is not found")

	_, err = mDB.GetItem(2)
	if !reflect.DeepEqual(err, ErrFound) {
		t.Error("delete fail")
	}
}

func TestMapDBUpdateItem(t *testing.T) {
	// inserting test data in DB
	mDB := mapDB{
		db:    make(map[int32]*models.Item, 5),
		maxID: 0,
	}

	currentID := int32(1)
	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_1",
		Price: 10.0,
	}
	currentID++

	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_2",
		Price: 15.0,
	}
	currentID++

	mDB.db[currentID] = &models.Item{
		ID:    currentID,
		Name:  "TestName_3",
		Price: 20.0,
	}
	// TEST BEGINS HERE
	newData := &models.Item{
		ID:    2,
		Name:  "TestName_2-1",
		Price: 17.0,
	}
	updatedItem, _ := mDB.UpdateItem(newData)

	if updatedItem.ID != newData.ID {
		t.Errorf("expected id: %v, got id: %v", newData.ID, updatedItem.ID)
	}
	if updatedItem.Name != newData.Name {
		t.Errorf("expected name: %v, got name: %v", newData.Name, updatedItem.Name)
	}
	if updatedItem.Price != newData.Price {
		t.Errorf("expected price: %v, got price: %v", newData.Price, updatedItem.Price)
	}

	errorData := &models.Item{
		ID:    4,
		Name:  "TestName_4",
		Price: 30.0,
	}
	var ErrFound = errors.New("Item with ID: 4 is not found")

	_, err := mDB.UpdateItem(errorData)
	if !reflect.DeepEqual(err, ErrFound) {
		t.Error("update fail")
	}
}
