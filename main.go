package main

import "fmt"

//Cube кубик
type Cube struct {
	State int // 0 - точка, 1, 2, 3 - пустые грани
	links []*Cube
}

// поворот кубика
func (c *Cube) rotate() {
	c.State++
	if c.State > 3 {
		c.State = 0
	}
}

//RotateDeep поворот кубика и связаных кубиков
func (c *Cube) RotateDeep() {
	c.rotate()
	for _, link := range c.links {
		link.rotate()
	}
}

//Add добавление связи
func (c *Cube) Add(other *Cube) {
	c.links = append(c.links, other)
}

func main() {
	n := 5
	cubes := make([]Cube, n)
	cubes[1].RotateDeep()
	cubes[0].Add(&cubes[1])
	cubes[0].Add(&cubes[2])
	cubes[0].RotateDeep()

	for _, cube := range cubes {
		fmt.Println(cube)
	}
}
