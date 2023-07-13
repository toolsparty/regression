package regression

import (
	"math"

	"gonum.org/v1/gonum/mat"
)

type Quadratic[T Number] struct {
	ax, ay Axis[T]
}

func (r *Quadratic[T]) Append(x, y T) {
	r.ax = append(r.ax, x)
	r.ay = append(r.ay, y)
}

func (r *Quadratic[T]) GetAbc() (a, b, c float64, err error) {
	const (
		pow4 = 4
		pow3 = 3
		pow2 = 2
	)

	var (
		s4, s3, s2        float64
		sx, sx2y, sxy, sy float64
		n                 = float64(len(r.ax))
	)

	for i, x := range r.ax {
		vx := float64(x)
		vy := float64(r.ay[i])
		s4 += math.Pow(vx, pow4)
		s3 += math.Pow(vx, pow3)
		s2 += math.Pow(vx, pow2)
		sx += vx
		sx2y += math.Pow(vx, pow2) * float64(r.ay[i])
		sxy += vx * vy
		sy += vy
	}

	x4 := s4 / n
	x3 := s3 / n
	x2 := s2 / n
	x := sx / n
	x2y := sx2y / n
	xy := sxy / n
	y := sy / n

	am := mat.NewDense(3, 3, []float64{
		x4, x3, x2,
		x3, x2, x,
		x2, x, 1,
	})
	bm := mat.NewVecDense(3, []float64{x2y, xy, y})
	vec := mat.NewVecDense(3, []float64{0, 0, 0})

	_ = vec.SolveVec(am, bm)
	//if err != nil {
	//	return a, b, c, fmt.Errorf("solve vector: %w", err)
	//}

	raw := vec.RawVector()
	a = raw.Data[0]
	b = raw.Data[1]
	c = raw.Data[2]

	return a, b, c, nil
}

func (r *Quadratic[T]) Predict(x T) (T, error) {
	a, b, c, err := r.GetAbc()
	if err != nil {
		return 0, err
	}

	return T(a*math.Pow(float64(x), 2) + b*float64(x) + c), nil
}

func NewQuadratic[T Number](x, y []T) (*Quadratic[T], error) {
	ax, ay, err := newAxes(x, y)
	if err != nil {
		return nil, err
	}

	return &Quadratic[T]{
		ax: ax,
		ay: ay,
	}, nil
}
