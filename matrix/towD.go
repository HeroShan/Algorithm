package matrix

import (
	"fmt"
	"math/rand"
	"time"
)


type TwoDSource struct {
	High int64
	Long int64
}

func Print2Matrix(m []*Source,n int){
	data := make([]interface{},0)
	for _,v := range m{
		data = append(data,v.Val)
		if len(data)== n{
			fmt.Println(data)
			data = nil
		}
	}
	if len(data) != 0{
		fmt.Println(data)
	}
}

//空矩阵
func (t *TwoDSource) NewEmpty() []*Source {
	s := make([]*Source, 0)
	for h := int64(0); h < t.High; h++ {
		for l := int64(0); l < t.Long; l++ {
			s = append(s, &Source{
				X:   0,
				Y:   0,
				Val: nil,
			})
		}
	}
	return s
}

//散点矩阵
func (t *TwoDSource) NewSplash(spNum int) []*Source {

	r := make(map[int64]bool)
	rr := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		if len(r)==spNum{
			break
		}
		r[rr.Int63n(t.High*t.Long)] = true
	}
	s := make([]*Source, 0)
	for h := int64(0); h < t.High; h++ {
		for l := int64(0); l < t.Long; l++ {
			if r[h*l] {
				s = append(s, &Source{
					X:   l,
					Y:   h,
					Val: "^",
				})

			} else {
				s = append(s, &Source{
					X:   l,
					Y:   h,
					Val: "*",
				})
			}

		}
	}
	return s
}
