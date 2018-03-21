package main

import (
	"fmt"
	"testing"
)

func init() {
	db := Connect()
	db.FlushDB()
	db.Close()
}

func TestSetGetGoodSimpleValues(t *testing.T) {
	db := Connect()
	goodItems := []Item{
		Item{Key: "Dog", Value: "Pepper"},
		Item{Key: "", Value: "Empty Key Okay"},
		Item{Key: "Empty Value Okay", Value: ""},
		Item{Key: "Int values", Value: 123},
		Item{Key: "Float values", Value: 123.456},
		Item{Key: "Negative int values", Value: -123},
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
		itemValueStr := fmt.Sprintf("%v", item.Value)
		if value != itemValueStr || err != nil {
			t.Error(err)
		}
	}
	db.FlushDB()
	db.Close()
}

func TestGetMissingKey(t *testing.T) {
	db := Connect()
	_, err := db.GetValue("there is nothing in the db")
	if err != KeyNotFound {
		t.Error(err)
	}
	db.Close()
}
