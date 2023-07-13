package regression

// Hyperbolic y = b + a/x
type Hyperbolic[T Number] struct {
	ax, ay Axis[T]
	az     Axis[float64]
}

func (r *Hyperbolic[T]) Append(x, y T) {
	r.ax = append(r.ax, x)
	r.ay = append(r.ay, y)
	r.az = append(r.az, 1/float64(x))
}

func (r *Hyperbolic[T]) GetA() float64 {
	ay := make(Axis[float64], len(r.ay))

	for i, y := range r.ay {
		ay[i] = float64(y)
	}

	return Covariance(r.az, ay) / Dispersion(r.az)
}

func (r *Hyperbolic[T]) GetB() float64 {
	return r.ay.Avg() - r.GetA()*r.az.Avg()
}

func (r *Hyperbolic[T]) Predict(x T) (T, error) {
	return T(r.GetB() + r.GetA()/float64(x)), nil
}

func NewHyperbolic[T Number](x, y []T) (*Hyperbolic[T], error) {
	ax, ay, err := newAxes(x, y)
	if err != nil {
		return nil, err
	}

	az := make(Axis[float64], len(ax))

	for i, vx := range ax {
		az[i] = 1 / float64(vx)
	}

	return &Hyperbolic[T]{
		ax: ax,
		ay: ay,
		az: az,
	}, nil
}
