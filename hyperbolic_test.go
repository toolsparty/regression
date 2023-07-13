package regression

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"math"
	"testing"
)

func TestNewHyperbolic(t *testing.T) {
	reg, err := NewHyperbolic(nil, []int{})
	assert.Error(t, err)
	assert.Nil(t, reg)
}

func TestHyperbolic_GetAB(t *testing.T) {
	// 1 + 2/x
	x := []float32{1, 2, 4, 5, 8, 10}
	y := []float32{3, 2, 1.5, 1.4, 1.25, 1.2}

	reg, err := NewHyperbolic(x, y)
	require.NoError(t, err)
	require.NotNil(t, reg)

	assert.Equal(t, 2., math.Round(reg.GetA()))
	assert.Equal(t, 1., math.Round(reg.GetB()))
}

func TestHyperbolic_Predict(t *testing.T) {
	// 1 + 2/x
	x := []float32{1, 2, 4, 5, 8, 10}
	y := []float32{3, 2, 1.5, 1.4, 1.25, 1.2}

	reg, err := NewHyperbolic(x, y)
	require.NoError(t, err)
	require.NotNil(t, reg)

	got, err := reg.Predict(20)
	assert.NoError(t, err)
	assert.Equal(t, 1.1, math.Round(float64(got)*10)/10)

	reg.Append(3, 4)
	reg.Append(6, 3)
	reg.Append(9, 2)

	got, err = reg.Predict(20)
	assert.NoError(t, err)
	assert.Equal(t, 1.8, math.Round(float64(got)*10)/10)
}
