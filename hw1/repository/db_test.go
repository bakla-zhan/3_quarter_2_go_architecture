package repository

import (
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
		t.Fatal("expected id == ")
	}
	if createdItem.Name != newItem.Name {
		t.Fatal("expected name == ")
	}
	if createdItem.Price != newItem.Price {
		t.Fatal("expected name == ")
	}

	existingItem := mDB.db[currentID]
	if existingItem == nil {
		t.Fatal("got nil item")
	}

	if existingItem.ID != currentID {
		t.Fatal("expected id == ")
	}
	if existingItem.Name != newItem.Name {
		t.Fatal("expected name == ")
	}
	if existingItem.Price != newItem.Price {
		t.Fatal("expected name == ")
	}
}
