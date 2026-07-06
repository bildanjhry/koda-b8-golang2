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

func main() {
	task1()
}
