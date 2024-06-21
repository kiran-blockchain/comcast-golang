package main

import "testing"

func TestVerifyParallel(t *testing.T){
	t.Run("Test 3 in Parallel", func(t *testing.T){
		t.Parallel()
		result :=Verify(3)
		if result !="true"{
			t.Errorf("got:%s and expected %s",result,"true")
		}

	})
	t.Run("Test 4 in Parallel", func(t *testing.T){
		t.Parallel()
		result :=Verify(4)
		if result !="true"{
			t.Errorf("got:%s and expected %s",result,"4")
		}

	})
}