package algorithm

import (
	"cubes/models"
	"errors"
)

//steps список возможных комбинаций
var steps []models.Step

//Find поиск кратчайшего пути к выигрышной комбинации
func Find(cubes []models.Cube) ([]int, error) {
	for i := 0; i < len(cubes); i++ {
		tmp := make([]models.Cube, len(cubes))
		copy(tmp, cubes)
		steps = append(steps, models.Step{Cubes: tmp, NextCubeNumber: i, Path: []int{i}})
	}
	return findPath()
}

//findPath поиск кратчайшего пути к выигрышной комбинации
func findPath() ([]int, error) {
	for k := 0; len(steps) > 0 && k < 1000; k++ {
		step := steps[0]
		steps = steps[1:]
		step.Cubes[step.NextCubeNumber].Rotate(step.Cubes)
		if step.IsWin() {
			return step.Path, nil
		}
		for i := 0; i < len(step.Cubes); i++ {
			tmpCubes := make([]models.Cube, len(step.Cubes))
			copy(tmpCubes, step.Cubes)
			tmpPath := make([]int, len(step.Path))
			copy(tmpPath, step.Path)
			steps = append(steps, models.Step{Cubes: tmpCubes, NextCubeNumber: i, Path: append(tmpPath, i)})
		}
	}
	return nil, errors.New("path not found")
}
