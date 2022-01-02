package flesch_go

import "strings"

// Word is contiguous sequence of alphabetic characters.
// Whitespace defines word boundaries.
type Word struct {
	allRunes []rune
	Start    int
	End      int
}

func (w Word) Runes() []rune {

	return w.allRunes[w.Start : w.End+1]
}

func (w Word) String() string {
	var b strings.Builder
	for _, r := range w.Runes() {
		b.WriteRune(r)
	}
	return b.String()
}

// Syllables are considered to have been encountered whenever
// you detect a vowel at the start of a word or a vowel
// following a consonant in a word. A lone ‘e’ at the end
// of a word does not count as a syllable. Three letter words
// or less are always one syllable. One is the minimum.
func (w Word) Syllables() int {
	word := w.Runes()

	return syllablesFromRunes(word)
}

func syllablesFromRunes(runes []rune) int {
	var syllables int
	if len(runes) <= 3 {
		return 1
	}

	// vowel at start of runes
	if TypeOfRune(runes[0]) == RuneTypeVowel {
		syllables++
	}

	// ...or a vowel following a consonant in a runes
	for i, r := range runes {
		// Not relevant for first rune
		if i == 0 {
			continue
		}
		if TypeOfRune(r) == RuneTypeVowel && TypeOfRune(runes[i-1]) == RuneTypeConsonant {
			// do not count if last rune and a LONE 'e'
			if i == /* last: */ len(runes)-1 && /* is 'e': */ r == 'e' && /* lone: */ runes[i-1] != 'e' {
				continue
			}
			// otherwise, count it
			syllables++
		}
	}

	// all words are at least one syllable
	if syllables == 0 {
		syllables = 1
	}

	return syllables
}
