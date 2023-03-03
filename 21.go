package main

import "fmt"

// интерфейс символов
type simbol interface {
	printlatin()
}

// структура латинские буквы
type latin struct {
}

// метод которые выводит латиницу
func (l latin) printlatin() {
	fmt.Println("latin")
}

// структура русские буквы
type rus struct{}

// метод которые выводит кириллицу
func (r rus) printrus() {
	fmt.Println("кириллица")
}

// адаптер,который реализует simbol
type adaptertranslation struct {
	r *rus
}

// вызывает printrus() у rus и реализует метод интерфейса printlatin()
func (ad adaptertranslation) printlatin() {
	ad.r.printrus()
}

// переводчик,который может понимать что за символы
type interpreter struct {
}

// перевод текста
func (i *interpreter) translation(s simbol) {
	s.printlatin()
}

func main() {
	interpreter := &interpreter{}
	latinWord := &latin{}
	interpreter.translation(latinWord)
	rusWord := &rus{}
	adaptertranslation := &adaptertranslation{r: rusWord}
	interpreter.translation(adaptertranslation)

}
