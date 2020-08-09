package models

// {“history”:[{“apple”:”gapple”},{“my”:”ymogo”},….]}

// TranslationHistoryList struct with list of translation items
type TranslationHistoryList struct {
	History []map[string]string `json:"history"`
}