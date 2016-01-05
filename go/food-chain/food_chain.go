package foodchain

import "fmt"

//Exercism variable.
const TestVersion = 1

type verse struct {
	animal string
	verse  string
}

var song = []verse{
	{"fly", "I don't know why she swallowed the fly. Perhaps she'll die."},
	{"spider", " that wriggled and jiggled and tickled inside her"},
	{"bird", "How absurd to swallow a bird!"},
	{"cat", "Imagine that, to swallow a cat!"},
	{"dog", "What a hog, to swallow a dog!"},
	{"goat", "Just opened her throat and swallowed a goat!"},
	{"cow", "I don't know how she swallowed a cow!"},
}

//Verse takes an int and return a string that contains the corrisponding verse.
func Verse(i int) string {
	// adjust 'i' to the array index
	var s string
	for ; i >= 0; i-- {
		v := song[i]
		s += fmt.Sprintf("I know an old lady who swallowed a %s", v.animal)
	}
	return s
}

//Verses takes two int and return a string that contains the corrisponding part.
func Verses(first, last int) string {
	return ""
}

//Song return a string that contains the whole song.
func Song() string {
	return Verses(1, 8)
}
