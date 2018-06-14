package cstring

import (
	"fmt"
	"testing"
)

func Test_ToMoney(t *testing.T) {

	val := 1234.5678
	fmt.Println("ToMoney :", ToMoney(val))

	val2 := -12000000000000
	fmt.Println("ToMoney :", ToMoney(val2))

	var val3 int64
	val3 = 99999999999999
	fmt.Println("ToMoney :", ToMoney(val3))

	var val4 int32
	val4 = 999999999
	fmt.Println("ToMoney :", ToMoney(val4))

	//different types
	fmt.Println("ToMoney :", ToMoney([]byte{}))

	fmt.Println("ToMoney :", ToMoney(tVector{}))

}

type tVector struct {
	x float64
	y float64
	z float64
}
