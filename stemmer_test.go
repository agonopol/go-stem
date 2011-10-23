package stemmer

import "testing"

func compare(t *testing.T, expected, actual interface{}, msg ... string) {
	if expected != actual {
		t.Errorf("[%v] -- value differs. Expected [%v], actual [%v]",msg, expected, actual)
	}
}

func TestConsonant(t *testing.T) {
	word := []byte("TOY")
	compare(t, true, Consonant(word, 0),"T") //T
	compare(t, false, Consonant(word, 1),"O") //O
	compare(t, true, Consonant(word, 2),"Y") //Y
	word = []byte("SYZYGY")
	compare(t, true, Consonant(word, 0),"S") //S
	compare(t, false, Consonant(word, 1),"Y") //Y
	compare(t, true, Consonant(word, 2),"Z") //Z
	compare(t, false, Consonant(word, 3),"Y") //Y
	compare(t, true, Consonant(word, 4),"G") //G
	compare(t, false, Consonant(word, 5),"Y") //Y
	word = []byte("yoke")
	compare(t, true, Consonant(word,0), "YOKE")
}

func TestMeasure(t *testing.T) {
	compare(t, 0, Meansure([]byte("TR")))
	compare(t, 0, Meansure([]byte("EE")))
	compare(t, 0, Meansure([]byte("TREE")))
	compare(t, 0, Meansure([]byte("Y")))
	compare(t, 0, Meansure([]byte("BY")))
	compare(t, 1, Meansure([]byte("TROUBLE")))
	compare(t, 1, Meansure([]byte("OATS")))
	compare(t, 1, Meansure([]byte("TREES")))
	compare(t, 1, Meansure([]byte("IVY")))
	compare(t, 2, Meansure([]byte("TROUBLES")))
	compare(t, 2, Meansure([]byte("PRIVATE")))
	compare(t, 2, Meansure([]byte("OATEN")))
	compare(t, 2, Meansure([]byte("ORRERY")))
}