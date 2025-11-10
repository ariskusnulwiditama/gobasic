// make golang unit test
package basic

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	println("TestHello")
}


func TestMap(t *testing.T) {
	var myMap map[string] any 
	/* myMap["name"] = "John"
	myMap["age"] = 30
	myMap["city"] = "New York"
	myMap["country"] = "USA" */

	fmt.Println(myMap)

	 for key, value := range myMap {
		fmt.Println(key+" : ", value)
	} 

	myMap2 := map[string] any {
		"name": "John",
		"age": 30,
		"city": "New York",
		"country": "USA",

	}

	fmt.Println(myMap2["name"])
}

type User struct {
	name string
	age int
	address string
}

// make constructor
func NewUser(name string, age int, address string) *User {
	return &User{name, age, address}
}

func (u *User) String() string {
	return fmt.Sprintf("Name: %s, Age: %d, Address: %s", u.name, u.age, u.address)
}

func TestStruct(t *testing.T) {
	user := NewUser("John", 30, "New York")
	t.Log(user.String())
}

func TestUserString(t *testing.T) {
	user := User{}

	t.Log(user.String())
}

func TestMap2(t *testing.T) {
	myMap := map[string] any {
		"name": "John",
		"age": 30,
		"city": "New York",
		"country": "USA",
	}

	for key, value := range myMap {
		fmt.Println(key+" : ", value)
	}
}

func TestSlice(t *testing.T) {
	var mySlice []int
	mySlice = append(mySlice, 10)
	mySlice = append(mySlice, 20)
	mySlice = append(mySlice, 30)
	
	fmt.Println(mySlice)
}

type Queue struct {
	elements []int
}

func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

func (q *Queue) Dequeue() int {
	if len(q.elements) == 0 {
		return -1
	}

	value := q.elements[0]
	q.elements = q.elements[:1]
	return value

}

func TestSet(t *testing.T) {
	set := make(map[string] struct{})
	set["apple"] =struct{}{}
	set["banana"] = struct{}{}
	set["orange"] = struct{}{}
	set["apple"] = struct{}{}
	t.Log(set)
	for key := range set {
		t.Log(key)
	}

	if _, exists := set["apple"]; exists {
		t.Log("apple exists")
	}
}


























