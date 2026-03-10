package arrays

// Reduce is a higher-order function that takes a collection of elements, a binary function, and an initial value.
// It applies the function cumulatively to the elements of the collection, starting with the initial value, and returns the final result.
func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, element := range collection {
		result = f(result, element)
	}

	return result
}

// Find is a higher-order function that takes a collection of elements and a predicate function.
func Find[A any](collection []A, predicate func(A) bool) (value A, found bool) {
	for _, v := range collection {
		if predicate(v) {
			return v, true
		}
	}
	return
}
