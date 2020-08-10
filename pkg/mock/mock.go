package mock

import (
	"gopher-translator/pkg/models"
	"gopher-translator/pkg/handlers"
	"gopher-translator/pkg/translation"
	"gopher-translator/pkg/storage"
)

// NewMock returns a new mock struct
type NewMock struct {}

// Mock is an interface with all mock functions
type Mock interface {
	CreateNewRedisRepository() storage.Repository
	CreateNewGopherTranslator() translation.GopherTranslator
	CreateNewTranslationHandler(repository storage.Repository, translator translation.GopherTranslator) handlers.TranstatorHandler
	CreateNewTranslationHistoryStructInstance() models.TranslationHistoryList
}

// CreateNewMock creates a new Mock struct instance
func CreateNewMock () *NewMock {
	return &NewMock{}
}

// CreateNewRedisRepository creates a new redis repository
func (mock *NewMock) CreateNewRedisRepository() storage.Repository {
	return storage.CreateNewRedisStorage("localhost:6379","",0)
}

// CreateNewGopherTranslator creates a new gopher translator
func (mock *NewMock) CreateNewGopherTranslator() translation.GopherTranslator {
	return translation.CreateNewGopherTranslator()
}

// CreateNewTranslationHandler creates a new translation handler
func (mock *NewMock) CreateNewTranslationHandler(repository storage.Repository, translator translation.GopherTranslator) handlers.TranstatorHandler {
	return handlers.CreateNewTranslatorHandler(repository, translator)
}

// CreateNewTranslationHistoryStructInstance creates a new translation handler
func (mock *NewMock) CreateNewTranslationHistoryStructInstance(englishWord string, gopherWord string) models.TranslationHistoryList {
	return models.TranslationHistoryList{
		History: []map[string]string{
			{
				englishWord: gopherWord,
			},
		},
	}
}

// CreateNewTranslationWordStructInstance creates a new instance of model TranslationWord
func (mock *NewMock) CreateNewTranslationWordStructInstance(englishWord string, gopherWord string) models.TranslationWord {
	return models.TranslationWord{
		EnglishWord: englishWord,
		GopherWord: gopherWord,
	}
}

// CreateNewTranslationSentenceStructInstance creates a new instance of model TranslationSentence
func (mock *NewMock) CreateNewTranslationSentenceStructInstance(englishSentence string, gopherSentence string) models.TranslationSentence {
	return models.TranslationSentence{
		EnglishSentence: englishSentence,
		GopherSentence: gopherSentence,
	}
}

// CreateNewTranslationWelcomeStructInstance creates a new instance of model TranslationSentence
func (mock *NewMock) CreateNewTranslationWelcomeStructInstance(welcome string, routes map[string]string) models.TranslationWelcome {
	return models.TranslationWelcome{
		Welcome: welcome,
		Routes: routes,
	}
}

// TranslateEnglishWordToGopher translate english word to gopher
func (mock *NewMock) TranslateEnglishWordToGopher(englishWord string) string {
	translator := translation.CreateNewGopherTranslator()
	return translator.TranslateEnglishWordToGopher(englishWord)
}

// TranslateEnglishSentenceToGopher translate english sentence to gopher
func (mock *NewMock) TranslateEnglishSentenceToGopher(englishSentence string) string {
	translator := translation.CreateNewGopherTranslator()
	return translator.TranslateEnglishSentenceToGopher(englishSentence)
}