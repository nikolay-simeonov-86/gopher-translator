package models_test

import (
	"testing"
	"gopher-translator/pkg/mock"
)

func TestTranslationWord(t *testing.T) {
	englishWord := "test"
	mockInstance := mock.CreateNewMock()
	gopherWord := mockInstance.TranslateEnglishWordToGopher(englishWord)
	wordTranslation := mockInstance.CreateNewTranslationWordStructInstance(englishWord, gopherWord)

	if wordTranslation.EnglishWord != englishWord {
		t.Errorf("English word expected %v got %v", englishWord, wordTranslation.EnglishWord)
	}

	if wordTranslation.GopherWord != gopherWord {
		t.Errorf("Gopher word expected %v got %v", gopherWord, wordTranslation.GopherWord)
	}
}