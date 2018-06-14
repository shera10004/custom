package cstring

import (
	"fmt"
	"log"
	"runtime/debug"
	"strconv"
)

// param : int , float32 , float64
func ToMoney(param interface{}) string {
	sv := ""
	//v, ok := param.(float64)

	switch param.(type) {

	case float32:
		sv = strconv.FormatFloat(param.(float64), 'f', -1, 32)
	case float64:
		sv = strconv.FormatFloat(param.(float64), 'f', -1, 64)
	//*/
	case int, int8, int16, int32, int64:
		sv = fmt.Sprintf("%v", param)

	default:
		ev := fmt.Sprintf("param is no digit --> %#+v", param)
		log.Println(ev)
		log.Println(string(debug.Stack()))
		return ev
	}

	bs := []byte(sv)
	size := len(bs)

	result := []byte{}
	div := 0
	dotIndex := -1

	for i, v := range bs {
		if v == '.' {
			dotIndex = i
		}
	} //for

	if dotIndex == -1 {
		dotIndex = size
	}

	fv := []byte{}
	div = 2
	for i := dotIndex - 1; i >= 0; i-- {
		fv = append(fv, bs[i])
		if div == 0 && i != 0 {
			fv = append(fv, ',')
			div = 2
		} else {
			div--
		}
	}

	//정수부
	for i := len(fv) - 1; i >= 0; i-- {
		result = append(result, fv[i])
	}

	//소수부
	for i := dotIndex; i < size; i++ {
		result = append(result, bs[i])
	}

	return string(result)
}
