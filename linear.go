package regression

import (
	"math"
)

type Linear[T Number] struct {
	ax, ay Axis[T]
}

// Append add x or y values
// ex. l.Append(nil, y) or l.Append(x, nil) or l.Append(x, y)
func (r *Linear[T]) Append(x, y T) {
	r.ax = append(r.ax, x)
	r.ay = append(r.ay, y)
}

func (r *Linear[T]) GetK() float64 {
	return Covariance(r.ax, r.ay) / Dispersion(r.ax)
}

func (r *Linear[T]) GetB() float64 {
	return r.ay.Avg() - r.GetK()*r.ax.Avg()
}

func (r *Linear[T]) Predict(x T) (T, error) {
	return T(r.GetK()*float64(x) + r.GetB()), nil
}

func (r *Linear[T]) GetTheta() float64 {
	return Covariance(r.ax, r.ay) / math.Sqrt(Dispersion(r.ax)*Dispersion(r.ay))
}

func NewLinear[T Number](x, y []T) (*Linear[T], error) {
	ax, ay, err := newAxes(x, y)
	if err != nil {
		return nil, err
	}

	return &Linear[T]{
		ax: ax,
		ay: ay,
	}, nil
}
