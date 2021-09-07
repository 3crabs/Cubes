package main

import (
	"errors"
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
	for i := 0; i < len(cubes)-1; i++ {
		if cubes[i].State != cubes[i+1].State {
			return false
		}
	}
	return true
}

func findPath() ([]int, error) {
	for len(steps) > 0 {
		step := steps[0]
		steps = steps[1:]
		step.cubes[step.nextCubeNumber].RotateDeep(step.cubes)
		if isWin(step.cubes) {
			return step.path, nil
		}
		for i := 0; i < len(step.cubes); i++ {
			tmpCubes := make([]Cube, len(step.cubes))
			copy(tmpCubes, step.cubes)
			tmpPath := make([]int, len(step.path))
			copy(tmpPath, step.path)
			steps = append(steps, Step{cubes: tmpCubes, nextCubeNumber: i, path: append(tmpPath, i)})
		}
	}
	return nil, errors.New("path not found")
}

func main() {
	n := 5
	cubes := make([]Cube, n)
	cubes[1].Add(2)
	cubes[1].Add(4)
	cubes[1].RotateDeep(cubes)
	cubes[2].RotateDeep(cubes)

	for i := 0; i < n; i++ {
		tmp := make([]Cube, len(cubes))
		copy(tmp, cubes)
		steps = append(steps, Step{cubes: tmp, nextCubeNumber: i, path: []int{i}})
	}
	path, err := findPath()
	if err != nil {
		panic(err)
	}
	fmt.Println("path", path)
	display("start ", cubes)
	for _, p := range path {
		cubes[p].RotateDeep(cubes)
		display("("+strconv.Itoa(p)+") -> ", cubes)
	}
}
