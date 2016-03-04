package procedures

// AnonyFunc is an anonymous function type that can be used as a return value.
type AnonyFunc func() []int

// Parentfunc simply returns AnonyFunc.
func Parentfunc() AnonyFunc {
	return func() []int {
		return []int{1, 2}
	}
}

// ExampleFunc is a simple function in Go that adds two numbers.
func ExampleFunc(a int, b int) (int, error) {
	return a + b, nil
}
