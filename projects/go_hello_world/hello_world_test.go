//revive:disable:var-naming
package go_hello_world

//revive:enable:var-naming

import "testing"

func TestGreeter(t *testing.T) {
	expected := "Hello World!"
	actual := HelloWorld()
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}
