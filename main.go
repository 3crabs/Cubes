package main

import (
	"cubes/algorithm"
	"cubes/models"
	"fmt"
	"strconv"
)

//display вывод комбинации кубиков
func display(prefix string, cubes []models.Cube) {
	fmt.Print(prefix)
	for _, cube := range cubes {
		fmt.Printf("%v ", cube.State)
	}
	fmt.Println()
}

func main() {
	n := 5
	cubes := make([]models.Cube, n)
	cubes[1].Rotate(cubes)
	cubes[1].Add(2)

	path, err := algorithm.Find(cubes)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("path", path)
	display("start ", cubes)
	for _, p := range path {
		cubes[p].Rotate(cubes)
		display("("+strconv.Itoa(p)+") -> ", cubes)
	}
}
