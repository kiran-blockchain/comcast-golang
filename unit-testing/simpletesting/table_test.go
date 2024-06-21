package main

import "testing"

func TestVerifyTableDriven(t *testing.T){

	var tests =[]struct{
		name string
		input int
		result string
	}{
		{"9 shoudl be verifiable", 9, "true"},
		{"7 shouldnot be verifiable", 7, "7"},
		{"0 shoudl be verifiable", 0, "true"},
	}
	for _,tt := range tests{
		t.Run(tt.name, func(t *testing.T){
			ans := Verify(tt.input)
			if ans !=tt.result{
				t.Errorf("Result was incorrect got:%s and expected %s",ans,tt.result)
			}
		})
	}
}