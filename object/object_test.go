package object

import "testing"

func TestStringHashKey(t *testing.T) {
	val1 := &String{Value: "Hello World"}
	val2 := &String{Value: "Hello World"}
	vala := &String{Value: "My name is johnny"}
	valb := &String{Value: "My name is johnny"}

	if val1.HashKey() != val2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if vala.HashKey() != valb.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if val1.HashKey() == vala.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestIntegerHashKey(t *testing.T) {
	val1 := &Integer{Value: 2}
	val2 := &Integer{Value: 2}
	vala := &Integer{Value: 4}
	valb := &Integer{Value: 4}

	if val1.HashKey() != val2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if vala.HashKey() != valb.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if val1.HashKey() == vala.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}

func TestBooleanHashKey(t *testing.T) {
	val1 := &Boolean{Value: true}
	val2 := &Boolean{Value: true}
	vala := &Boolean{Value: false}
	valb := &Boolean{Value: false}

	if val1.HashKey() != val2.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if vala.HashKey() != valb.HashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if val1.HashKey() == vala.HashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}
