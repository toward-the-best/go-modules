package greeting

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hi, jinkwon.kim. Welcome!"
	if got := Hello("jinkwon.kim"); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
	
}