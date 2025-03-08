package test

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func FindMin(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, errors.New("Array is empty")
	}
	min := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
	}

	return min, nil
}

func TestFind(t *testing.T) {
	res := []struct {
		name string
		input [] int
		expected int
		expectedErr bool
	} {
		{"PositiveNumbers", []int{3,5,7,2,8}, 2, false},
		{"NegativeNumbers", []int{-3,-5,-7,-2,-8}, -8, false},
		{"EmptyArray", []int{}, 0, true},
	}

	for _, tt := range res {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FindMin(tt.input)
			if (err != nil) != tt.expectedErr {
				t.Errorf("Unexpected error state: got %v, want error=%v", err, tt.expectedErr)
			}
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}

type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}

func (n *TreeNode) Insert(value int) {
	if value <= n.Value {
		if n.Left ==  nil {
			n.Left = &TreeNode{Value: value}
		} else {
			n.Left.Insert(value)
		}
	} else {
		if n.Right ==  nil {
			n.Right = &TreeNode{Value: value}
		} else {
			n.Right.Insert(value)
		}
	}
}

func (n *TreeNode) InOrder() {
	if n == nil {
		return
	}
	n.Left.InOrder()
	fmt.Printf("%d ", n.Value)
	n.Right.InOrder()
}

func (n *TreeNode) PreOrder() {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.Value)
	n.Left.PreOrder()
	n.Right.PreOrder()
}

func (n *TreeNode) PostOrder() {
	if n == nil {
		return
	}
	n.Left.PostOrder()
	n.Right.PostOrder()
	fmt.Printf("%d ", n.Value)
}

func TestBinaryTree(t *testing.T) {
	root := &TreeNode{Value: 100}
	root.Insert(50)
	root.Insert(150)
	root.Insert(20)
	root.Insert(70)
	root.Insert(120)
	root.Insert(180)

	fmt.Println("InOrder")
	root.InOrder()
	fmt.Println("")

	fmt.Println("PreOrder")
	root.PreOrder()
	fmt.Println("")

	fmt.Println("PostOrder")
	root.PostOrder()
	fmt.Println("")
}



func generateIdPengajuanSb() (idPengajuanSb string) {
	idPengajuanSb = "SB-" + strconv.Itoa(rand.Intn(100000)) + "-" + time.Now().Format("20060102")
	
	return idPengajuanSb
}
func TestGenerateIdPengajuanSb(t *testing.T) {
	
		id := generateIdPengajuanSb()
	
	fmt.Println(id)
}