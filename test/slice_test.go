package test

import "testing"

func TestSlice1(t *testing.T) {
	arr := []string{"a", "b", "c", "d", "e"}
	sl := arr[2:]
	t.Log(sl)
}

func TestSlice2(t *testing.T) {
	hari := [...]string{"minggu", "senin", "selasa", "rabu", "kamis", "jumat", "sabtu"}
	t.Log(hari)
	hariBaru := hari[5:]
	t.Log(hariBaru)
	hariBaru[0] = "JUMAT"
	hariBaru[1] = "SABTU"

	t.Log(hariBaru)
}

func TestSlice3(t *testing.T) {
	newSlice := make([]string, 0, 5)
	newSlice = append(newSlice, "andi")
	newSlice = append(newSlice, "budi")
	newSlice = append(newSlice, "caca")
	t.Log(len(newSlice))
	t.Log(cap(newSlice))

	fromSlice := newSlice[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)
	t.Log(toSlice)
	t.Log(fromSlice)
}

func TestSlice4(t *testing.T) {
	mySlice := make([]int, 0, 5)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 3)
	mySlice = append(mySlice, 4)
	mySlice = append(mySlice, 5)
	mySlice = append(mySlice, 6)
	mySlice = append(mySlice, 7)

	for _, v := range mySlice {
		t.Log(v) 	
	}
}

func TestThreeDimensionSlice(t *testing.T) {
	// var threeDimensionSlice [][][]int
	threeDimensionSlice := make([][][]int, 0, 5)

	for i := 0; i < 5; i++ {
		threeDimensionSlice = append(threeDimensionSlice, make([][]int, 0, 5))
		for j := 0; j < 5; j++ {
			threeDimensionSlice[i] = append(threeDimensionSlice[i], make([]int, 0, 5))
			for k := 0; k < 5; k++ {
				threeDimensionSlice[i][j] = append(threeDimensionSlice[i][j], k)
			}
		}
	}

	t.Log(threeDimensionSlice)
}