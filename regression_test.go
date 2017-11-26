package regression

import (
	"testing"
	"reflect"
)

func TestRegression_Len(t *testing.T) {
	X := Axis{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	n := X.Len()
	if n != 5 {
		t.Error("Len error")
	}

	if reflect.TypeOf(n).Kind() != reflect.Float64 {
		t.Error("Len Type error")
	}
}

func TestAxis_Avg(t *testing.T) {
	X := Axis{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	if X.Avg() != 3 {
		t.Error("Avg error")
	}
}

func TestRegression_CreateAx(t *testing.T) {
	Y := Axis{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	X := CreateAxis(Y)
	if len(Y) != len(X) {
		t.Error("Create axis error")
	}
}
