package main

import (
	"fmt"
	"strconv"
)

//Cube кубик
type Cube struct {
	State int // 0 - точка, 1, 2, 3 - пустые грани
	links []int
}

// поворот кубика
func (c *Cube) rotate() {
	c.State++
	if c.State > 3 {
		c.State = 0
	}
}

//RotateDeep поворот кубика и связаных кубиков
func (c *Cube) RotateDeep(cubes []Cube) {
	c.rotate()
	for _, link := range c.links {
		cubes[link].rotate()
	}
}

//Add добавление связи
func (c *Cube) Add(other int) {
	c.links = append(c.links, other)
}

type Step struct {
	cubes          []Cube
	nextCubeNumber int
	path           []int
}

var steps []Step

func display(prefix string, cubes []Cube) {
	fmt.Print(prefix)
	for _, cube := range cubes {
		fmt.Printf("%v ", cube.State)
	}
	fmt.Println()
}

func isWin(cubes []Cube) bool {
	win := true
	for i := 0; i < len(cubes)-1; i++ {
		if cubes[i].State != cubes[i+1].State {
			win = false
		}
	}
	return win
}

func findPath() {
	for len(steps) > 0 {
		step := steps[0]
		steps = steps[1:]
		step.cubes[step.nextCubeNumber].RotateDeep(step.cubes)
		display("("+strconv.Itoa(step.nextCubeNumber)+") -> ", step.cubes)
		if isWin(step.cubes) {
			display("end ", step.cubes)
			fmt.Println("path:", step.path)
			return
		}
		for i := 0; i < len(step.cubes); i++ {
			tmp := make([]Cube, len(step.cubes))
			copy(tmp, step.cubes)
			steps = append(steps, Step{cubes: tmp, nextCubeNumber: i, path: append(step.path, i)})
		}
	}
}

func main() {
	n := 5
	cubes := make([]Cube, n)
	cubes[1].RotateDeep(cubes)
	cubes[1].Add(2)

	for i := 0; i < n; i++ {
		tmp := make([]Cube, len(cubes))
		copy(tmp, cubes)
		steps = append(steps, Step{cubes: tmp, nextCubeNumber: i, path: []int{i}})
	}
	display("start ", cubes)
	findPath()
}
