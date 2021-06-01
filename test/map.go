package test

import "fmt"

type Person struct {
	Name string `json:"name"`
}

func TestMap()  {
	var person Person
	person = Person{Name: "lisa"}

	var mapData = make(map[string]interface{})
	mapData["key1"] = "value1"
	mapData["key2"] = 222
	mapData["key3"] = person

	fmt.Println(mapData)
}
