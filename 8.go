package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n int64 = 9223372036854775807
	var i int64 = 1
	var res1 int64
	var res2 int64

	// для того чтобы поменять бит на 1 , нужна маска вида 000..i...000 и операция побитового или ,тем самым мы поменяем бит только в нужном месте
	res1 = n | 1<<(i-1)

	// для того чтобы поменять бит на 0 , нужна маска вида 111..i...111 и операция побитового и ,тем самым мы поменяем бит только в нужном месте
	res2 = res1 & ^(1 << (i - 1))

	fmt.Println(n)
	fmt.Println(strconv.FormatInt(n, 10) + ":" + strconv.FormatInt(res1, 10))
	fmt.Println(strconv.FormatInt(res1, 10) + ":" + strconv.FormatInt(res2, 10))
}
