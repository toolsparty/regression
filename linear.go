package regression

import (
	"math"
	"fmt"
)

type Linear struct {
	x, y Axis
	k, b float64
}

// add x or y values
// ex. l.Append(nil, y) or l.Append(x, nil) or l.Append(&x, &y)
func (lr *Linear) Append(x, y *AxEl) {
	if x == nil && y == nil {
		panic("the x and y can not be empty")
	}

	var ax AxEl
	n := len(lr.y)
	if x == nil {
		if n == 0 {
			lr.x[n] = 1
		}

		var s float64
		for i := 1; i < n; i++ {
			s += float64(lr.x[i]) - float64(lr.x[i - 1])
		}
		ax = AxEl(float64(lr.x[n - 1]) + s / float64(n - 1))
	} else {
		ax = *x
	}

	lr.x[n] = ax
	lr.y[n] = *y
}

func (lr Linear) Dispersion(a Axis) float64 {
	n := a.Len()
	if n == 0 {
		return float64(0)
	}

	avg := a.Avg()

	d := float64(0)
	for _, val := range a {
		d += math.Pow(avg-float64(val), 2)
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
		cov += (float64(x[i]) - avgX) * (float64(y[i]) - avgY)
	}

	return cov / nx
}

func (lr *Linear) GetK() float64 {
	lr.k = lr.Covariance(lr.x, lr.y) / lr.Dispersion(lr.x)
	return lr.k
}

func (lr *Linear) GetB() float64 {
	lr.b = lr.y.Avg() - lr.GetK()*lr.x.Avg()
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
		res[i] = AxEl(k*float64(val) + b)
	}
	return res, nil
}

func (lr Linear) GetTheta() float64 {
	return lr.Covariance(lr.x, lr.y) / math.Sqrt(lr.Dispersion(lr.x)*lr.Dispersion(lr.y))
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
	reg.x = aX
	reg.y = aY
	return reg, nil
}
