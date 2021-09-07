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
	states := []int{1, 0, 1, 0}
	n := len(states)
	links := [][]int{{0, 3}, {1, 2}, {1, 3}}

	// создание кубов
	cubes := make([]models.Cube, n)

	// задание первоначального положения
	for i := 0; i < n; i++ {
		if states[i] < 0 || 3 < states[i] {
			continue
		}
		cubes[i].State = states[i]
	}

	// задание связей
	for _, link := range links {
		if link[0] < 0 || n <= link[0] {
			continue
		}
		if link[1] < 0 || n <= link[1] {
			continue
		}
		cubes[link[0]].Add(link[1])
	}

	// поиск пути
	path, err := algorithm.Find(cubes)
	if err != nil {
		fmt.Println(err)
		return
	}

	// вывод пути
	fmt.Println("path", path)
	display("start ", cubes)
	for _, p := range path {
		cubes[p].Rotate(cubes)
		display("("+strconv.Itoa(p)+") -> ", cubes)
	}
}
