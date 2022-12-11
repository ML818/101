package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// testNumber()
	// testCompareStruct()
	// testAnonymousTask()
	// testPointMethod()
	// testRectangle()
	testEmployee()
}

/*

	task4
	Define a struct employee with a field salary and make a method giveRaise() for this type to increase the salary with a certain percentage.

*/

type employee struct {
	salary float32
}

func (e *employee) giveRaise(percentage float32) {
	e.salary += percentage * e.salary
}

func testEmployee() {
	e1 := new(employee)
	e1.salary = 20000

	e1.giveRaise(0.4)

	fmt.Println("Salary change to: ", e1.salary)
}

/*

	task3
	focus on method.
	Define a struct Rectangle with int properties length and width. Give this type two methods Area() and Perimeter() and test it out.

*/

type Rectangle struct {
	length int
	width  int
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func (r *Rectangle) Perimeter() int {
	return 2*r.length + 2*r.width
}

func testRectangle() {
	r := &Rectangle{5, 9}
	fmt.Println(r.Area())
	fmt.Println(r.Perimeter())
}

/*

	task2:
	focus on method.
	Define a 2 dimensional Point with coordinates X and Y as a struct. Implement a method Abs() that calculates the length of the vector represented by a Point, and a method Scale () that multiplies the coordinates of a point with a scale factor.

*/

type Point struct {
	X float64
	Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p *Point) Scale(s float64) {
	p.X = p.X * s
	p.Y = p.Y * s
}

func testPointMethod() {
	p1 := new(Point)
	p1.X = 6
	p1.Y = 8

	fmt.Println("the length of the vector p1 is:", p1.Abs())

	p2 := &Point{4, 5}
	fmt.Println("the length of the vector p2 is:", p2.Abs())

	p1.Scale(7)
	fmt.Printf("Point p1 scaled 7 has the following coordinates: X %f - Y %f\n", p1.X, p1.Y)
}

/*
	task1:
	focus on anonymous and tags
	Make a struct with 1 named float field, and 2 anonymous fields of type int and string. Create a value of the struct with a literal expression and print the content.

*/

type anonymousTask struct {
	point float32 "This is a point of float"

	int    "set a number whatever you want"
	string "set a string whatever you want"
}

func testAnonymousTask() {
	anonym := new(anonymousTask)
	anonym.point = 3.14
	anonym.int = 3
	anonym.string = "This is a string"

	anonymTags := reflect.TypeOf(*anonym)

	fmt.Println(anonymTags.Field(0).Tag)
	fmt.Println(anonymTags.Field(1).Tag)
	fmt.Println(anonymTags.Field(2).Tag)

	fmt.Println(anonym)
}

/*
	convert struct types

	If different struct types both have a same element name and type, it can be same by convert type.

*/

type T1 struct {
	a int
}

type T2 struct {
	b int
}

type T3 struct {
	a int
}

func testCompareStruct() {
	t1 := T1{10}
	t2 := T2{10}
	t3 := T3{10}
	t4 := T3{20}

	fmt.Println(t1, t2, t3, t4)

	fmt.Println(t1 == T1(t3)) // True
	fmt.Println(t3 == t4)     // False

}

/*
	struct type alias

	when converting struct type, if use equal operator with type, it is an alias, and the value totally same as the variable.
	if without equal operator, it will create a new distinct type.
*/

type number struct {
	f float32
}

type nr number        // new distinct type
type nrAlias = number // alias type

func testNumber() {
	a := number{5.0}
	b := nr{5.0}
	c := nrAlias{5.0}

	fmt.Printf("type of a is %T.\n", a)
	fmt.Printf("type of b is %T.\n", b)
	fmt.Printf("type of c is %T.\n", c)
	// var i float32 = b   // compile-error: cannot use b (type nr) as type float32 in assignment
	// var i = float32(b)  // compile-error: cannot convert b (type nr) to type float32
	// var d number = b    // compile-error: cannot use b (type nr) as type number in assignment
	// needs a conversion:
	var d = number(b)
	fmt.Printf("type of d is %T.\n", d)
	// an alias doesn't need conversion:
	var e number = c
	fmt.Printf("type of e is %T.\n", e)

	fmt.Println(a, b, c, d, e)
}
