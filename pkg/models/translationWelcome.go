package models

// TranslationWelcome struct with basic API informatoin
type TranslationWelcome struct {
	Welcome string `json:"welcome"`
	Routes map[string]string `json:"routes"`
}