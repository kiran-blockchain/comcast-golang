package main
import "testing"
func BenchmarkVerify(b *testing.B){
	for i:=0;i<b.N;i++{
		Verify(i)
	}
}