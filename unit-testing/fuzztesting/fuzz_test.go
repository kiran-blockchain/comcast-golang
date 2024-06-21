package fuzztesting

import (
	"testing"
	"unicode/utf8"
)

func FuzzReverseTest(f *testing.F){
	testcases :=[]string{"Hello Guys"," ","!12345"}
	for _,tc := range testcases{
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, orig string){
		rev :=Reverse(orig)
		doubleRev := Reverse(rev)
		if orig != doubleRev{
			t.Errorf("Before  %q, after %q",orig,doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev){
			t.Errorf("Reverse produced an invalid UTF-8 Sting %q",rev)
		}
	})
}