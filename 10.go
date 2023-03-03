package main

import "fmt"

type temperatura struct {
}

// классифицируем нашу температуру от t, t+10. Используем деление 10 и берем целую част в качестве ключа.
// В отрицательных значениях вычитаем 10 и также делим на 10 для получения ключаю,чтобы выполнялось условие
func classification(array []float64) map[int][]float64 {
	mappa := make(map[int]map[int][]float64)
	result := make(map[int][]float64)

	for i := 0; i < len(array); i++ {
		if array[i] > 0 {
			index1 := int(array[i] / float64(10))
			if mappa[index1] == nil {
				mappa[index1] = make(map[int][]float64)
			}
			mappa[index1][index1*10] = append(mappa[index1][index1*10], array[i])
		} else {
			index1 := int((array[i] - 10) / float64(10))
			if mappa[index1] == nil {
				mappa[index1] = make(map[int][]float64)
			}
			mappa[index1][index1*10] = append(mappa[index1][index1*10], array[i])
		}
	}
	for key, _ := range mappa {
		if key < 0 {
			result[(key+1)*10] = mappa[key][key*10]
		} else {
			result[key*10] = mappa[key][key*10]
		}
	}
	return result

}

func main() {
	array := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	result := classification(array)

	for key, value := range result {

		fmt.Println("ключ:", key, "", "значение", value)
	}
}
