package stemmer

import "testing"
import "bufio"
import "strings"
import "os"

func compare(t *testing.T, expected, actual interface{}, msg ...string) {
	if expected != actual {
		t.Errorf("[%v] -- value differs. Expected [%v], actual [%v]", msg, expected, actual)
	}
}


func TestConsonant(t *testing.T) {
	word := []byte("TOY")
	compare(t, true, Consonant(word, 0), "T")  //T
	compare(t, false, Consonant(word, 1), "O") //O
	compare(t, true, Consonant(word, 2), "Y")  //Y
	word = []byte("SYZYGY")
	compare(t, true, Consonant(word, 0), "S")  //S
	compare(t, false, Consonant(word, 1), "Y") //Y
	compare(t, true, Consonant(word, 2), "Z")  //Z
	compare(t, false, Consonant(word, 3), "Y") //Y
	compare(t, true, Consonant(word, 4), "G")  //G
	compare(t, false, Consonant(word, 5), "Y") //Y
	word = []byte("yoke")
	compare(t, true, Consonant(word, 0), "YOKE")
}

func TestMeasure(t *testing.T) {
	compare(t, 0, Measure([]byte("TR")))
	compare(t, 0, Measure([]byte("EE")))
	compare(t, 0, Measure([]byte("TREE")))
	compare(t, 0, Measure([]byte("Y")))
	compare(t, 0, Measure([]byte("BY")))
	compare(t, 1, Measure([]byte("TROUBLE")))
	compare(t, 1, Measure([]byte("OATS")))
	compare(t, 1, Measure([]byte("TREES")))
	compare(t, 1, Measure([]byte("IVY")))
	compare(t, 2, Measure([]byte("TROUBLES")))
	compare(t, 2, Measure([]byte("PRIVATE")))
	compare(t, 2, Measure([]byte("OATEN")))
	compare(t, 2, Measure([]byte("ORRERY")))
}

func Test1A(t *testing.T) {
	compare(t, "caress", string(one_a([]byte("caresses"))))
	compare(t, "poni", string(one_a([]byte("ponies"))))
	compare(t, "ti", string(one_a([]byte("ties"))))
	compare(t, "caress", string(one_a([]byte("caress"))))
	compare(t, "cat", string(one_a([]byte("cats"))))
}

func Test1B(t *testing.T) {
	compare(t, "feed", string(one_b([]byte("feed"))))
	compare(t, "agree", string(one_b([]byte("agreed"))))
	compare(t, "plaster", string(one_b([]byte("plastered"))))
	compare(t, "bled", string(one_b([]byte("bled"))))
	compare(t, "motor", string(one_b([]byte("motoring"))))
	compare(t, "sing", string(one_b([]byte("sing"))))
	compare(t, "motor", string(one_b([]byte("motoring"))))
	compare(t, "conflate", string(one_b([]byte("conflated"))))
	compare(t, "trouble", string(one_b([]byte("troubled"))))
	compare(t, "size", string(one_b([]byte("sized"))))
	compare(t, "hop", string(one_b([]byte("hopping"))))
	compare(t, "tan", string(one_b([]byte("tanned"))))
	compare(t, "fail", string(one_b([]byte("failing"))))
	compare(t, "file", string(one_b([]byte("filing"))))
}

func Test1C(t *testing.T) {
	compare(t, "sky", string(one_c([]byte("sky"))))
	compare(t, "happi", string(one_c([]byte("happy"))))

}

func Test2(t *testing.T) {
	compare(t, "relate", string(two([]byte("relational"))))
	compare(t, "condition", string(two([]byte("conditional"))))
	compare(t, "rational", string(two([]byte("rational"))))
	compare(t, "valence", string(two([]byte("valenci"))))
	compare(t, "hesitance", string(two([]byte("hesitanci"))))
	compare(t, "digitize", string(two([]byte("digitizer"))))
	compare(t, "conformable", string(two([]byte("conformabli"))))
	compare(t, "radical", string(two([]byte("radicalli"))))
	compare(t, "different", string(two([]byte("differentli"))))
	compare(t, "vile", string(two([]byte("vileli"))))
	compare(t, "analogous", string(two([]byte("analogousli"))))
	compare(t, "vietnamize", string(two([]byte("vietnamization"))))
	compare(t, "predicate", string(two([]byte("predication"))))
	compare(t, "operate", string(two([]byte("operator"))))
	compare(t, "feudal", string(two([]byte("feudalism"))))
	compare(t, "decisive", string(two([]byte("decisiveness"))))
	compare(t, "hopeful", string(two([]byte("hopefulness"))))
	compare(t, "callous", string(two([]byte("callousness"))))
	compare(t, "formal", string(two([]byte("formaliti"))))
	compare(t, "sensitive", string(two([]byte("sensitiviti"))))
	compare(t, "sensible", string(two([]byte("sensibiliti"))))
}

func Test3(t *testing.T) {
	compare(t, "triplic", string(three([]byte("triplicate"))))
	compare(t, "form", string(three([]byte("formative"))))
	compare(t, "formal", string(three([]byte("formalize"))))
	compare(t, "electric", string(three([]byte("electriciti"))))
	compare(t, "electric", string(three([]byte("electrical"))))
	compare(t, "hope", string(three([]byte("hopeful"))))
	compare(t, "good", string(three([]byte("goodness"))))
}

func Test4(t *testing.T) {
	compare(t, "reviv", string(four([]byte("revival"))))
	compare(t, "allow", string(four([]byte("allowance"))))
	compare(t, "infer", string(four([]byte("inference"))))
	compare(t, "airlin", string(four([]byte("airliner"))))
	compare(t, "gyroscop", string(four([]byte("gyroscopic"))))
	compare(t, "adjust", string(four([]byte("adjustable"))))
	compare(t, "defens", string(four([]byte("defensible"))))
	compare(t, "irrit", string(four([]byte("irritant"))))
	compare(t, "replac", string(four([]byte("replacement"))))
	compare(t, "adjust", string(four([]byte("adjustment"))))
	compare(t, "depend", string(four([]byte("dependent"))))
	compare(t, "adopt", string(four([]byte("adoption"))))
	compare(t, "homolog", string(four([]byte("homologou"))))
	compare(t, "commun", string(four([]byte("communism"))))
	compare(t, "activ", string(four([]byte("activate"))))
	compare(t, "angular", string(four([]byte("angulariti"))))
	compare(t, "homolog", string(four([]byte("homologous"))))
	compare(t, "effect", string(four([]byte("effective"))))
	compare(t, "bowdler", string(four([]byte("bowdlerize"))))
}

func Test5A(t *testing.T) {
	compare(t, "probat", string(five_a([]byte("probate"))))
	compare(t, "rate", string(five_a([]byte("rate"))))
	compare(t, "ceas", string(five_a([]byte("cease"))))
}

func Test5B(t *testing.T) {
	compare(t, "control", string(five_b([]byte("controll"))))
	compare(t, "roll", string(five_b([]byte("roll"))))
}

func TestVocal(t *testing.T) {
	f, err := os.Open("in.txt")
	if err != nil {
		panic(err)
	}
	in := bufio.NewReader(f)
	f, err = os.Open("out.txt")
	if err != nil {
		panic(err)
	}
	out := bufio.NewReader(f)
	for word, err := in.ReadSlice('\n'); err == nil; word, err = in.ReadSlice('\n') {
		stem, err := out.ReadSlice('\n')
		if err != nil {
			panic(err)
		}
		compare(t, strings.TrimSpace(string(stem)), string(Stem(word)), string(word))
	}
}
