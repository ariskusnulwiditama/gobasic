package test

import (
	"fmt"
	"testing"
)

type somStruct struct {
	name   string
	rate   float64
	active bool
}

func changeSomStruct(s somStruct) {
	s.name = "Jane"
}
func TestNewsomStruct(t *testing.T) {
	s := somStruct{"John", 0.5, true}
	t.Log(s)
	changeSomStruct(s)
	t.Log(s)
}

type Number interface{
	~int | ~float64
}

type TypeStringBool interface{
	~string | ~bool
}

func StringBool[T TypeStringBool](a T) T {
	return a
}

// implementasi stringbool
func StringBool2[T TypeStringBool](a T) {
 // do something

}
func Add2[T Number](a T, b T) T {
	return a + b
}

func TestAdd2(t *testing.T) {
	t.Log(Add2(1, 2))
	t.Log(Add2(1.1, 2.2))
}

func TestStringBool(t *testing.T) {
	t.Log(StringBool("hello"))
	t.Log(StringBool(true))
}

// two sum
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
		// [2:0], [11:1], [7:2], [0:3]
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	t.Log(twoSum([]int{2, 11, 7,0 , 1, 4, 11, 15}, 13))
	t.Log(twoSum([]int{3, 2, 4, 1,2,4}, 6))
	// [3:0], [4:1], [2:2], [5:3], [4:4], [2:5]
}

func printNumbers(numbers []int) {
    for _, num := range numbers {
        fmt.Println(num)
    }
}

func TestPrintNumbers(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	printNumbers(slice)
}