package models

// TranslationSentence struct with the translation sentence object
type TranslationSentence struct {
	EnglishSentence string `json:"english-sentence,omitempty"`
	GopherSentence string `json:"gopher-sentence,omitempty"`
}

// ListTranslationSentences contains a list of translation sentences
type ListTranslationSentences struct {
	TranslationSentences []TranslationSentence `json:"translation-sentences"`
}