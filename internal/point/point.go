package point

import "fmt"

type Vec[T any] interface {
	Add(T) T
}

type Vec2 struct {
	X, Y float64
}

type Vec3 struct {
	X, Y, Z float64
}

func (p1 Vec2) Add(p2 Vec2) Vec2 {
	return Vec2{X: p1.X + p2.X, Y: p1.Y + p2.Y}
}

func (p1 Vec3) Add(p2 Vec3) Vec3 {
	return Vec3{X: p1.X + p2.X, Y: p1.Y + p2.Y, Z: p1.Z + p2.Z}
}

func (p *Vec2) String() string {
	return fmt.Sprintf("Vec2(%.2f, %.2f)\n", p.X, p.Y)
}

func (p *Vec3) String() string {
	return fmt.Sprintf("Vec3(%.2f, %.2f, %.2f)\n", p.X, p.Y, p.Z)
}
