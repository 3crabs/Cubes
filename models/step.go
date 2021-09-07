package models

//Step шаг алгоритма
type Step struct {
	Cubes          []Cube // положение кубиков
	NextCubeNumber int    // номер кубика который будем поворачивать
	Path           []int  // путь - список кубиков которые уже повернули
}

//IsWin проверка комбинации кубиков на выигрышную
func (s Step) IsWin() bool {
	for i := 0; i < len(s.Cubes)-1; i++ {
		if s.Cubes[i].State != s.Cubes[i+1].State {
			return false
		}
	}
	return true
}
