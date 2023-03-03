package main

import "fmt"

//функция нахождения общих элементов множества среди set1 и set2
func innerSelect(set1, set2 map[int]struct{}) map[int]struct{} {
	var res = make(map[int]struct{})

	if len(set1) > len(set2) {
		for value := range set1 {
			_, found := set2[value]
			if found {
				res[value] = struct{}{}
			}
		}
	} else {
		for value := range set2 {
			_, found := set1[value]
			if found {
				res[value] = struct{}{}
			}
		}

	}
	return res
}

func main() {

	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})

	for i := 0; i <= 5; i++ {
		set1[i] = struct{}{}
	}

	for i := 3; i < 10; i++ {
		set2[i] = struct{}{}
	}

	res := innerSelect(set1, set2)

	for key, _ := range res {
		fmt.Println("Key:", key)
	}

}
