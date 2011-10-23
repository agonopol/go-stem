package stemmer

func consonant(body []byte, offset int) bool {
	switch body[offset] {
	case 'A', 'E', 'I', 'O', 'U':
		return false
	case 'Y':
		if len(body)+1 < offset {
			return !consonant(body, offset+1)
		}
		return false
	}
	return true
}

func vowel(body []byte, offset int) bool {
	return !consonant(body, offset)
}

const (
	vowel_state = iota
	consonant_state
)

func meansure(body []byte) int {
	meansure := 0
	var state int
	if vowel(body, 0) {
		state = vowel_state
	} else {
		state = consonant_state
	}
	for i := 0; i < len(body); i++ {
		if vowel(body, i) && state == consonant_state {
			state = vowel_state
			meansure++
		} else if consonant(body, i) && state == vowel_state {
			state = consonant_state
			meansure++
		}
	}
	return meansure
}

func Stem(word []byte) []byte {
	return []byte("")
}
