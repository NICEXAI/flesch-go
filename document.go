package flesch_go

import "strings"

type Document struct {
	Sentences []Sentence
}

func (d Document) Score() float32 {
	return 206.835 - (84.6 * d.avgSylPerWord()) - (1.015 * d.avgWordPerSen())
}

func (d Document) WordCount() int {
	var count int
	for _, sentence := range d.Sentences {
		count += len(sentence.Words)
	}

	return count
}

func (d Document) Words() []Word {
	var words []Word
	for _, sentence := range d.Sentences {
		words = append(words, sentence.Words...)
	}

	return words
}

func (d Document) UniqueWords() []Word {
	keys := make(map[string]bool)
	var list []Word
	for _, word := range d.Words() {
		// case invariant
		wordString := strings.ToUpper(word.String())
		if _, exists := keys[wordString]; !exists {
			keys[wordString] = true
			list = append(list, word)
		}
	}
	return list
}

func (d Document) Syllables() int {
	var count int
	for _, s := range d.Sentences {
		count += s.Syllables()
	}

	return count
}

func (d Document) SentencesCount() int {
	return len(d.Sentences)
}

func (d Document) Kincaid() float32 {
	score := .39*d.avgWordPerSen() + 11.8*d.avgSylPerWord() - 15.59

	return score
}

func (d Document) avgWordPerSen() float32 {
	words := float32(d.WordCount())
	sentences := float32(len(d.Sentences))

	return words / sentences
}

func (d Document) avgSylPerWord() float32 {
	syllables := float32(d.Syllables())
	words := float32(d.WordCount())

	return syllables / words
}

func (d Document) ReadableScore() string {
	var scoreMessage string
	score := d.Score()
	switch {
	case score <= 100 && score > 90:
		scoreMessage = "5th grade "
	case score <= 90 && score > 80:
		scoreMessage = "6th grade "
	case score <= 80 && score > 70:
		scoreMessage = "7th grade "
	case score <= 70 && score > 60:
		scoreMessage = "8th & 9th grade "
	case score <= 60 && score > 50:
		scoreMessage = "10th to 12th grade"
	case score <= 50 && score > 30:
		scoreMessage = "College"
	case score <= 30 && score >= 0:
		scoreMessage = "College graduate"
	default:
		scoreMessage = "(invalid)"
	}

	return scoreMessage
}
