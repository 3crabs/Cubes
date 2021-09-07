package models

//Cube кубик
type Cube struct {
	State int // 0 - точка, 1, 2, 3 - пустые грани
	links []int
}

//Add добавление связи
func (c *Cube) Add(other int) {
	c.links = append(c.links, other)
}

//Rotate поворот кубика и связаных кубиков
func (c *Cube) Rotate(cubes []Cube) {
	c.rotate()
	for _, link := range c.links {
		cubes[link].rotate()
	}
}

// поворот кубика
func (c *Cube) rotate() {
	c.State++
	if c.State > 3 {
		c.State = 0
	}
}
