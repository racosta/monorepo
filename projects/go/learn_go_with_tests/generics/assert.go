package generics

import "testing"

// AssertEqual is a helper function that asserts that two values of any comparable type are equal.
// If they are not, it reports an error with the testing.T instance.
func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

// AssertNotEqual is a helper function that asserts that two values of any comparable type are not equal.
// If they are equal, it reports an error with the testing.T instance.
func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("didn't want %+v", got)
	}
}

// AssertTrue is a helper function that asserts that a boolean value is true.
// If it is not, it reports an error with the testing.T instance.
func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %+v, want true", got)
	}
}

// AssertFalse is a helper function that asserts that a boolean value is false.
// If it is not, it reports an error with the testing.T instance.
func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %+v, want false", got)
	}
}
