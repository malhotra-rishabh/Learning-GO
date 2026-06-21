package main

import "fmt"

type Coordinates struct {
	lat, long float64
}

func main() {
	m := make(map[string]Coordinates)
	m["Home"] = Coordinates{
		lat:  65.23237,
		long: 21.23233,
	}
	m["Work"] = Coordinates{
		lat:  -98.23237,
		long: 32.23233,
	}
	fmt.Println(m)

	fmt.Println("Map value retrivel: ")
	home := m["Home"]
	fmt.Println("Home: ", home)

	fmt.Println("Update key in map: ")
	m["Home"] = Coordinates{
		lat:  -87.8276,
		long: 1.92786,
	}
	fmt.Println(m)
	fmt.Println(home) // home will remain the same as it was before the update, because it is a copy of the value in the map, not a reference to it.

	fmt.Println("Delete element in map: ")
	delete(m, "Home")
	fmt.Println(m)

	fmt.Println("Check if key is present in map: ")
	elem, ok := m["Home"]
	fmt.Println(elem, ok) // if not present, ok is false and elem is 0
	elem, ok = m["Work"]
	fmt.Println(elem, ok) // if present, ok is true and elem is the value
}
