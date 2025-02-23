package main

import "fmt"

type MyInterface interface {
	Add_int(x int) int
	Multiply(x float64) float64
}

type MyInt int
type MyFloat float64

// Implement MyInterface for MyInt type
func (i MyInt) Add_int(j int) int {
	return int(i) + j
}
func (i MyInt) Multiply(j float64) float64 {
	return float64(i) * j
}

// Implement MyInterface for MyFloat type
func (i MyFloat) Add_int(j int) int {
	return int(i) + j
}
func (i *MyFloat) Multiply(j float64) float64 {
	return float64(*i) * j
}

func main() {
	var inter MyInterface
	fmt.Println(inter)

	i := MyInt(1)
	f := MyFloat(5.5)

	
	var null_obj MyInt
	
	// Hold the value of i
	inter = i
	fmt.Printf("\n%v %T %d %f\n", inter, inter, inter.Add_int(10), inter.Multiply(10.0))
	
	// Hold the value of f
	// inter = f This will cause error
	inter = &f
	fmt.Printf("\n%v %T %d %f\n", inter, inter, inter.Add_int(10), inter.Multiply(10.0))
	
	fmt.Printf("\n The Underlying float value is %f\n",f)

	// other := 100
	// inter = other // this should give error as `int` doesn't implements my interface.

	// Interface with a nil unerlying concreet type
	inter = null_obj

	fmt.Printf("\n%v %T %d %f\n", inter, inter, inter.Add_int(10), inter.Multiply(10.0))
}
