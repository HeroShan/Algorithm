package matrix

import (
	"fmt"
	"testing"
)

func TestNewSplash(t *testing.T) {

	two := new(TwoDSource)
	two.Long = 5
	two.High = 5
	i:=0
	for{
		data := two.NewEmpty()
		Print2Matrix(data,5)
		fmt.Println("---------------------------------")
		i++
		if i >3{
			break
		}
		data = two.NewSplash(8)
		Print2Matrix(data,5)
	}
	//for _,v := range data{
	//	fmt.Printf("%+v\n",*v)
	//}
}
