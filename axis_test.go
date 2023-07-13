package regression

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAxisFrom(t *testing.T) {
	values := []int{1, 2, 4, 7, 8, 9, 12}
	ay := NewAxis(values...)
	ax := NewAxisFrom(ay)
	assert.Len(t, ax, len(ay))
	assert.Equal(t, ax, Axis[int]{1, 2, 3, 4, 5, 6, 7})
}

func TestDispersion(t *testing.T) {
	iv := []int32{-11, -7, -5, -3, -1, 0, 1, 2, 4, 8, 12, 16, 32}
	ax := NewAxis(iv...)

	fv := []float32{-11, -7, -5, -3, -1, 0, 1, 2, 4, 8, 12, 16, 32}
	ay := NewAxis(fv...)
	assert.Equal(t, Dispersion(ax), Dispersion(ay))
	assert.Equal(t, 118.2130177514793, Dispersion(ay))

	az := Axis[uint32]{}
	assert.Equal(t, .0, Dispersion(az))
}

func TestCovariance(t *testing.T) {
	xv := []int64{0, 1, 2, 3, 4, 5, 10}
	yv := []int64{0, 2, 4, 6, 8, 10, 20}

	ax := NewAxis(xv...)
	ay := NewAxis(yv...)

	c := Covariance(ax, ay)
	assert.Equal(t, 18.775510204081634, c)

	c = Covariance(ax, ax)
	assert.Equal(t, 9.387755102040817, c)

	c = Covariance(ay, ay)
	assert.Equal(t, 37.55102040816327, c)

	az := Axis[int64]{}
	assert.Equal(t, .0, Covariance(ax, az))
	assert.Equal(t, .0, Covariance(ay, az))
	assert.Equal(t, .0, Covariance(az, ax))
	assert.Equal(t, .0, Covariance(az, az))
}

func TestAxis_Avg(t *testing.T) {
	xv := []float64{0, 1, 2, 3, 4, 5}
	ax := NewAxis(xv...)
	avg := ax.Avg()
	assert.Equal(t, 2.5, avg)

	ay := Axis[uint]{}
	assert.Equal(t, .0, ay.Avg())
}
