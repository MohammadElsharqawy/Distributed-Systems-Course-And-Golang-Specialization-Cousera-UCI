package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sec/data"
	"strconv"
	"strings"
)

type grades int

const (
	A grades = iota
	B
	C
	D
)

func foo(x *[3]int) {
	// pass by reference
	(*x)[0] = (*x)[0] + 1
}

func foo1(sli []int) {
	sli[0] = sli[0] + 1
}

func add(x, y int) (z, t int) {
	z = x + y
	return
}

var i, j = "hi", 6

func gooo(afunc func(int) int, val int) int { //functions as arguments
	return afunc(val)
}

func applyIt(afunc func(int) int, val int) int {
	return afunc(val)
}

func inc(x int) int { return x + 1 }
func dec(x int) int { return x - 1 }

// o_x, o_y originx, orginy
func MakeDistFunction(o_x, o_y float64) func(float64, float64) float64 { // return a func which does compute the dis from the orig
	// now i can compute the distance from any origin i want

	fn := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-o_x, 2) + math.Pow(y-o_y, 2))
	}

	return fn
}

// Variadic
// variadic ---->>> pass variable argument number
// like c++ use ellipsis ... python just use * (as far as i remember)

func getMax(vals ...int) (retMax int) {
	retMax = -1
	for _, val := range vals {
		if val > retMax {
			retMax = val
		}
	}
	return // returns retMax implictly
}

func fA() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

type myint int

func (mi myint) Double() myint { // type should be defined in the same package as its method
	return mi * 2
}

func main() {
	fmt.Println(i, j)
	fmt.Print(math.Sqrt(9), " ")

	XX := 5
	if XX > 5 {
		fmt.Print("yuuup\n")
	} else if XX < 5 {
		fmt.Print("NOOOOOO\n")
	} else {
		fmt.Print("Yeeeeeees\n")
	}

	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	/*
		fmt.Printf("Enter the number of apples ")
		var apple, ioj int
		num, err := fmt.Scan(&apple, &ioj)

		fmt.Println(apple, ioj, num, err)
	*/

	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr1)

	for idx, v := range arr1 {
		fmt.Println(idx, v)
	}

	i, _ := strconv.Atoi("10")
	y := i * 2
	fmt.Println(y)

	s := strings.Replace("ianianian", "ni", "in", 2)
	fmt.Println(s)

	var xtemp int
	x1 := 0
	x2 := 1
	for x := 0; x < 5; x++ {
		xtemp = x2
		x2 = x2 + x1
		x1 = xtemp
	}
	fmt.Println(x2)

	x := 7
	switch {
	case x > 3:
		fmt.Printf("1\n")
	case x > 5:
		fmt.Printf("2")
	case x == 7:
		fmt.Printf("3")
	default:
		fmt.Printf("4")
	}

	sli := make([]int, 0, 3)
	sli = append(sli, 100)
	fmt.Print(sli, len(sli), cap(sli))
	fmt.Println()

	//hashMap

	var idMap map[string]int // nil should be allocated with make function

	idMap2 := make(map[string]int) // allocated

	idMap3 := map[string]int{
		"mo":    1,
		"ahmed": 1}
	m := map[string]int{} // used a map literal, initialize empty map

	idMap2["hi"] = 10
	//idMap["hi"] = 10 error idmap is nil
	idMap = make(map[string]int)

	idMap["hello"] = 10

	val := idMap["hello"]
	fmt.Println(val)
	delete(idMap3, "ahmed")

	fmt.Println(idMap, idMap2, idMap3, m)

	fmt.Println(idMap3["m0o"], idMap3["mo"]) // returns 0 if not exist

	id, p := idMap3["mo"] //id takes the value, p bool if the key is present or not

	fmt.Println(id, p, len(idMap3))

	fmt.Println(&val)

	for key, val := range idMap3 {
		fmt.Println(key, val)
		fmt.Println(&val)
	}

	//structs

	type Person struct {
		name    string
		address string
		phone   string
	}

	var p1 *Person = new(Person) //pointer
	*p1 = Person{name: "mo", address: "tanta", phone: "1234"}

	p2 := new(Person) // initialized fileds with zero, strings with empty strings (POINTER)

	fmt.Println(*p1, p2)

	// with struct literal
	p3 := Person{ // Not Pointer
		name: "mo", address: "tanta", phone: "1234"}

	fmt.Printf("%T %T %T\n", p3, p1, p2)

	xe := [...]int{4, 8, 5}
	ye := xe[0:2]
	ze := xe[1:3]
	ye[0] = 1
	ze[1] = 3
	fmt.Print(xe, "\n")

	xw := [...]int{1, 2, 3, 4, 5}
	yw := xw[0:2]
	zw := xw[1:4]
	fmt.Print(len(yw), cap(yw), len(zw), cap(zw), "\n")

	xt := map[string]int{
		"ian": 1, "harris": 2}

	for i, j := range xt {
		if i == "harris" {
			fmt.Print(i, j)
		}
	}
	fmt.Println()

	type P struct {
		x string
		y int
	}
	b := P{"x", -1}
	aw := [...]P{P{"a", 10},
		P{"b", 2},
		{"c", 3}} // we can remove P, like the last one here
	for _, z := range aw {
		if z.y > b.y {
			b = z
		}
	}
	fmt.Println(b.x)

	//Json

	//generate json from object

	p12 := Person{name: "mjo", address: "tanta", phone: "1234"}

	barr, err := json.Marshal(p12) // byte arr and err if no errors return nil

	if err != nil {
		fmt.Print(err, "hi err")
	} else {
		fmt.Println(string(barr))
	}

	type Employee struct {
		Name    string
		Age     int
		Address string
	}
	emp := Employee{Name: "George Smith", Age: 30, Address: "Newyork, USA"}
	empData, err := json.Marshal(emp)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(empData))

		fmt.Printf("%T \n", empData)
	}
	//////////////////////////////

	//p5 := Person{
	//	name: "hi", address: "uuu", phone: "cgd"}
	jsondata := `{"Name":"mohamed", "Address":"tanta", "Age":15}`
	fmt.Printf("%T \n", jsondata)
	dataBytes := []byte(jsondata)

	var p5 Employee
	err1 := json.Unmarshal(dataBytes, &p5) // if no error will return nil, (p5 & barr []byte must fit)

	if err1 != nil {
		fmt.Print(err1, "hi err")
	} else {
		fmt.Println(p5) // attributes in struct should start with capital letter, print the values
		fmt.Printf("%T \n", p5)
	}

	type Response struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Address string `json:"address"`
	}

	empJsonData := `{"name":"George Smith","age":30,"address":"Newyork, USA"}` // here first char not capital bcz of the struct `json:"name"`
	empBytes := []byte(empJsonData)
	var emp2 Response
	json.Unmarshal(empBytes, &emp2)
	fmt.Println(emp2)
	fmt.Println(emp2.Name)
	fmt.Println(emp2.Age)
	fmt.Println(emp2.Address)

	//files
	// file access is linear, not random access
	{
		data, _ := ioutil.ReadFile("test.txt")
		fmt.Println(string(data))
	}
	// data is []byte, readfile (explicit open/close are not needed)
	// large files cause problems

	dat := "Hello, World!"

	err2 := ioutil.WriteFile("test1.txt", []byte(dat), 0777)
	// 0777 unix-style permisssion
	fmt.Println(err2)

	//os.Open() return file descriptor (file construct)
	//os.Close()
	//os.Read() reads from a file into []byte
	//os.Write() writes a []byte into a file

	f, err := os.Open("test.txt")
	barr1 := make([]byte, 10) // length 10
	nb, err := f.Read(barr1)  // return the number of byted
	f.Close()
	fmt.Println(nb, string(barr1))

	f2, err := os.Create("out.txt")
	barr3 := []byte{1, 2, 3}
	nb2, err := f2.Write(barr3)

	nb3, err := f2.WriteString("hi")

	fmt.Println(nb2, nb3)

	aa := [3]int{1, 2, 3}

	foo(&aa) // messy and unnecessary, use slices instead
	// slice is a structure, first contains the first pointer of the array, length, capacity
	// when we pass a slice we pass by value but its value is contains a pinter so it is like coping a pointer
	fmt.Println(aa)

	bb := []int{2, 2, 3} // is a slice
	foo1(bb)             // slice is a structure, first contains the first pointer of the array, length, capacity
	// when we pass a slice we pass by value but its value is contains a pinter so it is like coping a pointer
	fmt.Println(bb)

	/////////////////////////////////////////////////////////////////////////////////////////////
	var funcVar func(int) int
	funcVar = inc // inc is a function without ()
	fmt.Println(funcVar(1))

	fmt.Println(applyIt(inc, 2))
	fmt.Println(applyIt(dec, 2))

	//anonymous functions
	v := applyIt(func(x int) int { return x + 1 }, 5)
	fmt.Println(v)

	dist1 := MakeDistFunction(0, 0) // origin 0,0
	dist2 := MakeDistFunction(2, 2) // origin 2,2

	fmt.Println(dist1(1, 1))
	fmt.Println(dist2(2, 2))

	//closure === function + environment
	// when we pass fn to another we also pass its environment

	// Variadic and Deferred

	// variadic ---->>> pass variable argument number
	mx := getMax(1, 2, 3, 4, 6, 8, 6, 5, 4, 5, 10, 100)
	fmt.Println(mx)

	sli5 := []int{1, 50, 900, 10}
	mx2 := getMax(sli5...)
	fmt.Println(mx2)

	//defer function call
	//used for clean up activity
	// called when the surrounding functions complete
	//eg. closes all the files at the end

	//defer fmt.Println("Bye!") // last called function in the main
	fmt.Println("Hello!")

	// keep in your mind that the arguments dont get deferred, they executed immediately, but the call is deferred.

	//i2 := 1
	//defer fmt.Println(i2 + 1)
	fmt.Println("Hello!")

	fB := fA()
	fmt.Print(fB())
	fmt.Println(fB())

	vv := myint(5)
	fmt.Println(vv)
	vv.Double() // it is like passing vv as an argument, and it is call by value
	fmt.Println(vv)
	{
		var p1 data.Point

		p1.InitMe(5.0, 10.0)
		// InitMe has a reciever of type pointer to a point struct
		// no need to reference either --->> (p1).Initme(), we include p1 itself not a reference --->> compiler recognize it

		fmt.Println(p1.GetX(), p1.GetY())
		p1.Scale(5)
		p1.PrintMe()
		//fmt.Println(data.x)
	}

	// type specifies an interface, if type defines all methods satisfied in the interface.

	//var s1 Speaker
	//d1 := &Dog{"brian"}
	//s1 = d1
	//s1.Speak()

	// panic, b3daha mafesh code byshtghal.
	//panic("PAAAAAAAAANNNIIIIC")
	// we can use fmt.println(recover())
	/////////////////////////////////////////////////////////////////////////

	// concrete types and interface types
	// concrete types are a regular types, specify the exact representation of the data and
	// the methods(methods that are used in the type of the reciever type)
	// so they are fully specified and has a complete implementation of the methods.
	// any method that uses these type as a reciever type

	// concrete type -->> you will have a bunch of data that are associated with the type
	// interface typr -->> just sprcifies some method signatures, so no data is specified, just the methods.
	// even the method, the implementations are abstrcted -->> you dont have implementations, you just have the signature of the methods
	// the interface eventually get mapped to a concrete type

	// interface values have two components
	// 1- dynamic type and 2- the dynamic value
	// dynamic type is the concrete type that is assigned to
	// dynamic value is actually the value of that dynamic type

	// interface value is actually a pair of (dynamic type and dynamic value)

	var s2 Speaker // s1 is an interface value
	var d2 = Dog{"brian"}
	// dynamic type is Dog, and the dynamic value is d2
	s2 = d2

	s2.Speak()

}

// type specifies(implements) an interface, if type defines all methods specified in the interface.
type Shape2d interface {
	Area() float64
	perimeter() float64
}

type Triangle struct {
	x float64
	y float64
}

func (t *Triangle) Area() float64 {
	return t.x * t.y
}

func (t *Triangle) perimeter() float64 {
	return 2 * (t.x + t.y)
}

type Speaker interface {
	Speak()
}

type Dog struct {
	name string
}

func (d Dog) Speak() {
	fmt.Println(d.name)
}
