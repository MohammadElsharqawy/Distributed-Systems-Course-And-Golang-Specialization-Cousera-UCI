package data

import "fmt"

type Point struct {
	x float64
	y float64
}

// good practise to have all methods have a pointer reciever
// or all methods have non-pointer reciever
// mixing them for a type will get confusing
// pointer reciever allows modification

// when we start it with capital letter, is public, we can access it from outside say (main package)
func (p *Point) InitMe(xn, yn float64) {
	(*p).x = xn
	p.y = yn // it makes the dereference implicitly--->> compiler recognize it

	// no need to reference either --->> (p1).Initme(), we include p1 itself not a reference --->> compiler recognize it
}

func (p *Point) Scale(factor float64) {
	p.x *= factor
	p.y *= factor
}

func (p *Point) PrintMe() {
	fmt.Println(p.x, p.y)
}

func (p *Point) GetX() float64 {
	return (*p).x
}

func (p *Point) GetY() float64 {
	return p.y
}
