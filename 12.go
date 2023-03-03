package main

import "fmt"

func listToSet(x []string) {

	set := make(map[string]struct{})

	for _, value := range x {
		set[value] = struct{}{}
	}
	for key, _ := range set {
		fmt.Println(key)
	}
}

func main() {

	massiv := []string{"cat", "cat", "dog", "cat", "tree"}

	listToSet(massiv)

}
