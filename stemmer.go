package stemmer

import "regexp"

var one_a = regexp.MustCompile("(.*)?(ss|i)es$")

func Stem(word []byte) []byte {
  if (len(word) < 3) {
    return word 
  } 
  if (word[0] == 'Y') {
    word[0] = 'y'
  }
  if (one_a.Match(word)) {
    match := one_a.FindSubmatch(word)
    word = append(match[1],match[2]...)
  }
  return word
}
