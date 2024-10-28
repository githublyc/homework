package main

import "fmt"

func main() {
	//tst := []any{4, 5, 9, 22}
	tst3 := []int{4, 5, 9, 22}
	//ans, err := Yaoqiu1(tst, 4)
	//fmt.Printf("%v\n%s", ans, err)
	//ans, err := Yaoqiu2(tst, 2)
	//fmt.Printf("%v\n%s", ans, err)
	//ans := Yaoqiu3(tst3, 2)
	//fmt.Printf("%v\n", ans)
	ans := Yaoqiu4(tst3, 2)
	fmt.Printf("%v\n", ans)
	fmt.Printf("cap: %d\n", cap(ans))
}
