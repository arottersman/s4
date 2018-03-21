package main

import "testing"

func TestSetGetGoodSimpleValues(t *testing.T) {
	db := Connect()
	goodItems := []Item{
		Item{Key: "Dog", Value: "Pepper"},
		Item{Key: "", Value: "Empty Key Okay"},
		Item{Key: "Empty Value Okay", Value: ""},
		Item{Key: "Numeric values", Value: 123},
	}

	for _, item := range goodItems {
		err := db.SetValue(item)
		if err != nil {
			t.Error(err)
		}
	}
	for _, item := range goodItems {
		value, err := db.GetValue(item.Key)
		if value != item.Value || err != nil {
			t.Error(err)
		}
	}
	db.FlushAll()
	db.Close()
}

func TestGetMissingKey(t *testing.T) {
	db := Connect()
	value, err := db.GetValue("there is nothing in the db")
	if value != nil && err != nil {
		t.Error(err)
	}
	db.Close()
}
