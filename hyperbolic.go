package regression

import (
	"fmt"
)

// y = b + a/x
type Hyperbolic struct {
	Linear
	a, b float64
	z Axis
}

func (h *Hyperbolic) setZ() {
	if len(h.z) == 0 {
		n := len(h.X)
		h.z = make(Axis, n)
		for i, x := range h.X {
			h.z[i] = 1 / x
		}
	}
}

func (h *Hyperbolic) GetA() float64 {
	h.setZ()
	h.a = h.Covariance(h.z, h.Y) / h.Dispersion(h.z)
	return h.a
}

func (h Hyperbolic) GetB() float64 {
	h.GetA()
	h.b = h.Y.Avg() - h.a * h.z.Avg()
	return h.b
}

func (h Hyperbolic) Predict(x Axis) (Axis, error) {
	n := len(x)
	if n == 0 {
		return nil, fmt.Errorf("aX is empty")
	}

	a, b := h.GetA(), h.GetB()
	res := make(Axis, n)
	for i, val := range x {
		res[i] = b + a / val
	}
	return res, nil
}

func NewHyperbolic(x, y []float64) (*Hyperbolic, error) {
	nx, ny := len(x), len(y)
	if ny == 0 {
		return nil, fmt.Errorf("axis y is empty")
	}

	if nx > 0 && nx != ny {
		return nil, fmt.Errorf("axes x and y has different lengths")
	}

	aY, _ := NewAxis(y)

	var aX Axis
	if nx == 0 {
		aX = CreateAxis(aY)
	} else {
		aX, _ = NewAxis(x)
	}

	reg := &Hyperbolic{}
	reg.X = aX
	reg.Y = aY
	return reg, nil
}