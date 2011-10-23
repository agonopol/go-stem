package stemmer

import "testing"

func compare(t *testing.T, expected, actual interface{}, msg string) {
	if expected != actual {
		t.Errorf("[%v] -- value differs. Expected [%v], actual [%v]",msg, expected, actual)
	}
}

func TestConsonant(t *testing.T) {
	word := []byte("TOY")
	compare(t, true, consonant(word, 0),"T") //T
	compare(t, false, consonant(word, 1),"O") //O
	compare(t, true, consonant(word, 2),"Y") //Y
	word = []byte("SYZYGY")
	compare(t, true, consonant(word, 0),"S") //S
	compare(t, false, consonant(word, 1),"Y") //Y
	compare(t, true, consonant(word, 2),"Z") //Z
	compare(t, false, consonant(word, 3),"Y") //Y
	compare(t, true, consonant(word, 4),"G") //G
	compare(t, false, consonant(word, 5),"Y") //Y
	word = []byte("yoke")
	compare(t, true, consonant(word,0), "YOKE")
}