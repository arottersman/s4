package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockDB struct{ Item }

func (db *mockDB) GetValue(key string) (string, error) {
	itemValueStr := fmt.Sprintf("%v", db.Item.Value)
	return itemValueStr, nil
}
func (db *mockDB) SetValue(item Item) error {
	db.Item = item
	return nil
}
func (db *mockDB) Close() error   { return nil }
func (db *mockDB) FlushDB() error { return nil }

type erroringMockDB struct{}

func (db *erroringMockDB) GetValue(key string) (string, error) {
	return "", errors.New("DB error!")
}
func (db *erroringMockDB) SetValue(item Item) error {
	return errors.New("DB error!")
}
func (db *erroringMockDB) Close() error   { return nil }
func (db *erroringMockDB) FlushDB() error { return nil }

type keyNotFoundMockDB struct{}

func (db *keyNotFoundMockDB) GetValue(key string) (string, error) {
	return "", KeyNotFound
}
func (db *keyNotFoundMockDB) SetValue(item Item) error { return nil }
func (db *keyNotFoundMockDB) Close() error             { return nil }
func (db *keyNotFoundMockDB) FlushDB() error           { return nil }

func InitMockDB(d DB) {
	db = d
}

func TestSetHandler(t *testing.T) {
	goodDB := &mockDB{}
	InitMockDB(goodDB)

	data, _ := json.Marshal(Item{Key: "ShouldBeOkay", Value: "Yes!"})
	body := bytes.NewBuffer(data)

	req, _ := http.NewRequest("PUT", "/items", body)
	res := httptest.NewRecorder()
	CreateRouter().ServeHTTP(res, req)

	statusCode := res.Result().StatusCode
	if statusCode != http.StatusCreated {
		t.Error("Set request failed with status code:", statusCode)
	}
	if value, err := goodDB.GetValue("ShouldBeOkay"); value != "Yes!" && err != nil {
		t.Error("Set didn't set proper value. value: ", value, "error: ", err)
	}
}

func TestGetHandler(t *testing.T) {
	item := Item{Key: "ShouldBeOkay", Value: "Yes!"}
	goodDB := &mockDB{Item: item}
	InitMockDB(goodDB)

	req, _ := http.NewRequest("GET", "/items/ShouldBeOkay", nil)
	res := httptest.NewRecorder()
	CreateRouter().ServeHTTP(res, req)

	statusCode := res.Result().StatusCode
	if statusCode != http.StatusOK {
		t.Error("Get request failed with status code:", statusCode)
	}

	expectedResponseBody, _ := json.Marshal(item)
	if strings.TrimSpace(res.Body.String()) != string(expectedResponseBody) {
		t.Error("Did not return proper value for key. Got", res.Body.String(), "Expected", string(expectedResponseBody))
	}
}

func TestGetHandlerMissingKey(t *testing.T) {
	missingKeyDB := &keyNotFoundMockDB{}
	InitMockDB(missingKeyDB)

	req, _ := http.NewRequest("GET", "/items/123", nil)
	res := httptest.NewRecorder()
	CreateRouter().ServeHTTP(res, req)

	statusCode := res.Result().StatusCode
	if statusCode != http.StatusNotFound {
		t.Error("Expected Get missing key to respond with 404. Got ", statusCode)
	}
}

func TestSetHandlerErroringDB(t *testing.T) {
	erroringDB := &erroringMockDB{}
	InitMockDB(erroringDB)

	data, _ := json.Marshal(Item{Key: "ShouldBeOkay", Value: "Yes!"})
	body := bytes.NewBuffer(data)

	req, _ := http.NewRequest("PUT", "/items", body)
	res := httptest.NewRecorder()
	CreateRouter().ServeHTTP(res, req)

	statusCode := res.Result().StatusCode
	if statusCode != http.StatusInternalServerError {
		t.Error("Expected error on Set to respond with 500. Got ", statusCode)
	}
}

func TestGetHandlerErroringDB(t *testing.T) {
	erroringDB := &erroringMockDB{}
	InitMockDB(erroringDB)

	req, _ := http.NewRequest("GET", "/items/123", nil)
	res := httptest.NewRecorder()
	CreateRouter().ServeHTTP(res, req)

	statusCode := res.Result().StatusCode
	if statusCode != http.StatusInternalServerError {
		t.Error("Expected error on Get to respond with 500. Got ", statusCode)
	}
}

func TestSetHandlerBadRequestBody(t *testing.T) {
	goodDB := &mockDB{}
	InitMockDB(goodDB)

	data, _ := json.Marshal("{Iam:bad}")
	body := bytes.NewBuffer(data)

	req, _ := http.NewRequest("PUT", "/items", body)
	res := httptest.NewRecorder()
	CreateRouter().ServeHTTP(res, req)

	statusCode := res.Result().StatusCode
	if statusCode != http.StatusBadRequest {
		t.Error("Expected error on Set to respond with 400. Got ", statusCode)
	}
}

func TestSetHandlerNoRequestBody(t *testing.T) {
	goodDB := &mockDB{}
	InitMockDB(goodDB)

	req, _ := http.NewRequest("PUT", "/items", nil)
	res := httptest.NewRecorder()
	CreateRouter().ServeHTTP(res, req)

	statusCode := res.Result().StatusCode
	if statusCode != http.StatusBadRequest {
		t.Error("Expected error on Set to respond with 400. Got ", statusCode)
	}
}
