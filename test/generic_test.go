package test

import (
	"cmp"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}

	return -1
}

func TestGeneric(t *testing.T) {
	numbers := []int{3, 5, 7, 9, 11}

	t.Log(Index(numbers, 7))
	t.Log(Index(numbers, 12))
	t.Log(Index([]string{"apple", "banana", "cherry"}, "banana"))
	t.Log(Index([]string{"apple", "banana", "cherry"}, "kiwi"))
	t.Log(Index([]float64{3.14, 2.71, 1.61}, 2.71))
	t.Log(Index([]float64{3.14, 2.71, 1.61}, 1.41))
	t.Log(Index([]string{"apple", "banana", "cherry"}, "banana"))
	t.Log(Index([]string{"apple", "banana", "cherry"}, "kiwi"))
	t.Log(Index([]float64{3.14, 2.71, 1.61}, 2.71))
	t.Log(Index([]float64{3.14, 2.71, 1.61}, 1.41))
}

// contoh generic function untuk map
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

func TestMap(t *testing.T) {
	numbers := []int{3, 5, 7, 9, 11}
	t.Log(Map(numbers, func(x int) int { return x * x }))
	t.Log(Map(numbers, func(x int) string { return fmt.Sprintf("%d", x) }))
	t.Log(Map(numbers, func(x int) float64 { return float64(x) }))
	t.Log(Map([]string{"apple", "banana", "cherry"}, func(x string) int { return len(x) }))
	t.Log(Map([]string{"apple", "banana", "cherry"}, func(x string) float64 { return float64(len(x)) }))
	t.Log(Map([]string{"apple", "banana", "cherry"}, func(x string) string { return strings.ToUpper(x) }))
}

// contoh generic function dengan param map[string]any menjadi map[string]string atau map[string]int

func MapString2[T any](m map[string]T, f func(T) string, g func(string) int) (map[string]string, map[string]int) {
	result1 := make(map[string]string)
	result2 := make(map[string]int)
	for k, v := range m {
		result1[k] = f(v)
		result2[k] = g(k)
	}
	return result1, result2
}

func Min[T cmp.Ordered]( x, y T) T {
	if x < y {
		return x
	}
	return y
}

func TestMin(t *testing.T) {
	t.Log(Min(3, 5))
	t.Log(Min(3.14, 2.71))
	t.Log(Min("apple", "banana"))
}

func Sum[T ~int | ~float64](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

func TestSum(t *testing.T) {
	result := Sum([]int{3, 5, 7, 9, 11})
	t.Log(result)
	tipe := reflect.TypeOf(result)
	t.Log(tipe)
	t.Log(Sum([]int{3, 5, 7, 9, 11}))
	t.Log(Sum([]float64{3.14, 2.71, 1.61}))
}

type Num interface {
	~int | ~int64 | ~float32 | ~float64
}

func Sum2[T Num](s []T) T {
	var sum T
	for _, v := range s {
		sum += v
	}
	return sum
}

type Person struct {
	Name string
	Age  int
}

func PrintPerson[T any] (val T) {
	fmt.Println(val)
}

func CheckAge[T any](val T) bool {
	switch v := any(val).(type) {
	case Person:
		return v.Age > 18
	default:
		return false
	}
}

func TestPrintPerson(t *testing.T) {
	PrintPerson(Person{"John Doe", 30})
}

func TestCheckAge(t *testing.T) {
	t.Log(CheckAge(Person{"John Doe", 30}))
	t.Log(CheckAge(Person{"Jane Doe", 15}))
}

type HasAge interface {
	GetAge() int
}

func CheckAge2[T HasAge](val T) bool {
	return val.GetAge() > 18
}
func(p *Person) GetAge() int {
	return p.Age
}
func TestCheckAge2(t *testing.T) {
	t.Log(CheckAge2(&Person{"John Doe", 30}))
	t.Log(CheckAge2(&Person{"Jane Doe", 15}))
}

func SumNumber[K comparable, V Num](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}

	return sum
}

func TestSumNumber(t *testing.T) {
	t.Log(SumNumber(map[int]int{1: 3, 2: 5, 3: 7, 4: 9, 5: 11}))
	t.Log(SumNumber(map[int]int{1: 3, 2: 5, 3: 7, 4: 9, 5: 11}))
}