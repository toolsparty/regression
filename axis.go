package regression

import (
	"fmt"
	"math"
)

type Number interface {
	int | int16 | int8 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

type Axis[T Number] []T

func (a Axis[T]) Avg() float64 {
	n := len(a)

	if n == 0 {
		return 0
	}

	sum := T(0)

	for _, v := range a {
		sum += v
	}

	return float64(sum) / float64(n)
}

func NewAxis[T Number](values ...T) Axis[T] {
	n := len(values)

	axis := make(Axis[T], n)

	for i, value := range values {
		axis[i] = value
	}

	return axis
}

func NewAxisFrom[T Number](a Axis[T]) Axis[T] {
	n := len(a)
	x := make(Axis[T], n)

	for i := 0; i < n; i++ {
		x[i] = T(i + 1)
	}

	return x
}

func newAxes[T Number](x, y []T) (ax Axis[T], ay Axis[T], err error) {
	nx, ny := len(x), len(y)
	if ny == 0 {
		return nil, nil, fmt.Errorf("axis y is empty")
	}

	if nx > 0 && nx != ny {
		return nil, nil, fmt.Errorf("axes x and y has a different lengths")
	}

	ay = NewAxis(y...)

	if nx == 0 {
		ax = NewAxisFrom(ay)
	} else {
		ax = NewAxis(x...)
	}

	return ax, ay, nil
}

func Dispersion[T Number](axis Axis[T]) float64 {
	n := len(axis)

	if n == 0 {
		return 0
	}

	avg := axis.Avg()

	d := float64(0)
	for _, value := range axis {
		d += math.Pow(avg-float64(value), 2)
	}

	return d / float64(n)
}

func Covariance[T Number](ax, ay Axis[T]) float64 {
	nx, ny := len(ax), len(ay)
	if nx == 0 || ny == 0 || nx != ny {
		return float64(0)
	}

	avgX, avgY := ax.Avg(), ay.Avg()
	cov := float64(0)
	for i := range ax {
		cov += (float64(ax[i]) - avgX) * (float64(ay[i]) - avgY)
	}

	return cov / float64(nx)
}
