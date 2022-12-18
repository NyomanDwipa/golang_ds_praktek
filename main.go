package main

import "fmt"

func doLoop(any interface{}) { // interface{} menerima semua nilai 
	switch v := any.(type) { // get tipe data dari interface{}
	case []string: // kalau slice -> []string
		for key, value := range v {
			fmt.Println("ini key : ", key+1)
			fmt.Println("ini value : ", value)
		}
		fmt.Println()
	case map[string]string: // kalau map -> map[string]string
		for key, value := range v {
			fmt.Println("ini key : ", key)
			fmt.Println("ini value : ", value)
		}
		fmt.Println()
	case [3]string:
		for key, value := range v { // kalau array -> [3]string
			fmt.Println("ini key : ", key+1)
			fmt.Println("ini value : ", value)
		}
		fmt.Println()
	}
}

func main() {
	var names []string // slice
	names = append(names, "afandy")
	names = append(names, "alexia")
	names = append(names, "asford")

	doLoop(names)

	class := [3]string{ // array
		"Kelas A", "Kelas B", "Kelas C",
	}

	doLoop(class)

	mk := make(map[string]string) // map
	mk["kelas A"] = "Matematika"
	mk["kelas B"] = "IPA"
	mk["kelas C"] = "Penjas"

	doLoop(mk)
}