package point

import "fmt"

type Point struct {
	X, Y float64
}

func (p *Point) String() string {
	return fmt.Sprintf("Point(%.2f , %.2f)\n", p.X, p.Y)
}
