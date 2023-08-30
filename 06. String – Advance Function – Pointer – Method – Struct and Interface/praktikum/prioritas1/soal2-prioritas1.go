package main

import (
	"fmt"

)

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	scoreall := 0
	for _ , v := range s.score{
		scoreall += v
	}
	rata2 := float64(scoreall) / float64(len(s.score))
	return rata2
}

func (s Student) Min() (min int, name string) {
	min = s.score[0]
	name = s.name[0]
	for i, score := range s.score {
		if score < min {
			min = score
			name = s.name[i]
		}
	}

	return min, name
}

func (s Student) Max() (max int, name string) {
	max = s.score[0]
	name = s.name[0]
	for i, score := range s.score {
		if score > max {
			max = score
			name = s.name[i]
		}
	}
	return max, name
}

func main() {
	var a = Student{}

	for i := 0; i < 5; i++ {
		var name string
		fmt.Print("Input Student's Name	: ")
		fmt.Scan(&name)
		a.name = append(a.name, name)
		var score int
		fmt.Print("Input " + name + " Score	: ")
		fmt.Scan(&score)
		a.score = append(a.score, score)
	}

	fmt.Printf("\n\nAverage Score Students is : %.f\n",a.Average())

	scoreMax, nameMax := a.Max()
	fmt.Println("Max Score Students is	: "+nameMax+" (", scoreMax, ")")

	scoreMin, nameMin := a.Min()
	fmt.Println("Min Score Students is	: "+nameMin+" (", scoreMin, ")")
}