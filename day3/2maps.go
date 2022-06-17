package main

import "fmt"

func maps() {
	fmt.Println("A map is a data structure that provides you with an unordered collection of key/value pairs (maps are also sometimes called associative arrays in Php, hash tables in Java, or dictionaries in Python). Maps are used to look up a value by its associated key. You store values into the map based on a key.")

	//map declarations

	var m map[string]int //empty declaration
	//m["a"] = 1           //cant assing value before allocating space
	fmt.Printf("Value %v \n", m)
	var m1 map[string]int = make(map[string]int) //allocate space
	m1["a"] = 2
	fmt.Printf("Value %v \n", m1)
	m2 := map[string]int{"a": 1, "b": 2} // map literal
	fmt.Printf("Value %v %v\n", m2, m2["a"])

	//adding items

	m2["c"] = 4
	m2["b"] = 3 //update value
	delete(m2, "a")
	fmt.Printf("Value %v %v\n", m2, m2["c"])

	//loop over map

	for k, v := range m2 {
		fmt.Printf("m2[%s]=%v \n", k, v)
	}

	//truncate table

	for k, _ := range m2 {
		delete(m2, k)
	}
	fmt.Printf("Value %v\n", m2)

	a := map[int]int{1: 1, 2: 2, 3: 3}
	b := map[int]int{4: 4, 5: 5}
	for k, v := range b {
		a[k] = v
	}
	fmt.Printf("Value %v\n", a)

}

func main() {
	maps()
}
