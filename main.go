package main

import "stemmer"
import "bufio"
import "os"

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	
	for word, err := in.ReadSlice(' '); err == nil; word, err = in.ReadSlice(' ') {
		out.Write(stemmer.Stem(word))
		out.WriteString("\n")
	}
}
