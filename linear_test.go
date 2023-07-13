package regression

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewLinear(t *testing.T) {
	testCases := []*struct {
		Name    string
		X, Y    []float64
		WantErr string
	}{
		{
			Name:    "axis y is empty",
			X:       nil,
			Y:       nil,
			WantErr: "axis y is empty",
		},
		{
			Name:    "axis y is empty",
			X:       nil,
			Y:       []float64{},
			WantErr: "axis y is empty",
		},
		{
			Name:    "axis y is empty",
			X:       []float64{},
			Y:       []float64{},
			WantErr: "axis y is empty",
		},
		{
			Name:    "axis y is empty",
			X:       []float64{1},
			Y:       []float64{},
			WantErr: "axis y is empty",
		},
		{
			Name:    "different lengths",
			X:       []float64{1},
			Y:       []float64{1, 2},
			WantErr: "axes x and y has a different lengths",
		},
		{
			Name:    "success",
			X:       []float64{},
			Y:       []float64{1, 2},
			WantErr: "",
		},
		{
			Name:    "success",
			X:       []float64{},
			Y:       []float64{1, 2},
			WantErr: "",
		},
		{
			Name:    "success",
			X:       []float64{2, 4},
			Y:       []float64{1, 2},
			WantErr: "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			got, err := NewLinear(testCase.X, testCase.Y)
			if testCase.WantErr != "" {
				assert.EqualError(t, err, testCase.WantErr)
				assert.Nil(t, got)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, got)
			}
		})
	}
}

func TestLinear_GetKB(t *testing.T) {
	y := []int{3, 5, 7, 9, 11}
	reg, err := NewLinear(nil, y)
	require.NoError(t, err)
	require.NotNil(t, reg)

	assert.Equal(t, 2.0, reg.GetK())
	assert.Equal(t, 1.0, reg.GetB())
}

func TestLinear_Predict(t *testing.T) {
	y := []int{3, 5, 7, 9, 11}
	reg, err := NewLinear(nil, y)
	require.NoError(t, err)
	require.NotNil(t, reg)

	got, err := reg.Predict(2)
	assert.NoError(t, nil)
	assert.Equal(t, 5, got)

	got, err = reg.Predict(5)
	assert.NoError(t, nil)
	assert.Equal(t, 11, got)

	got, err = reg.Predict(10)
	assert.NoError(t, nil)
	assert.Equal(t, 21, got)
}

func TestLinear_GetTheta(t *testing.T) {
	y := []int{3, 5, 7, 9, 11}
	reg, err := NewLinear(nil, y)
	require.NoError(t, err)
	require.NotNil(t, reg)

	theta := reg.GetTheta()
	assert.Equal(t, 1.0, theta)

	reg.Append(6, 20)

	assert.Equal(t, 3.0, reg.GetK())
	assert.Equal(t, -1.333333333333334, reg.GetB())
}
