package models_test

import (
	"testing"
	"gopher-translator/pkg/mock"
)

func TestTranslationSentence(t *testing.T) {
	englishSentence := "Test test."
	mockInstance := mock.CreateNewMock()
	gopherSentence := mockInstance.TranslateEnglishSentenceToGopher(englishSentence)
	sentenceTranslation := mockInstance.CreateNewTranslationSentenceStructInstance(englishSentence, gopherSentence)

	if sentenceTranslation.EnglishSentence != englishSentence {
		t.Errorf("English sentence expected %v got %v", englishSentence, sentenceTranslation.EnglishSentence)
	}

	if sentenceTranslation.GopherSentence != gopherSentence {
		t.Errorf("Gopher sentence expected %v got %v", gopherSentence, sentenceTranslation.GopherSentence)
	}
}