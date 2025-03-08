package test

import "testing"

type Student struct {
	Name  string
	Age   int
	Marks []int
}

func NewStudent(name string, age int, marks []int) *Student {
	return &Student{
		Name: name,
		Age:  age,
		Marks: marks,
	}
}

func (s *Student) GetMarks() int {
	sum := 0
	for _, v := range s.Marks {
		sum += v
	}
	return sum
}

func (s *Student) GetAverage() float64 {
	return float64(s.GetMarks()) / float64(len(s.Marks))
}

func (s *Student) GetGrade() string {
	avg := s.GetAverage()
	if avg >= 90 {
		return "A"
	} else if avg >= 80 {
		return "B"
	} else if avg >= 70 {
		return "C"
	} else if avg >= 60 {
		return "D"
	} else {
		return "F"
	}
}

func TestStudent(t *testing.T) {
	s := new (Student)
	s.Name = "John"
	s.Age = 20
	s.Marks = []int{90, 80, 70, 60, 50}
	

	t.Log("Name: ", s.Name)
	t.Log("Age: ", s.Age)
	t.Log("Marks: ", s.Marks)
	t.Log("Average: ", s.GetAverage())
}