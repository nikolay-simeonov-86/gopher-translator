package storage

import (
	"strings"
	"log"
	"gopher-translator/pkg/models"
	"github.com/go-redis/redis"
)

// NewRedisStorage New redis client with custom functions
type NewRedisStorage struct {
	redisPrefixString string
	client *redis.Client
}

// Repository has methods to interact with storage (database or other)
type Repository interface {
	AddWordTranslation(translation models.TranslationWord) (err error)
	GetWordTranslation(key string) (res models.TranslationWord, err error)
	AddSentenceTranslation(translation models.TranslationSentence) (err error)
	GetSentenceTranslation(key string) (res models.TranslationSentence, err error)
	GetAllTranslations() (res models.TranslationHistoryList)
}

// CreateNewRedisStorage returns a new redis storage client
func CreateNewRedisStorage(address string, password string, dbNumber int) *NewRedisStorage {
	if address == "" {
		address = "localhost:6379"
	}
	// Create a new Redis client
	client := redis.NewClient(&redis.Options{
		Addr: address,
		Password: password,
		DB: dbNumber,
	})
	redisPrefixString := "gopher-translator:"
	return &NewRedisStorage{
		redisPrefixString: redisPrefixString,
		client: client,
	}
}

// AddWordTranslation Adds translation from english to gopher for a word
func (redisStorage *NewRedisStorage) AddWordTranslation(translation models.TranslationWord) (err error) {
	err = redisStorage.client.Set(redisStorage.redisPrefixString + translation.EnglishWord, translation.GopherWord, 0).Err()
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

// GetWordTranslation Adds translation from english to gopher for a word
func (redisStorage *NewRedisStorage) GetWordTranslation(key string) (res models.TranslationWord, err error) {
	translation, err := redisStorage.client.Get(redisStorage.redisPrefixString + key).Result()
	translationRes := models.TranslationWord{}
	if err != nil {
		log.Print(err)
		return translationRes, err
	}
	
	// Uncomment to return full struct object
	// translationRes.EnglishWord = key
	translationRes.GopherWord = translation

	return translationRes, nil
}

// AddSentenceTranslation Adds translation from english to gopher for a sentence
func (redisStorage *NewRedisStorage) AddSentenceTranslation(translation models.TranslationSentence) (err error) {
	err = redisStorage.client.Set(redisStorage.redisPrefixString + translation.EnglishSentence, translation.GopherSentence, 0).Err()
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

// GetSentenceTranslation Adds translation from english to gopher for a word
func (redisStorage *NewRedisStorage) GetSentenceTranslation(key string) (res models.TranslationSentence, err error) {
	translation, err := redisStorage.client.Get(redisStorage.redisPrefixString + key).Result()
	translationRes := models.TranslationSentence{}
	if err != nil {
		log.Print(err)
		return translationRes, err
	}
	
	// Uncomment to return full struct object
	// translationRes.EnglishSentence = key
	translationRes.GopherSentence = translation

	return translationRes, nil
}

// GetAllTranslations gets all translations from the storage
func (redisStorage *NewRedisStorage) GetAllTranslations() (res models.TranslationHistoryList) {
	translations := models.TranslationHistoryList{}
	translationsKeys, _, err := redisStorage.client.Scan(0, "", 0).Result()
	if err != nil {
		log.Print(err)
		translations.History = append(translations.History, map[string]string{})
		return translations
	}
	
	if len(translationsKeys) == 0 {
		translations.History = append(translations.History, map[string]string{})
		return translations
	}

	for _, translationKey := range translationsKeys {
		translationKey = strings.Replace(translationKey, redisStorage.redisPrefixString, "", 1)
		translationValue, err := redisStorage.GetWordTranslation(translationKey)
		if err != nil {
			log.Print(err)
		}
		translation := map[string]string{
			translationKey: translationValue.GopherWord,
		}
		translations.History = append(translations.History, translation)
	}

	return translations
}