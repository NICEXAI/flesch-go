package flesch_go

import "strings"

// Sentence is encountered whenever you find a word that
// ends in a specific punctuation symbol: a period, question
// mark, or exclamation point.
type Sentence struct {
	allRunes []rune
	Start    int
	End      int
	Words    []Word
}

func (s Sentence) Runes() []rune {
	return s.allRunes[s.Start : s.End+1]
}

func (s Sentence) String() string {
	var b strings.Builder
	for _, r := range s.Runes() {
		b.WriteRune(r)
	}
	return b.String()
}

func (s Sentence) Syllables() int {
	var count int
	for _, w := range s.Words {
		count += w.Syllables()
	}

	return count
}
