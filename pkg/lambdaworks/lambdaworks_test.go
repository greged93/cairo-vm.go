package lambdaworks_test

import (
	"reflect"
	"testing"

	"github.com/lambdaclass/cairo-vm.go/pkg/lambdaworks"
)

func TestFromHex(t *testing.T) {
	var h_one = "1a"
	expected := lambdaworks.FeltFromUint64(26)

	result := lambdaworks.FeltFromHex(h_one)
	if result != expected {
		t.Errorf("TestFromHex failed. Expected: %v, Got: %v", expected, result)
	}

}

func TestFromDecString(t *testing.T) {
	var s_one = "435"
	expected := lambdaworks.FeltFromUint64(435)

	result := lambdaworks.FeltFromDecString(s_one)
	if result != expected {
		t.Errorf("TestFromDecString failed. Expected: %v, Got: %v", expected, result)
	}
}

func TestFromNegDecString(t *testing.T) {
	var s_one = "-1"
	expected := lambdaworks.FeltFromHex("800000000000011000000000000000000000000000000000000000000000000")

	result := lambdaworks.FeltFromDecString(s_one)
	if result != expected {
		t.Errorf("TestFromNegDecString failed. Expected: %v, Got: %v", expected, result)
	}
}

func TestToLeBytes(t *testing.T) {
	expected := [32]uint8{
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	actual := *lambdaworks.FeltOne().ToLeBytes()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestToLeBytes failed. Expected: %v, Got: %v", expected, actual)
	}
}

func TestToBeBytes(t *testing.T) {
	expected := [32]uint8{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
	}
	actual := *lambdaworks.FeltOne().ToBeBytes()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("TestToBeBytes failed. Expected: %v, Got: %v", expected, actual)
	}
}

func TestFromLeBytes(t *testing.T) {
	bytes := [32]uint8{
		1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	felt_from_bytes := lambdaworks.FeltFromLeBytes(&bytes)

	if !reflect.DeepEqual(felt_from_bytes, lambdaworks.FeltOne()) {
		t.Errorf("TestFromLeBytes failed. Expected 1, Got: %v", felt_from_bytes)
	}
}

func TestFromBeBytes(t *testing.T) {
	bytes := [32]uint8{
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1,
	}
	felt_from_bytes := lambdaworks.FeltFromBeBytes(&bytes)

	if !reflect.DeepEqual(felt_from_bytes, lambdaworks.FeltOne()) {
		t.Errorf("TestToFromBeBytes failed. Expected 1, Got: %v", felt_from_bytes)
	}
}

func TestFeltSub(t *testing.T) {
	f_one := lambdaworks.FeltOne()
	expected := lambdaworks.FeltZero()

	result := f_one.Sub(f_one)
	if result != expected {
		t.Errorf("TestFeltSub failed. Expected: %v, Got: %v", expected, result)
	}
}

func TestFeltAdd(t *testing.T) {
	f_zero := lambdaworks.FeltZero()
	f_one := lambdaworks.FeltOne()
	expected := lambdaworks.FeltOne()

	result := f_zero.Add(f_one)
	if result != expected {
		t.Errorf("TestFeltAdd failed. Expected: %v, Got: %v", expected, result)
	}
}

func TestFeltMul1(t *testing.T) {
	f_one := lambdaworks.FeltOne()
	expected := lambdaworks.FeltOne()

	result := f_one.Mul(f_one)
	if result != expected {
		t.Errorf("TestFeltMul1 failed. Expected: %v, Got: %v", expected, result)
	}
}

func TestFeltMul0(t *testing.T) {
	f_one := lambdaworks.FeltOne()
	f_zero := lambdaworks.FeltZero()
	expected := lambdaworks.FeltZero()

	result := f_zero.Mul(f_one)
	if result != expected {
		t.Errorf("TestFeltMul0 failed. Expected: %v, Got: %v", expected, result)
	}
}

func TestFeltMul9(t *testing.T) {
	f_three := lambdaworks.FeltFromUint64(3)
	expected := lambdaworks.FeltFromUint64(9)

	result := f_three.Mul(f_three)
	if result != expected {
		t.Errorf("TestFeltMul9 failed. Expected: %v, Got: %v", expected, result)
	}
}

func TestFeltDiv3(t *testing.T) {
	f_three := lambdaworks.FeltFromUint64(3)
	expected := lambdaworks.FeltFromUint64(1)

	result := f_three.Div(f_three)
	if result != expected {
		t.Errorf("TestFeltDiv3 failed. Expected: %v, Got: %v", expected, result)
	}
}

func TestFeltDiv4(t *testing.T) {
	f_four := lambdaworks.FeltFromUint64(4)
	f_two := lambdaworks.FeltFromUint64(2)

	expected := lambdaworks.FeltFromUint64(2)

	result := f_four.Div(f_two)
	if result != expected {
		t.Errorf("TestFeltDiv4 failed. Expected: %v, Got: %v", expected, result)
	}
}

func TestFeltDiv4Error(t *testing.T) {
	f_four := lambdaworks.FeltFromUint64(4)
	f_one := lambdaworks.FeltFromUint64(1)

	expected := lambdaworks.FeltFromUint64(45)

	result := f_four.Div(f_one)
	if result == expected {
		t.Errorf("TestFeltDiv4Error failed. Expected: %v, Got: %v", expected, result)
	}
}
