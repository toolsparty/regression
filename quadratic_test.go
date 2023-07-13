package regression

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewQuadratic(t *testing.T) {
	reg, err := NewQuadratic(nil, []float64{})
	assert.Error(t, err)
	assert.Nil(t, reg)
}

func TestQuadratic_GetAbc(t *testing.T) {
	// x^2 + 2x - 1
	x := []float64{1, 2, 3, 5, 6}
	y := []float64{2, 7, 14, 34, 47}

	reg, err := NewQuadratic(x, y)
	require.NoError(t, err)
	require.NotNil(t, reg)

	a, b, c, err := reg.GetAbc()
	assert.NoError(t, err)
	assert.Equal(t, 1.0, math.Round(a))
	assert.Equal(t, 2.0, math.Round(b))
	assert.Equal(t, -1.0, math.Round(c))
}

func TestQuadratic_Predict(t *testing.T) {
	// x^2 + 2x - 1
	x := []int{1, 2, 3, 5, 6}
	y := []int{2, 7, 14, 34, 47}

	reg, err := NewQuadratic(x, y)
	require.NoError(t, err)
	require.NotNil(t, reg)

	a, b, c, err := reg.GetAbc()
	assert.NoError(t, err)
	assert.Equal(t, 1.0, math.Round(a))
	assert.Equal(t, 2.0, math.Round(b))
	assert.Equal(t, -1.0, math.Round(c))

	got, err := reg.Predict(7)
	assert.NoError(t, err)
	assert.Equal(t, 62, got)

	reg.Append(7, -65)

	a, b, c, err = reg.GetAbc()
	assert.NoError(t, err)
	assert.Equal(t, -7.423469387755153, a)
	assert.Equal(t, 55.78061224489836, b)
	assert.Equal(t, -63.204081632653605, c)

	got, err = reg.Predict(10)
	assert.NoError(t, err)
	assert.Equal(t, -247, got)
}
