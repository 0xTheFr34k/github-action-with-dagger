package tests

import (
	"m_dagger/internal/point"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	p := point.Vec2{X: 1.5, Y: 2.75}

	p1 := point.Vec3{X: 1.5, Y: 2.75, Z: 3.0}

	assert.Equal(t, p.String(), "Vec2(1.50, 2.75)\n")
	assert.Equal(t, p1.String(), "Vec3(1.50, 2.75, 3.00)\n")
}

func TestInterface(t *testing.T) {
	arr := make([]point.Vec2, 2)
	arr1 := make([]point.Vec3, 2)

	for i := range 2 {
		arr[i] = point.Vec2{}
		arr1[i] = point.Vec3{}
	}

}

func TestAdd(t *testing.T) {
	p := point.Vec2{X: 1.5, Y: 2.75}

	p = p.Add(point.Vec2{X: 1.0, Y: 1.0})

	assert.Equal(t, p.X, 2.5)
	assert.Equal(t, p.Y, 3.75)
}
