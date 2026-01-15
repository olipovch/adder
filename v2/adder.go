// Package adder provides a function to add two integers.
package adder

import "golang.org/x/exp/constraints"

// Number is a type constraint that allows any integer or floating-point type.
type Number interface {
	constraints.Integer | constraints.Float
}

// Adder returns the sum of Number types a and b.
//
// https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
