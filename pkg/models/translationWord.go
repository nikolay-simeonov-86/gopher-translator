package models

// TranslationWord struct with the translation object
type TranslationWord struct {
	EnglishWord string `json:"english-word,omitempty"`
	GopherWord string `json:"gopher-word,omitempty"`
}

// ListTranslationWords contains a list of translations
type ListTranslationWords struct {
	TranslationWords []TranslationWord `json:"translation-words"`
}