package fuzztesting

import (
	"testing"
)
func TestGuess(f *testing.F){
	f.Fuzz(func(t *testing.T,input int){
		Guess(input)
	})
}