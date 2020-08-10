package handlers_test

import (
	"bytes"
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
	
	req, err := http.NewRequest("GET", "/", nil)
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
		t.Errorf("Welcome response body Welcome key expected %v got %v", expectedWelcomeRes.Welcome, actualWelcomeRes.Welcome)
	}

	eq := reflect.DeepEqual(expectedWelcomeRes.Routes, actualWelcomeRes.Routes)
	if !eq {
		t.Errorf("Welcome response body Routes keys don't match")
	}
}

func TestTranslateWord(t *testing.T)  {
	englishWord := "test"
	mockInstance := mock.CreateNewMock()
	gopherWord := mockInstance.TranslateEnglishWordToGopher(englishWord)
	expectedTranslateWordRes := mockInstance.CreateNewTranslationWordStructInstance(englishWord, gopherWord)
	repository := mockInstance.CreateNewRedisRepository()
	translator := mockInstance.CreateNewGopherTranslator()
	handler := mockInstance.CreateNewTranslationHandler(repository, translator)
	translateWordRes := handler.TranslateWord()
	reqBody, err := json.Marshal(expectedTranslateWordRes)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/word", bytes.NewBuffer(reqBody))
    if err != nil {
        t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	
	res := httptest.NewRecorder()
	translateWordRes(res, req)

	if status := res.Code; status != http.StatusOK {
        t.Errorf("TranslateWord handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
	}

	actualTranslateWordRes := models.TranslationWord{}
	err = json.NewDecoder(res.Body).Decode(&actualTranslateWordRes)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if expectedTranslateWordRes.GopherWord != actualTranslateWordRes.GopherWord {
		t.Errorf(
			"TranslateWord response body gopher word expected %v got %v", 
			expectedTranslateWordRes.GopherWord,
			actualTranslateWordRes.GopherWord,
		)
	}
}

func TestTranslateSentence(t *testing.T)  {
	englishSentence := "Test test."
	mockInstance := mock.CreateNewMock()
	gopherSentence := mockInstance.TranslateEnglishSentenceToGopher(englishSentence)
	expectedTranslateSentenceRes := mockInstance.CreateNewTranslationSentenceStructInstance(englishSentence, gopherSentence)
	repository := mockInstance.CreateNewRedisRepository()
	translator := mockInstance.CreateNewGopherTranslator()
	handler := mockInstance.CreateNewTranslationHandler(repository, translator)
	translateSentenceRes := handler.TranslateSentence()
	reqBody, err := json.Marshal(expectedTranslateSentenceRes)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/sentence", bytes.NewBuffer(reqBody))
    if err != nil {
        t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	
	res := httptest.NewRecorder()
	translateSentenceRes(res, req)

	if status := res.Code; status != http.StatusOK {
        t.Errorf("TranslateSentence handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
	}

	actualTranslateSentenceRes := models.TranslationSentence{}
	err = json.NewDecoder(res.Body).Decode(&actualTranslateSentenceRes)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	if expectedTranslateSentenceRes.GopherSentence != actualTranslateSentenceRes.GopherSentence {
		t.Errorf(
			"TranslateSentence response body gopher sentence expected %v got %v", 
			expectedTranslateSentenceRes.GopherSentence,
			actualTranslateSentenceRes.GopherSentence,
		)
	}
}

func TestTranslateHistory(t *testing.T)  {
	mockInstance := mock.CreateNewMock()
	repository := mockInstance.CreateNewRedisRepository()
	translator := mockInstance.CreateNewGopherTranslator()
	handler := mockInstance.CreateNewTranslationHandler(repository, translator)
	translateHistoryRes := handler.TranslateHistory()

	req, err := http.NewRequest("GET", "/history", nil)
    if err != nil {
        t.Fatal(err)
	}
	
	res := httptest.NewRecorder()
	translateHistoryRes(res, req)

	if status := res.Code; status != http.StatusOK {
        t.Errorf("TranslateHistory handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
	}

	expectedTranslationHistoryRes := mockInstance.GetAllTranslationHistoryRecords()

	actualTranslateHistoryRes := models.TranslationHistoryList{}
	err = json.NewDecoder(res.Body).Decode(&actualTranslateHistoryRes)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	eq := reflect.DeepEqual(expectedTranslationHistoryRes.History, actualTranslateHistoryRes.History)
	if !eq {
		t.Errorf("TranslateHistory response body History keys don't match")
	}
}

func TestAddNewWordTranslation(t *testing.T) {
	englishWord := "test"
	mockInstance := mock.CreateNewMock()
	gopherWord := mockInstance.TranslateEnglishWordToGopher(englishWord)
	expectedTranslation := mockInstance.CreateNewTranslationWordStructInstance(englishWord, gopherWord)
	repository := mockInstance.CreateNewRedisRepository()
	translator := mockInstance.CreateNewGopherTranslator()
	actualTranslation, err := handlers.AddNewWordTranslation(englishWord, translator, repository)
	if err != nil {
		t.Fatal(err)
	}

	if expectedTranslation.GopherWord != actualTranslation.GopherWord {
		t.Errorf(
			"addNewWordTranslation response GopherWord expected %v got %v", 
			expectedTranslation.GopherWord,
			actualTranslation.GopherWord,
		)
	}
}

func TestAddNewSentenceTranslation(t *testing.T) {
	englishSentence := "Test test."
	mockInstance := mock.CreateNewMock()
	gopherSentencce := mockInstance.TranslateEnglishSentenceToGopher(englishSentence)
	expectedTranslation := mockInstance.CreateNewTranslationSentenceStructInstance(englishSentence, gopherSentencce)
	repository := mockInstance.CreateNewRedisRepository()
	translator := mockInstance.CreateNewGopherTranslator()
	actualTranslation, err := handlers.AddNewSentenceTranslation(englishSentence, translator, repository)
	if err != nil {
		t.Fatal(err)
	}

	if expectedTranslation.GopherSentence != actualTranslation.GopherSentence {
		t.Errorf(
			"addNewSentenceTranslation response GopherSentence expected %v got %v", 
			expectedTranslation.GopherSentence,
			actualTranslation.GopherSentence,
		)
	}
}

func TestValidateRequestBody(t *testing.T) {
	mockInstance := mock.CreateNewMock()
	englishWord := "test"
	gopherWord := mockInstance.TranslateEnglishWordToGopher(englishWord)
	translationWord := mockInstance.CreateNewTranslationWordStructInstance(englishWord, gopherWord)
	errMsg := handlers.ValidateRequestBody(translationWord)
	if errMsg != "" {
		t.Errorf("ValidateRequestBody response for TranslationWordStruct instance expected no error got %v", errMsg)
	}

	englishSentence := "Test test."
	gopherSentence := mockInstance.TranslateEnglishSentenceToGopher(englishSentence)
	translationSentence := mockInstance.CreateNewTranslationSentenceStructInstance(englishSentence, gopherSentence)
	errMsg = handlers.ValidateRequestBody(translationSentence)
	if errMsg != "" {
		t.Errorf("ValidateRequestBody response for TranslationSentenceStruct instance expected no error got %v", errMsg)
	}
}