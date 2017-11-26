package regression

import "fmt"

type Predictor interface {
	Predict(x Axis) (Axis, error)
}

type Axis map[int]float64

func (a Axis) Len() float64 {
	return float64(len(a))
}

func (a Axis) Avg() float64 {
	n := a.Len()
	if n == 0 {
		return float64(0)
	}

	sum := float64(0)
	for _, val := range a {
		sum += val
	}
	return sum / n
}

func (a Axis) ToArray() []float64 {
	arr := make([]float64, len(a))
	i := 0
	for _, val := range a {
		arr[i] = val
		i++
	}
	fmt.Println(arr)
	return arr
}

func NewAxis(a []float64) (Axis, error) {
	n := len(a)
	if n == 0 {
		return nil, fmt.Errorf("invalid argument")
	}

	ax := make(Axis, n)
	i := 0
	for _, val := range a {
		ax[i] = val
		i++
	}

	return ax, nil
}

type Regression struct {
	X, Y Axis
}

func CreateAxis(a Axis) Axis {
	n := len(a)
	x := make(Axis, n)
	for i := 0; i < n; i++ {
		x[i] = float64(i + 1)
	}
	return x
}