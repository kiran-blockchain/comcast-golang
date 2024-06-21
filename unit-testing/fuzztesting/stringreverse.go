package fuzztesting

import "fmt"

func main (){
	input := "kiran"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("Original %q \n",input)
	fmt.Printf("Reversed %q \n",rev)
	fmt.Printf("Reversed %q \n",doubleRev)
}
func Reverse (s string) string {
	b := []byte(s)
	for i, j  :=0, len(b)-1; i<len(b)/2; i,j=i+1,j-1{
		b[i],b[j]= b[j],b[i]
	}
	return string(b)
}