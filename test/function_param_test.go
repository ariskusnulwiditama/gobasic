package test

import (
	"fmt"
	"testing"
)

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	
	result := fmt.Sprintf("Hello %s ", filter(name))
	fmt.Println(result)
}

func spamFilter(name string) string {
	swearWords := []string{"anjing", "bangsat", "kontol", "memek", "pepek", "ngentot", "ngewe", "ngentod", "ngentotin", "ngentodin", "ngentotan", "ngentodan", "ngewe", "ngewean", "ngewek", "ngewekan", "ngewekin", "ngewekin", "ngewekan", "ngewekan"}
	for _, sw := range swearWords {
		if name == sw {
			return "..."
		} 
	}
	return name
}

func TestSayHelloWithFilter(t *testing.T) {
	sayHelloWithFilter("kontol", spamFilter)
}

func TestSayHelloWithFilter2(t *testing.T) {
	sayHelloWithFilter("seffy", spamFilter)
}