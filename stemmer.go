package stemmer

import "fmt"

func ingore() {
	fmt.Sprintf("")
}

func Consonant(body []byte, offset int) bool {
	switch body[offset] {
	case 'A', 'E', 'I', 'O', 'U':
		return false
	case 'Y':
		return offset > 0 && !Consonant(body, offset-1)
	}
	return true
}

func Vowel(body []byte, offset int) bool {
	return !Consonant(body, offset)
}

const (
	vowel_state = iota
	consonant_state
)

func Meansure(body []byte) int {
	meansure := 0
	var state int
	if Vowel(body, 0) {
		state = vowel_state
	} else {
		state = consonant_state
	}
	for i := 0; i < len(body); i++ {
		if Vowel(body, i) && state == consonant_state {
			state = vowel_state
		} else if Consonant(body, i) && state == vowel_state {
			state = consonant_state
			meansure++
		}
	}
	return meansure
}

func Stem(word []byte) []byte {
	return []byte("")
}
