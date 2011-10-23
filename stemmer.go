package stemmer

import "fmt"
import "bytes"

func ingore() {
	fmt.Sprintf("")
}

func Consonant(body []byte, offset int) bool {
	switch body[offset] {
	case 'A', 'E', 'I', 'O', 'U', 'a', 'e', 'i', 'o', 'u':
		return false
	case 'Y', 'y':
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

func HasVowel(body []byte) bool {
	for i := 0; i < len(body); i++ {
		if Vowel(body, i) {
			return true
		}
	}
	return false
}

func one_a(body []byte) []byte {
	if bytes.HasSuffix(body, []byte("sses")) || bytes.HasSuffix(body, []byte("ies")) {
		return body[:len(body)-2]
	} else if bytes.HasSuffix(body, []byte("ss")) {
		return body
	} else if bytes.HasSuffix(body, []byte("s")) {
		return body[:len(body)-1]
	}
	return body
}

func star_o(body []byte) bool {
	size := len(body) - 1
	if Consonant(body, size-2) && Vowel(body, size-1) && Consonant(body, size) {
		return body[size] != 'w' && body[size] != 'x' && body[size] != 'y'
	}
	return false
}
func one_b_a(body []byte) []byte {
	size := len(body)
	if bytes.HasSuffix(body, []byte("at")) {
		return append(body, 'e')
	} else if bytes.HasSuffix(body, []byte("bl")) {
		return append(body, 'e')
	} else if bytes.HasSuffix(body, []byte("iz")) {
		return append(body, 'e')
	} else if Consonant(body, size-1) && Consonant(body, size-2) && body[size-1] == body[size-2] {
		if body[size-1] != 'l' || body[size-1] != 's' || body[size-1] != 'z' {
			return body[:size-1]
		}
	} else if star_o(body) && Meansure(body) == 1 {
		return append(body, 'e')
	}
	return body
}

func one_b(body []byte) []byte {
	if bytes.HasSuffix(body, []byte("eed")) {
		if Meansure(body[:len(body)-1]) > 0 {
			return body[:len(body)-1]
		}
		return body
	} else if bytes.HasSuffix(body, []byte("ed")) && HasVowel(body[:len(body)-2]) {
		return one_b_a(body[:len(body)-2])
	} else if bytes.HasSuffix(body, []byte("ing")) && HasVowel(body[:len(body)-3]) {
		return one_b_a(body[:len(body)-3])
	}
	return body
}

func one_c(body []byte) []byte {
	if bytes.HasSuffix(body, []byte("y")) && HasVowel(body[:len(body)-1]) {
		body[len(body)-1] = 'i'
		return body
	}
	return body
}

func two(body []byte) []byte {
	if bytes.HasSuffix(body, []byte("ational")) && Meansure(body[:len(body)-7]) > 0 {
		return append(body[:len(body)-7], []byte("ate")...)
	} else if bytes.HasSuffix(body, []byte("tional")) && Meansure(body[:len(body)-6]) > 0 {
		return append(body[:len(body)-6], []byte("tion")...)
	} else if bytes.HasSuffix(body, []byte("enci")) && Meansure(body[:len(body)-4]) > 0 {
		return append(body[:len(body)-4], []byte("ence")...)
	} else if bytes.HasSuffix(body, []byte("anci")) && Meansure(body[:len(body)-4]) > 0 {
		return append(body[:len(body)-4], []byte("ance")...)
	} else if bytes.HasSuffix(body, []byte("izer")) && Meansure(body[:len(body)-4]) > 0 {
		return append(body[:len(body)-4], []byte("ize")...)
	} else if bytes.HasSuffix(body, []byte("abli")) && Meansure(body[:len(body)-4]) > 0 {
		return append(body[:len(body)-4], []byte("able")...)
	} else if bytes.HasSuffix(body, []byte("alli")) && Meansure(body[:len(body)-4]) > 0 {
		return append(body[:len(body)-4], []byte("al")...)
	} else if bytes.HasSuffix(body, []byte("entli")) && Meansure(body[:len(body)-5]) > 0 {
		return append(body[:len(body)-5], []byte("ent")...)
	} else if bytes.HasSuffix(body, []byte("eli")) && Meansure(body[:len(body)-3]) > 0 {
		return append(body[:len(body)-3], []byte("e")...)
	} else if bytes.HasSuffix(body, []byte("ousli")) && Meansure(body[:len(body)-5]) > 0 {
		return append(body[:len(body)-5], []byte("ous")...)
	} else if bytes.HasSuffix(body, []byte("ization")) && Meansure(body[:len(body)-7]) > 0 {
		return append(body[:len(body)-7], []byte("ize")...)
	} else if bytes.HasSuffix(body, []byte("ation")) && Meansure(body[:len(body)-5]) > 0 {
		return append(body[:len(body)-5], []byte("ate")...)
	} else if bytes.HasSuffix(body, []byte("ator")) && Meansure(body[:len(body)-4]) > 0 {
		return append(body[:len(body)-4], []byte("ate")...)
	} else if bytes.HasSuffix(body, []byte("alism")) && Meansure(body[:len(body)-5]) > 0 {
		return append(body[:len(body)-5], []byte("al")...)
	} else if bytes.HasSuffix(body, []byte("iveness")) && Meansure(body[:len(body)-7]) > 0 {
		return append(body[:len(body)-7], []byte("ive")...)
	} else if bytes.HasSuffix(body, []byte("fulness")) && Meansure(body[:len(body)-7]) > 0 {
		return append(body[:len(body)-7], []byte("ful")...)
	} else if bytes.HasSuffix(body, []byte("ousness")) && Meansure(body[:len(body)-7]) > 0 {
		return append(body[:len(body)-7], []byte("ous")...)
	} else if bytes.HasSuffix(body, []byte("aliti")) && Meansure(body[:len(body)-5]) > 0 {
		return append(body[:len(body)-5], []byte("al")...)
	} else if bytes.HasSuffix(body, []byte("iviti")) && Meansure(body[:len(body)-5]) > 0 {
		return append(body[:len(body)-5], []byte("ive")...)
	} else if bytes.HasSuffix(body, []byte("biliti")) && Meansure(body[:len(body)-6]) > 0 {
		return append(body[:len(body)-6], []byte("ble")...)
	}
	return body
}

func three(body []byte) []byte {
	if bytes.HasSuffix(body, []byte("icate")) && Meansure(body[:len(body)-5]) > 0 {
		return append(body[:len(body)-5], []byte("ic")...)
	} else if bytes.HasSuffix(body, []byte("ative")) && Meansure(body[:len(body)-5]) > 0 {
		return body[:len(body)-5]
	} else if bytes.HasSuffix(body, []byte("alize")) && Meansure(body[:len(body)-5]) > 0 {
		return append(body[:len(body)-5], []byte("al")...)
	} else if bytes.HasSuffix(body, []byte("iciti")) && Meansure(body[:len(body)-5]) > 0 {
		return append(body[:len(body)-5], []byte("ic")...)
	} else if bytes.HasSuffix(body, []byte("ical")) && Meansure(body[:len(body)-4]) > 0 {
		return append(body[:len(body)-4], []byte("ic")...)
	} else if bytes.HasSuffix(body, []byte("ful")) && Meansure(body[:len(body)-3]) > 0 {
		return body[:len(body)-3]
	} else if bytes.HasSuffix(body, []byte("ness")) && Meansure(body[:len(body)-4]) > 0 {
		return body[:len(body)-4]
	}
	return body
}

func four(body []byte) []byte {
	if bytes.HasSuffix(body, []byte("al")) && Meansure(body[:len(body)-2]) > 1 {
		return body[:len(body)-2]
	} else if bytes.HasSuffix(body, []byte("ance")) && Meansure(body[:len(body)-4]) > 1 {
		return body[:len(body)-4]
	} else if bytes.HasSuffix(body, []byte("ence")) && Meansure(body[:len(body)-4]) > 1 {
		return body[:len(body)-4]
	} else if bytes.HasSuffix(body, []byte("er")) && Meansure(body[:len(body)-2]) > 1 {
		return body[:len(body)-2]
	} else if bytes.HasSuffix(body, []byte("ic")) && Meansure(body[:len(body)-2]) > 1 {
		return body[:len(body)-2]
	} else if bytes.HasSuffix(body, []byte("able")) && Meansure(body[:len(body)-4]) > 1 {
		return body[:len(body)-4]
	} else if bytes.HasSuffix(body, []byte("ible")) && Meansure(body[:len(body)-4]) > 1 {
		return body[:len(body)-4]
	} else if bytes.HasSuffix(body, []byte("ant")) && Meansure(body[:len(body)-3]) > 1 {
		return body[:len(body)-3]
	} else if bytes.HasSuffix(body, []byte("ement")) && Meansure(body[:len(body)-5]) > 1 {
		return body[:len(body)-5]
	} else if bytes.HasSuffix(body, []byte("ment")) && Meansure(body[:len(body)-4]) > 1 {
		return body[:len(body)-4]
	} else if bytes.HasSuffix(body, []byte("ent")) && Meansure(body[:len(body)-3]) > 1 {
		return body[:len(body)-3]
	} else if bytes.HasSuffix(body, []byte("ion")) && Meansure(body[:len(body)-3]) > 1 {
		if len(body) > 4 && (body[len(body)-4] == 's' || body[len(body)-4] == 't') {
			return body[:len(body)-3]
		}
	} else if bytes.HasSuffix(body, []byte("ou")) && Meansure(body[:len(body)-2]) > 1 {
		return body[:len(body)-2]
	} else if bytes.HasSuffix(body, []byte("ism")) && Meansure(body[:len(body)-3]) > 1 {
		return body[:len(body)-3]
	} else if bytes.HasSuffix(body, []byte("ate")) && Meansure(body[:len(body)-3]) > 1 {
		return body[:len(body)-3]
	} else if bytes.HasSuffix(body, []byte("iti")) && Meansure(body[:len(body)-3]) > 1 {
		return body[:len(body)-3]
	} else if bytes.HasSuffix(body, []byte("ous")) && Meansure(body[:len(body)-3]) > 1 {
		return body[:len(body)-3]
	} else if bytes.HasSuffix(body, []byte("ive")) && Meansure(body[:len(body)-3]) > 1 {
		return body[:len(body)-3]
	} else if bytes.HasSuffix(body, []byte("ize")) && Meansure(body[:len(body)-3]) > 1 {
		return body[:len(body)-3]
	}
	return body
}

func five_a(body []byte) []byte {
	if bytes.HasSuffix(body, []byte("e")) && Meansure(body[:len(body)-1]) > 1 {
		return body[:len(body)-1]
	} else if bytes.HasSuffix(body, []byte("e"))  && Meansure(body[:len(body)-1]) == 1 && !star_o(body[:len(body)-1]) {
		return body[:len(body)-1]
	}
	return body
}

func five_b(body []byte) []byte {
	size := len(body)
	if Meansure(body) > 1 && Consonant(body, size-1) && Consonant(body, size-2) && body[size-1] == body[size-2] && body[size-1] == 'l' {
		return body[:len(body)-1]
	}
	return body
}

func Stem(body []byte) []byte {
	return five_b(five_a(four(three(two(one_c(one_b(one_a(bytes.TrimSpace(bytes.ToLower(body))))))))))
}
