package main

import "fmt"

//	func mera_func(x, y int, z string) int {
//		return 0
//	}
//
//	type MyType struct {
//		a int
//	}

//func modify(t *MyType) {
//	t = &MyType{a: 100}
//}

//	func modifySlice(s *MyType) {
//		s.a = 100
//	}
type T1 struct {
	f1 int
	f2 string
}

type T2 struct {
	f1 int
	f6 string
	f4 map[string]int
}

func main() {
	//s := []int{1, 2, 3}
	//a := [...]int{1, 2, 3}
	//s1 := a[:2]
	//s_nil := a[:0]
	//fmt.Println(s_nil == nil)

	//s := &MyType{a: 10}
	//modify(s)
	//var i []int = nil  GOOD
	//var i [4]int = nil  BAD
	//var i int = nil  BAD
	//var maxInt uint64 = math.MaxUint64
	//myZero := uint64(0)
	//f := float64(maxInt)

	//maxInt = maxInt / myZero  :Panic
	//maxInt = maxInt / 0  :Compile time error
	//f = f / 0  :Works , return +Inf
	//fmt.Println(f)

	//var m map[string]int   // Compile time error
	//m = make(map[string]string)
	//
	//fmt.Println(m)

	//var m map[string]int
	//
	//fmt.Println(m["Kuchbhi"])
	//
	//m = nil
	//
	//fmt.Println(m, m["Kuchbhi"])

	t1 := T1{
		f1: 1,
		f2: "2",
	}

	t2 := T1{
		f1: 3,
		f2: "4",
	}

	t3 := T1(t2)

	fmt.Println(t3)

}
