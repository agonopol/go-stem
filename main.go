package main


import "stemmer"

func main() {
  println(string(stemmer.Stem([]byte("caresses"))))
}
