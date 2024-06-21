package main
import "testing"

func TestVerify(t *testing.T){
	result :=Verify(6)
	if result !="true"{
		t.Errorf("Result was incorrect got:%s and expected %s",result,"true")
	}
}