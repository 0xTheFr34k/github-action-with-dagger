package main

import (
	"fmt"
	"m_dagger/internal/point"
)

func main() {
	p := point.Vec2{X: 1.2, Y: 3.4}
	fmt.Println(p.String())
}
