package handlers

import (
	"log"
	"regexp"
	"strings"
	"gopher-translator/pkg/translation"
	"gopher-translator/pkg/models"
	"encoding/json"
	"gopher-translator/pkg/storage"
	"net/http"
)

// TranstatorHandler Handles translate routes
type TranstatorHandler interface {
	Welcome() http.HandlerFunc
	TranslateWord(repository storage.Repository, translator translation.GopherTranslator) http.HandlerFunc
	TranslateSentence(repository storage.Repository, translator translation.GopherTranslator) http.HandlerFunc
	TranslateHistory(repository storage.Repository) http.HandlerFunc
}

// NewTranstatorHandler new TranstatorHandler
type NewTranstatorHandler struct {
	repository storage.Repository
	translator translation.GopherTranslator
}

// CreateNewTranslatorHandler creates a new NewTranstatorHandler
func CreateNewTranslatorHandler(repository storage.Repository, translator translation.GopherTranslator) *NewTranstatorHandler {
	return &NewTranstatorHandler{
		repository,
		translator,
	}
}

// Welcome Welcomes people to gopher translator
func (handler *NewTranstatorHandler) Welcome() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		welcomeRes := models.TranslationWelcome{
			Welcome:"Welcome to Gopher language Translator!", 
			Routes: map[string]string{
				"POST": "/word /sentence",
				"GET": "/history",
			},
		}
		json.NewEncoder(w).Encode(welcomeRes)
	}
}

// TranslateWord translates english words to gopher ones
func (handler *NewTranstatorHandler) TranslateWord() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		translationReq := models.TranslationWord{}
		err := json.NewDecoder(r.Body).Decode(&translationReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		errorString := handler.validateRequestBody(translationReq)
		if errorString != "" {
			json.NewEncoder(w).Encode(map[string]string{"error":errorString})
			return
		}

		translation, err := handler.repository.GetWordTranslation(translationReq.EnglishWord)
		if err == nil {
			json.NewEncoder(w).Encode(translation)
			return
		}

		newTranslation, err := handler.addNewWordTranslation(translationReq.EnglishWord);
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error":"Could not create new word translation!"})
			return
		}

		json.NewEncoder(w).Encode(newTranslation)
	}
}

// TranslateSentence TranslateSentence translates entire sentences from english to gopher
func (handler *NewTranstatorHandler) TranslateSentence() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		translationReq := models.TranslationSentence{}
		err := json.NewDecoder(r.Body).Decode(&translationReq)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		errorString := handler.validateRequestBody(translationReq)
		if errorString != "" {
			json.NewEncoder(w).Encode(map[string]string{"error":errorString})
			return
		}

		translation, err := handler.repository.GetSentenceTranslation(translationReq.EnglishSentence)
		if err == nil {
			json.NewEncoder(w).Encode(translation)
			return
		}

		newTranslation, err := handler.addNewSentenceTranslation(translationReq.EnglishSentence);
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"error":"Could not create new sentence translation!"})
			return
		}

		json.NewEncoder(w).Encode(newTranslation)
	}
}

// TranslateHistory returns a list with all translations from storage
func (handler *NewTranstatorHandler) TranslateHistory() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(handler.repository.GetAllTranslations())
	}
}

// addNewWordTranslation Adds new translation in the storage
func (handler *NewTranstatorHandler) addNewWordTranslation(key string) (translation models.TranslationWord, err error) {
	translation = models.TranslationWord{EnglishWord: key,}
	gopherWord := handler.translator.TranslateEnglishWordToGopher(key);
	translation.GopherWord = gopherWord
	err = handler.repository.AddWordTranslation(translation)
	if err != nil {
		return translation, err
	}

	// Remove to return full struct object
	translation.EnglishWord = ""
	return translation, nil 
}

// addNewSentenceTranslation Adds new translation in the storage
func (handler *NewTranstatorHandler) addNewSentenceTranslation(key string) (translation models.TranslationSentence, err error) {
	translation = models.TranslationSentence{EnglishSentence: key,}
	gopherSentence := handler.translator.TranslateEnglishSentenceToGopher(key);
	translation.GopherSentence = gopherSentence
	err = handler.repository.AddSentenceTranslation(translation)
	if err != nil {
		return translation, err
	}

	// Remove to return full struct object
	translation.EnglishSentence = ""
	return translation, nil 
}

// validateRequestBody validates the request body if it is empty or has apostrophes
func (handler *NewTranstatorHandler) validateRequestBody(reqObject interface{}) string {
	switch reqType := reqObject.(type) {
		case models.TranslationWord:
			if reqType.EnglishWord == "" {
				return `You must provide an english word to translate! Example JSON - {"english-word":"apple"}`
			}

			if strings.Contains(reqType.EnglishWord, " ") {
				return `You must provide an single english word! Example JSON - {"english-word":"apple"}`
			}

			if strings.Contains(reqType.EnglishWord, "'") {
				return `You must provide an english word without apostrophes! Example JSON - {"english-word":"apple"}`
			}
		case models.TranslationSentence:
			if reqType.EnglishSentence == "" {
				return `You must provide an english sentence to translate! Example JSON - {"english-sentence":"Apples are sweet!"}`
			}

			if strings.Contains(reqType.EnglishSentence, "'") {
				return `You must provide an english sentence without apostrophes! Example JSON - {"english-sentence":"Apples are sweet!"}`
			}

			if !strings.Contains(".?!", string(reqType.EnglishSentence[len(reqType.EnglishSentence)-1])) {
				return `You must provide an english sentence with a punctuation mark at the end! Example JSON - {"english-sentence":"Apples are sweet!"}`
			}

			reg, err := regexp.Compile("[ ]+[a-zA-Z]")
			if err != nil {
				log.Print(err)
			}

			if !reg.MatchString(reqType.EnglishSentence) {
				return `You must provide more than one english word! Example JSON - {"english-sentence":"Apples are sweet!"}`
			}
	}

	return ""
}