package main

import (
	"fmt"
)

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

func task1() {
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

type strWorld struct {
	world string
}

func task2() {
	hello := strWorld{
		world: "Hello World",
	}
	fmt.Println(hello.world)
}

// func task3() {
// 	type techSt struct {
// 		academy string
// 	}

// 	type manSt struct {
// 		tech techSt
// 	}

// 	type structArr struct {
// 		man manSt
// 	}

// 	type str struct {
// 		tech techSt
// 	}

// 	tech := techSt {
// 		academy: "Tech Academy",
// 	}

// 	man := manSt {
// 		tech: tech,
// 	}

//  structA:= []structArr{{}, {}}

//  type objStr struct {
// 	str [][][]kodaSt
//  }

//  strArr = []objStr {{}, {}, {}, {}}

//  var objSlice = [][]str {null, null, {structA[0].tech.academy},}

//  var strSlice = [][][]str{{}, {}, {null, structA}}

//  obj := objStr {
// 	str: structA,
//  }

// }

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

func task4() {
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

func task5() {
	nums := [3]int{4, 20, 12}
	fmt.Println(nums[1] + nums[2])
}

func main() {
	task1()
	task2()
	task4()
	task5()
}
