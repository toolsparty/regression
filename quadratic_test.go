package regression

import (
	"testing"
)

func TestQuadratic_GetAbc(t *testing.T) {
	x, _ := NewAxis([]float64{1.5, 2.5, 3.5, 4.5, 5.5})
	y, _ := NewAxis([]float64{3, 5, 7, 9, 11})
	reg := Quadratic{}
	reg.x = x
	reg.y = y
	a, b, c := reg.GetAbc()

	if a != 0 || b != 2 || c != 0 {
		t.Error("a, b, c error")
	}
}
