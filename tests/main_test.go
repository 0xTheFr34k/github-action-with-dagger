package tests

import (
	"m_dagger/internal/point"
	"testing"
)

func TestPointString(t *testing.T) {
	p := point.Point{X: 1.5, Y: 2.75}

	want := "Point(1.50 , 2.75)\n"

	if got := p.String(); got != want {
		t.Errorf("expected %q, got %q", want, got)
	}
}
