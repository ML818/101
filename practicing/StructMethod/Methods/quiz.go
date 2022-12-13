package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("================================================")
	fmt.Println("Change salary")
	testEmployee()

	fmt.Println("\n================================================")
	fmt.Println("Calculate Area and Perimeter of rectangle")
	testRectangle()

	fmt.Println("\n================================================")
	fmt.Println("Coordinates and Vector")
	testPointMethod()
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
