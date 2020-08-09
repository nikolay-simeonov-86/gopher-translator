package translation

import (
	"log"
	"regexp"
	"strings"
)

// GopherTranslator GopherTranslator can translate
type GopherTranslator interface {
	TranslateEnglishWordToGopher(word string) string
	TranslateEnglishSentenceToGopher(sentence string) string
}

// NewGopherTranslator NewGopherTranslator is a new trnslator instance
type NewGopherTranslator struct {
	vowelsString string
}

// CreateNewGopherTranslator CreateNewGopherTranslator creates a new instance of NewGopherTranslator
func CreateNewGopherTranslator() *NewGopherTranslator {
	return &NewGopherTranslator{
		vowelsString: "aeiouAEIOU",
	}
}

// TranslateEnglishWordToGopher Translates english words to gopher alternatives
func (gopherTranslator *NewGopherTranslator) TranslateEnglishWordToGopher(word string) string {
	wordLength := len(word)
	word = removeNonAlphabeticCharactersFromWord(word)

	return translateWithLoops(word, wordLength, gopherTranslator.vowelsString)
}

// TranslateEnglishSentenceToGopher Translates english sentences to gopher alternatives
func (gopherTranslator *NewGopherTranslator) TranslateEnglishSentenceToGopher(sentence string) string {
	wordsArray := strings.Split(sentence, " ")
	firstWord := strings.ToLower(wordsArray[0])
	firstWord = translateWithLoops(firstWord, len(firstWord), gopherTranslator.vowelsString)
	firstWord = strings.Title(firstWord)
	translatedWordsArray := []string{firstWord}
	for _, word := range wordsArray[1:] {
		word = removeNonAlphabeticCharactersFromWord(word)
		wordLength := len(word)
		wordTranslated := translateWithLoops(word, wordLength, gopherTranslator.vowelsString)
		translatedWordsArray = append(translatedWordsArray, wordTranslated)
	}
log.Print(translatedWordsArray)
	return strings.Join(translatedWordsArray, " ") + string(sentence[len(sentence)-1])
}

// removeNonAlphabeticCharactersFromWord Removes non alphabetic characters from the provided string (word)
func removeNonAlphabeticCharactersFromWord(word string) string {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Print(err)
	}

	return reg.ReplaceAllString(word, "")
}

// translateWithLoops Translate using loops to catch any number of consonants at the beggining
func translateWithLoops(word string, wordLength int, vowelsString string) string {
	out:
	for index, value := range word {
		switch value {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
				word = ("g" + word)
				break out
			default:
				if ((wordLength >= 2) && 
					((word[index+1] == 'r') || 
					(word[index+1] == 'R'))) {
					word = ("ge" + word)
					break out
				} else if ((wordLength >= 3) && 
							(word[index+1] == 'q' || word[index+1] == 'Q') && 
							(word[index+2] == 'u' || word[index+2] == 'U')) {
					word = (word[3:] + word[:3] + "ogo")
					break out
				} else {
					for key, val := range word[1:] {
						if strings.Contains(vowelsString, string(val)) {
							word = (word[key+1:] + word[:key+1] + "ogo")
							break out
						}
					}
					break out
				}
		}
	}

	return word
}

// translateWithoutLoops Translate without loops but catches limited number of consonants from the beggining of words
func translateWithoutLoops(word string, wordLength int, vowelsString string) string {
	if strings.Contains(vowelsString, string(word[0])) {
		word = ("g" + word)
	} else {
		if wordLength > 1 {
			if ((word[0] == 'x' || word[0] == 'X') && (word[1] == 'r' || word[1] == 'R')) {
				word = ("ge" + word)
			} else {
				if wordLength > 2 {
					if !strings.Contains(vowelsString, string(word[1])) {
						if (!strings.Contains(vowelsString, string(word[2])) || (word[2] == 'u' || word[2] == 'U')) {
							word = (word[3:] + word[:3] + "ogo")
						} else {
							word = (word[2:] + word[:2] + "ogo")
						}
					} else {
						word = (word[2:] + word[:2] + "ogo")
					}
				} else {
					word = (word[1:] + word[:1] + "ogo")
				}
			}
		} else {
			word = ("ogo" + word)
		}
	}

	return word
}