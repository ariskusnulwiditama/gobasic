package test

import (
	"errors"
	"strings"
	"testing"
)

func TestMap1(t *testing.T) {
	person := map[string]string{}
	person["name"] = "andi"
	person["title"] = "programmer"

	t.Log(person)
	t.Log(person["name"])
	t.Log(person["title"])

	for key, value := range person {
		t.Log(key, value)
	}
}

func TestMap2(t *testing.T) {
	myMap := map[string][]string{}
	myMap["name"] = []string{"andi", "budi", "caca"}
	myMap["title"] = []string{"programmer", "designer", "manager"}
	
	map2 := make(map[string]int, 0)
	for i := 0; i < 7; i++ {
		map2["name"]++
	}
	
	t.Log(map2)
}

func TestMapCounter(t *testing.T) {
	words := []string{"apple", "banana", "cherry", "apple", "cherry", "cherry", "banana", "apple", "cherry"}
	counter := map[string]int{}
	for _, word := range words {
		counter[word]++
	}

	t.Log(counter)
}

func FindDuplicatePerson(people []Person) (map[string]int, map[int]int) {
	seen := make(map[Person]int)  //counter
	nameDuplicate := make(map[string]int)
	ageDuplicate := make(map[int]int)

	for _, person := range people {
		seen[person]++
		if seen[person] > 1 {
			nameDuplicate[person.Name]++
			ageDuplicate[person.Age]++
		}
	}

	return nameDuplicate, ageDuplicate
}

func FindDuplicateNumber(numbers []int) (result []map[int]int) {
	seen := make(map[int]int)
	sliceDuplicate := make(map[int]int)

	for _, num := range numbers {
		seen[num]++
		if seen[num] > 1 {
			sliceDuplicate[num]++
		}

		result = append(result, sliceDuplicate)
	}

	return result
}

func TestFindDuplicateNumber(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	result := FindDuplicateNumber(numbers)
	/* for _, v := range result {
		if len(v) > 0 {
			t.Log(v)
		}
	} */
	t.Log(result)
}

func TestFindDuplicatePerson(t *testing.T) {
	people := []Person{
		{"andi", 20},
		{"budi", 20},
		{"caca", 20},
		{"didi", 20},
		{"andi", 20},
		{"budi", 20},
		{"caca", 20},
		{"didi", 20},
	}

	nameDuplicate, ageDuplicate := FindDuplicatePerson(people)
	t.Log(nameDuplicate)
	t.Log(ageDuplicate)
}

func CountFrequency(nums []int) map[int]int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	return freq
}

func TestCountFrequency(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	freq := CountFrequency(nums)
	t.Log(freq)
}

func Add[T Num](x, y T) T {
	return x+y
}

func Divided[T Num](x, y T) (T, error) {
	if y == 0 {
		return 0, errors.New("cannot divided by zero")
	}

	return x/y, nil
}

func SetInt[T Num](x T) T {
	return x
}

func TestSetInt(t *testing.T) {
	t.Log(SetInt(10))
	t.Log(SetInt(10.5))
	
}

func TestFindDuplicate(t *testing.T) {
	
	maping := make(map[string]int)
	maping["andi"] = 1
	maping["budi"] = 2
	maping["caca"] = 3
	for key, value := range maping {
		// find duplicate
		if value > 1 {
			t.Log("Duplicate found:", key, value)
		} else {
			t.Log("No duplicate found:", key, value)
		}
	}

}	

func CountWords(text string) map[string]int  {
	words := strings.Fields(text)
	frequency := make(map[string]int)

	for _, word := range words {
		frequency[word]++
	}

	return frequency
}

func TestCountWords(t *testing.T) {
	text := "hello world hello universe hello world"
	frequency := CountWords(text)
	t.Log(frequency)
}

func HasDuplicate(nums []int) bool {
	seen := make(map[int]bool)
	for _, num := range nums {
		if seen[num] {
			return true
		}
		seen[num] = true
	}

	return false
}

func TestHasDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 1}
	if HasDuplicate(nums) {
		t.Log("Duplicate found")
	} else {
		t.Log("No duplicate found")
	}	
}

func HasDuplicateIn(nums []int) map[int]int {
	// yang mana angka dan posisi indexnya yang duplikat
	seen := make(map[int]int)
	duplicates := make(map[int]int)
	for index, num := range nums {
		if _, exists := seen[num]; exists {
			duplicates[num] = index // menyimpan index terakhir dari duplikat
		} else {
			seen[num] = index // menyimpan index pertama dari angka tersebut
		}
	}

	if len(duplicates) == 0 {
		return map[int]int{-1: -1} // mengembalikan -1 jika tidak ada duplikat
	}

	return duplicates
}

func TestHasDuplicateIn(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 1, 1}
	duplicates := HasDuplicateIn(nums)
	if len(duplicates) > 0 && duplicates[0] != -1 {
		t.Log("Duplicate found at index:", duplicates)
	} else {
		t.Log("No duplicate found")
	}
}

func FindDuplicatesWithIndices(nums []int) (duplicates []int, duplicatesIndices map[int][]int) {
	freq := make(map[int]int)
	indices := make(map[int][]int)
	duplicatesIndices = make(map[int][]int)

	for i, n := range nums {
		freq[n]++
		indices[n] = append(indices[n], i)

		if freq[n] > 1 {
			duplicates = append(duplicates, n)
			duplicatesIndices[n] = indices[n]
		}
	}

	return duplicates, duplicatesIndices
}

func TestFindDuplicatesWithIndices(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	duplicates, indices := FindDuplicatesWithIndices(nums)
	t.Log("Duplicates:", duplicates)
	t.Log("Indices of duplicates:", indices)
}

func TestFloatx(t *testing.T) {

	myMap := map[string]any{}
	myMap["name"] = "andi"
	myMap["age"] = 20
	myMap["height"] = true
	myMap["weight"] = 70.5
	t.Log(Divided(10.5, 0))
}



