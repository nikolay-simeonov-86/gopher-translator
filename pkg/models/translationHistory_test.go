package models_test

import (
	"gopher-translator/pkg/mock"
	"testing"
)

func TestTranslationHistoryList(t *testing.T) {
	englishWord := "test"
	mockInstance := mock.CreateNewMock()
	gopherWord := mockInstance.TranslateEnglishWordToGopher(englishWord)
	historyList := mockInstance.CreateNewTranslationHistoryStructInstance(englishWord,gopherWord)
	for _, val := range historyList.History {
		for _, mapVal := range val {
			if mapVal != gopherWord {
				t.Errorf("Expected %v, got %v", gopherWord, mapVal)
			}
		}
	}
}