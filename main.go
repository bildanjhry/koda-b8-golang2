package main

import "fmt"

func task1() {
	type str struct {
		the string
	}
	type strbest struct {
		best string
	}
	type strthe struct {
		the strbest
	}
	type strSt struct {
		are strthe
	}

	best := strbest{
		best: "Koda",
	}

	the := strthe{
		the: best,
	}

	we := strSt{
		are: the,
	}

	fmt.Println(we.are.the.best)
}

func task2() {
	type str struct {
		world string
	}

	hello := str{
		world: "Hello World",
	}

	fmt.Println(hello.world)
}

func task5() {
	nums := [3]int{4, 20, 12}
	fmt.Println(nums[1] + nums[2])
}

func main() {
	task1()
	task2()
	task5()
}
