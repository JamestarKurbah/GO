package main

import "fmt"

func main() {
	StudentAge := make(map[string]int)
	StudentAge["Aryan"] = 23
	StudentAge["Monika"] = 21
	StudentAge["Johan"] = 27
	StudentAge["Kritika"] = 22

	fmt.Println(StudentAge)
	fmt.Println(len(StudentAge))

	// map inside a map
	superhero := map[string]map[string]string{
		"Superman": map[string]string{
			"Realname": "Clark Kent",
			"City":     "Metropolis",
		},
		"BatMan": map[string]string{
			"Realname": "Bruce Wayne",
			"City":     "Gotham City",
		},
	}
	if temp, hero := superhero["BatMan"]; hero {
		fmt.Println(temp["Realname"], temp["City"])
	}
}
