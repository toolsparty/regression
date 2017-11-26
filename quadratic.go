package regression

import  (
	"math"
	"gonum.org/v1/gonum/mat"
	"fmt"
)

type Quadratic struct {
	Regression
	a, b, c float64
}

func (q Quadratic) avgXY() float64 {
	xy := float64(0)
	for i := range q.X {
		xy += q.X[i] * q.Y[i]
	}
	return xy / q.X.Len()
}

func (q Quadratic) xyAvg() float64 {
	return q.X.Avg() * q.Y.Avg()
}

func (q Quadratic) avgXPow(p float64) float64 {
	x2 := float64(0)
	for _, x := range q.X {
		x2 += math.Pow(x, p)
	}
	return x2 / q.X.Len()
}

func (q Quadratic) avgX2Y() float64 {
	x2y := float64(0)
	for i := range q.X {
		x2y += math.Pow(q.X[i], 2) * q.Y[i]
	}
	return x2y / q.X.Len()
}

func (q *Quadratic) GetAbc() (float64, float64, float64) {
	x4 := q.avgXPow(4)
	x3 := q.avgXPow(3)
	x2 := q.avgXPow(2)
	x := q.X.Avg()
	x2y := q.avgX2Y()
	xy := q.avgXY()
	y := q.Y.Avg()

	A := mat.NewDense(3, 3, []float64{
		x4, x3, x2,
		x3, x2, x,
		x2, x, 1,
	})
	B := mat.NewVecDense(3, []float64{x2y, xy, y})
	v := mat.NewVecDense(3, []float64{0, 0, 0})
	v.SolveVec(A, B)
	q.a = v.RawVector().Data[0]
	q.b = v.RawVector().Data[1]
	q.c = v.RawVector().Data[2]
	return q.a, q.b, q.c
}

func (q Quadratic) Predict(x Axis) (Axis, error) {
	n := len(x)
	if n == 0 {
		return nil, fmt.Errorf("aX is empty")
	}

	a, b, c := q.GetAbc()
	res := make(Axis, n)
	for i, val := range x {
		res[i] = a * math.Pow(val, 2) + b * val + c
	}
	return res, nil
}

func NewQuadratic(x, y []float64) (*Quadratic, error) {
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

	reg := &Quadratic{}
	reg.X = aX
	reg.Y = aY
	return reg, nil
}