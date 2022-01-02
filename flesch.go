package flesch_go

import "errors"

func ParseString(text string) (Document, error) {
	report := Document{}

	runes := []rune(text)
	var currentRuneIndex int
	for {
		sentence, err := GetSentence(runes, currentRuneIndex)
		if err != nil {
			break
		}
		// get words in sentence
		for {
			var word Word
			word, err = GetWord(runes, currentRuneIndex, sentence.End)
			if err != nil {
				break
			}
			sentence.Words = append(sentence.Words, word)
			currentRuneIndex = word.End + 1
		}
		report.Sentences = append(report.Sentences, sentence)
		currentRuneIndex = sentence.End + 1
	}
	return report, nil
}

var NoMoreSentences = errors.New("no more sentences")
var NoMoreWords = errors.New("no more words")

func GetSentence(allRunes []rune, start int) (Sentence, error) {
	i := start
	sentence := Sentence{allRunes: allRunes}
	var sentenceStarted bool
	for {
		if i >= len(allRunes) {
			return sentence, NoMoreSentences
		}
		r := allRunes[i]
		// if the sentence hasn't started yet...
		if !sentenceStarted {
			// if the rune is a vowel or consonant, the sentence will have started here
			if TypeOfRune(r) == RuneTypeVowel || TypeOfRune(r) == RuneTypeConsonant {
				sentenceStarted = true
				sentence.Start = i
			}
		} else {
			if TypeOfRune(r) == RuneTypeSentenceStop {
				sentence.End = i
				return sentence, nil
			}
		}

		i++
	}
}

func GetWord(allRunes []rune, start int, stop int) (Word, error) {
	i := start
	word := Word{allRunes: allRunes}
	var wordStarted bool
	for {
		if i > stop {

			return word, NoMoreWords
		}
		r := allRunes[i]
		// if the word hasn't started yet...
		if !wordStarted {
			// if the rune is a vowel or consonant, the word will have started here
			if TypeOfRune(r) == RuneTypeVowel || TypeOfRune(r) == RuneTypeConsonant {
				wordStarted = true
				word.Start = i
			}
		} else {
			if TypeOfRune(r) == RuneTypeWhiteSpace ||
				TypeOfRune(r) == RuneTypeWordStop ||
				TypeOfRune(r) == RuneTypeSentenceStop {

				word.End = i - 1

				return word, nil
			}
		}

		i++
	}
}
