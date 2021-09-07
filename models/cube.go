package models

//Cube кубик
type Cube struct {
	State int   // 0, 1, 2, 3 - номер грани повернутой к нам
	links []int // номера кубов которые поворачиваются при повороте этого куба
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
