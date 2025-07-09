package main

import (
	"fmt"
	"strings"

	"io.github.sowen.datastructure/search"
)

type Student struct {
	Name string
}

func isSameOneStudent(one interface{}, other interface{}) bool {
	if one == other {
		return true
	}

	oneStudent, ok := one.(Student)
	if ok {
		if otherStudent, ok2 := other.(Student); ok2 {
			return strings.EqualFold(oneStudent.Name, otherStudent.Name)
		}
	}

	return false

}

func main() {
	var data []int
	data = append(data, 24)
	data = append(data, 18)
	data = append(data, 11)
	data = append(data, 9)
	data = append(data, 16)
	data = append(data, 66)
	data = append(data, 32)
	data = append(data, 4)

	res := search.LinearSearchInt(data, 16)
	fmt.Println(res)

	var students []interface{}
	students = append(students, Student{"Owen"})
	students = append(students, Student{"Vava"})

	res2 := search.LinearSearch(students, Student{"vava"})
	fmt.Println(res2)

	res3 := search.LinearSearchWith(isSameOneStudent, students, Student{"vava"})
	fmt.Println(res3)
}
