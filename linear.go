package regression

import (
	"math"
	"fmt"
)

type Linear struct {
	Regression
	k, b float64
}

func (lr Linear) Dispersion(a Axis) float64 {
	n := a.Len()
	if n == 0 {
		return float64(0)
	}

	avg := a.Avg()

	d := float64(0)
	for _, val := range a {
		d += math.Pow(avg - val, 2)
	}

	return d / n
}

func (lr Linear) Covariance(x, y Axis) float64 {
	nx, ny := x.Len(), y.Len()
	if nx == 0 || ny == 0 || nx != ny {
		return float64(0)
	}

	avgX, avgY := x.Avg(), y.Avg()
	cov := float64(0)
	for i := range x {
		cov += (x[i] - avgX) * (y[i] - avgY)
	}

	return cov / nx
}

func (lr *Linear) GetK() float64 {
	lr.k = lr.Covariance(lr.X, lr.Y) / lr.Dispersion(lr.X)
	return lr.k
}

func (lr *Linear) GetB() float64 {
	lr.b = lr.Y.Avg() - lr.GetK() * lr.X.Avg()
	return lr.b
}

func (lr Linear) Predict(x Axis) (Axis, error) {
	n := len(x)
	if n == 0 {
		return nil, fmt.Errorf("aX is empty")
	}

	k, b := lr.GetK(), lr.GetB()
	res := make(Axis, n)
	for i, val := range x {
		res[i] = k * val + b
	}
	return res, nil
}

func (lr Linear) GetTheta() float64 {
	return lr.Covariance(lr.X, lr.Y) / math.Sqrt(lr.Dispersion(lr.X) * lr.Dispersion(lr.Y))
}

func NewLinear(x, y []float64) (*Linear, error) {
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

	reg := &Linear{}
	reg.X = aX
	reg.Y = aY
	return reg, nil
}