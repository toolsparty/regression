package regression

import (
	"testing"
)

func TestNewLinearRegression(t *testing.T) {
	var y []float64

	reg, err := NewLinear([]float64{}, y)
	if err == nil {
		t.Error("No error")
	}

	y = []float64{3, 5, 7, 9, 11}
	reg, err = NewLinear([]float64{}, y)
	if err != nil {
		t.Error(err)
	}

	if int(reg.X.Len()) != len(y) {
		t.Error("Lengths error")
	}

	if reg.GetK() != 2 {
		t.Error("invalid K")
	}

	if reg.GetB() != 1 {
		t.Error("invalid B")
	}

	x := []float64{1.5, 2.5, 3.5, 4.5, 5.5}
	reg, err = NewLinear(x, y)
	if err != nil {
		t.Error(err)
	}

	if int(reg.X.Len()) != len(y) {
		t.Error("Lengths error")
	}

	if reg.GetK() != 2 {
		t.Error("invalid K")
	}

	if reg.GetB() != 0 {
		t.Error("invalid B")
	}

	nX, _ := NewAxis([]float64{10, 20, 30})
	nY, _ := reg.Predict(nX)

	cY, _ := NewAxis([]float64{20, 40, 60})

	for i := range nY {
		if nY[i] != cY[i] {
			t.Error("Prediction error")
		}
	}
}
