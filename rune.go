package flesch_go

type RuneType int

func TypeOfRune(r rune) RuneType {
	runeType, ok := runeLookup[r]
	if !ok {
		return RuneTypeOther
	}
	return runeType
}

var runeLookup = map[rune]RuneType{
	// Numbers
	'0': RuneTypeNumber,
	'1': RuneTypeNumber,
	'2': RuneTypeNumber,
	'3': RuneTypeNumber,
	'4': RuneTypeNumber,
	'5': RuneTypeNumber,
	'6': RuneTypeNumber,
	'7': RuneTypeNumber,
	'8': RuneTypeNumber,
	'9': RuneTypeNumber,

	// Lowercase
	'a': RuneTypeVowel,
	'b': RuneTypeConsonant,
	'c': RuneTypeConsonant,
	'd': RuneTypeConsonant,
	'e': RuneTypeVowel,
	'f': RuneTypeConsonant,
	'g': RuneTypeConsonant,
	'h': RuneTypeConsonant,
	'i': RuneTypeVowel,
	'j': RuneTypeConsonant,
	'k': RuneTypeConsonant,
	'l': RuneTypeConsonant,
	'm': RuneTypeConsonant,
	'n': RuneTypeConsonant,
	'o': RuneTypeVowel,
	'p': RuneTypeConsonant,
	'q': RuneTypeConsonant,
	'r': RuneTypeConsonant,
	's': RuneTypeConsonant,
	't': RuneTypeConsonant,
	'u': RuneTypeVowel,
	'v': RuneTypeConsonant,
	'w': RuneTypeConsonant,
	'x': RuneTypeConsonant,
	'y': RuneTypeConsonant,
	'z': RuneTypeConsonant,

	// Uppercase
	'A': RuneTypeVowel,
	'B': RuneTypeConsonant,
	'C': RuneTypeConsonant,
	'D': RuneTypeConsonant,
	'E': RuneTypeVowel,
	'F': RuneTypeConsonant,
	'G': RuneTypeConsonant,
	'H': RuneTypeConsonant,
	'I': RuneTypeVowel,
	'J': RuneTypeConsonant,
	'K': RuneTypeConsonant,
	'L': RuneTypeConsonant,
	'M': RuneTypeConsonant,
	'N': RuneTypeConsonant,
	'O': RuneTypeVowel,
	'P': RuneTypeConsonant,
	'Q': RuneTypeConsonant,
	'R': RuneTypeConsonant,
	'S': RuneTypeConsonant,
	'T': RuneTypeConsonant,
	'U': RuneTypeVowel,
	'V': RuneTypeConsonant,
	'W': RuneTypeConsonant,
	'X': RuneTypeConsonant,
	'Y': RuneTypeConsonant,
	'Z': RuneTypeConsonant,

	// Whitespace
	' ':  RuneTypeWhiteSpace,
	'\t': RuneTypeWhiteSpace,
	'\n': RuneTypeWhiteSpace,
	'\r': RuneTypeWhiteSpace,

	// Word Stop
	'”': RuneTypeWordStop,
	'"': RuneTypeWordStop,
	',': RuneTypeWordStop,
	')': RuneTypeWordStop,
	':': RuneTypeWordStop,
	'—': RuneTypeWordStop,

	// Sentence Stop
	'.': RuneTypeSentenceStop,
	';': RuneTypeSentenceStop,
	'!': RuneTypeSentenceStop,
	'?': RuneTypeSentenceStop,
}

const (
	RuneTypeSentenceStop RuneType = iota
	RuneTypeWhiteSpace
	RuneTypeWordStop
	RuneTypeVowel
	RuneTypeConsonant
	RuneTypeNumber
	RuneTypeOther
)
