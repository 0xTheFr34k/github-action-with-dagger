package tests

import (
	"github.com/stretchr/testify/assert"
	"m_dagger/internal/point"
	"testing"
)

func TestPointString(t *testing.T) {
	p := point.Point{X: 1.5, Y: 2.75}
	assert.Equal(t, p.String(), "Point(1.50 , 2.75)\n")

}
