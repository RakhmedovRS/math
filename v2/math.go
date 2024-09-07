// Package math is a part of my exercises with Go package and module management
// It's part of chapter's 10 1-3 exercises in "Learning Go" book
package math

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds to Number values and returns result
// [addition]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](val1, val2 T) T {
	return val1 + val2
}
