package regression

import  (
	"math"
	"gonum.org/v1/gonum/mat"
	"fmt"
)

type Quadratic struct {
	x, y  Axis
	a, b, c float64
}

func (q Quadratic) avgXY() float64 {
	xy := float64(0)
	for i := range q.x {
		xy += float64(q.x[i]) * float64(q.y[i])
	}
	return xy / q.x.Len()
}

func (q Quadratic) xyAvg() float64 {
	return q.x.Avg() * q.y.Avg()
}

func (q Quadratic) avgXPow(p float64) float64 {
	x2 := float64(0)
	for _, x := range q.x {
		x2 += math.Pow(float64(x), p)
	}
	return x2 / q.x.Len()
}

func (q Quadratic) avgX2Y() float64 {
	x2y := float64(0)
	for i := range q.x {
		x2y += math.Pow(float64(q.x[i]), 2) * float64(q.y[i])
	}
	return x2y / q.x.Len()
}

func (q *Quadratic) GetAbc() (float64, float64, float64) {
	x4 := q.avgXPow(4)
	x3 := q.avgXPow(3)
	x2 := q.avgXPow(2)
	x := q.x.Avg()
	x2y := q.avgX2Y()
	xy := q.avgXY()
	y := q.y.Avg()

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
		res[i] = AxEl(a * math.Pow(float64(val), 2) + b * float64(val) + c)
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
	reg.x = aX
	reg.y = aY
	return reg, nil
}