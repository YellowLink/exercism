package hamming

import (
	"errors"
	"unicode/utf8"
)

func Distance(a, b string) (int, error) {
	distance := 0

	if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
		return 0, errors.New("Sequences must be the same length.")
	}
	
	for i := 0; i < utf8.RuneCountInString(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
