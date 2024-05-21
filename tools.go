package toolkit

import "crypto/rand"

const randomStringSource = "abcdefghkijklmnopqrstuvwxyzABCDEFGHIJKLMNORQRSTUVWXYZ0123456789_+"

// Tools is the type used to instantiate this module.
// Any variable of this type will have access to all the methods with the receiver *Tools
type Tools struct{}

// RandomString returns a string of random characters of length , using randomStringSource
// as the source for the string
func (t *Tools) RandomString(length int) string {
	s, r := make([]rune, length), []rune(randomStringSource)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}
