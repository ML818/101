package main

import (
	"fmt"
	"reflect"
)

func main() {
	// testNumber()
	// testCompareStruct()
	testAnonymousTask()
}

/*

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
