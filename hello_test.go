package hello

import "testing"

func TestHelloThere(t *testing.T) {
	got, error := Hello("there")
	if error != nil {
		t.Errorf("Whould not return error")
	}
	if got != "Hello, there!" {
		t.Errorf("Hello(there) = %v; want Hello, there!", got)
	}
}
