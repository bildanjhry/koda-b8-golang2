package main

import (
	"fmt"
)

func task1() {

	// declares all needed structs
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

func task4() {

	// declares all needed structs
	type isSt struct {
		is string
	}
	type fruitsSt struct {
		fruit isSt
	}
	type fav struct {
		favorite [4]fruitsSt
	}

	// the values of all struct of fruits
	isApple := isSt{
		is: "Apple",
	}
	isDurian := isSt{
		is: "Durian",
	}
	isManggo := isSt{
		is: "Manggo",
	}
	isPineapple := isSt{
		is: "Pineapple",
	}

	// the structs of all fruits
	apple := fruitsSt{
		fruit: isApple,
	}
	manggo := fruitsSt{
		fruit: isManggo,
	}
	durian := fruitsSt{
		fruit: isDurian,
	}
	pineapple := fruitsSt{
		fruit: isPineapple,
	}

	favorite := fav{
		favorite: [4]fruitsSt{manggo, durian, pineapple, apple},
	}

	my := []fav{favorite}

	fmt.Println(my[0].favorite[3].fruit.is)
}

func main() {
	task1()
	task2()
	task5()
	task4()
}
