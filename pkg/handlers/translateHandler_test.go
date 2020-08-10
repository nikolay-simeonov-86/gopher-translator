package handlers_test

import (
	"reflect"
	"gopher-translator/pkg/models"
	"encoding/json"
	"net/http/httptest"
	"net/http"
	"gopher-translator/pkg/handlers"
	"gopher-translator/pkg/mock"
	"testing"
)

func TestCreateNewTranslatorHandler(t *testing.T)  {
	mockInstance := mock.CreateNewMock()
	repository := mockInstance.CreateNewRedisRepository()
	translator := mockInstance.CreateNewGopherTranslator()
	handler := mockInstance.CreateNewTranslationHandler(repository, translator)

	switch v := handler.(type) {
	case handlers.TranstatorHandler:
		break
	default:
		t.Errorf("Epected type handlers.TranstatorHandler got %v", v)
	}
}

func TestWelcome(t *testing.T)  {
	mockInstance := mock.CreateNewMock()
	repository := mockInstance.CreateNewRedisRepository()
	translator := mockInstance.CreateNewGopherTranslator()
	handler := mockInstance.CreateNewTranslationHandler(repository, translator)
	welcomeRes := handler.Welcome()
	
	req, err := http.NewRequest("GET", "/welcome", nil)
    if err != nil {
        t.Fatal(err)
	}
	
	res := httptest.NewRecorder()
	welcomeRes(res, req)

	if status := res.Code; status != http.StatusOK {
        t.Errorf("Welcome handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
	}
	
	expectedWelcomeRes := mockInstance.CreateNewTranslationWelcomeStructInstance(
		"Welcome to Gopher language Translator!",
		map[string]string{
			"POST": "/word /sentence",
			"GET": "/history",
		},
	)

	actualWelcomeRes := models.TranslationWelcome{}
	err = json.NewDecoder(res.Body).Decode(&actualWelcomeRes)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	
	if expectedWelcomeRes.Welcome != actualWelcomeRes.Welcome {
		t.Errorf("Response body Welcome key expected %v got %v", expectedWelcomeRes.Welcome, actualWelcomeRes.Welcome)
	}

	eq := reflect.DeepEqual(expectedWelcomeRes.Routes, actualWelcomeRes.Routes)
	if !eq {
		t.Errorf("Response body Routes keys don't match")
	}
}