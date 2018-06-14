package csort

/*
Sort(func() int {
		return len(s)
	},
		func(x, y int) bool {
			return s[x].name < s[y].name
		},
		func(x, y int) {
			s[x], s[y] = s[y], s[x]
		},
	)
*/

import (
	"sort"
)

type cLen func() int
type cLess func(x, y int) bool //각 상황별 정렬 함수를 저장할 타입
type cSwap func(x, y int)

type stDataSorter struct {
	size    cLen
	compair cLess //func(s1, s2 *Student) bool
	change  cSwap
}

//Sort(len func()int , less func(x,y int)bool , swap func(x,y int)){}
func Sort(len cLen, less cLess, swap cSwap) {
	sorter := &stDataSorter{len, less, swap}
	sort.Sort(sorter)
}

func (s *stDataSorter) Len() int {
	return s.size()
}
func (s *stDataSorter) Less(x, y int) bool {
	return s.compair(x, y)
}
func (s *stDataSorter) Swap(x, y int) {
	s.change(x, y)
}
