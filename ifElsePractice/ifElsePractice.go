package ifElsePractice

import (
	"fmt"
)

func TestIfSugarSyntax()  {
	statePopulations := map[string]int{
		"California": 300000,
		"Texas": 200000,
		"Florida": 300000,
		"New York": 100000,
		"Illinois": 3000000,
		"Ohio": 50000,
	}

	// 第一個參數表示 map 的 value
	// 第二個參數表示 bool, 且後面的括號表示 success 才會進入
	if pop, ok := statePopulations["Texas"]; ok {
		fmt.Println(ok == true)
		fmt.Println(pop)
	}
	// 這邊是不能執行的
	//fmt.Println(pop)
}

func TestSimpleGuess() {
	number := 50
	guess := 5

	if guess < 1 {
		fmt.Println("The guess is too low")
	} else if guess > 100 {
		fmt.Println("The guess is too high")
	} else {
		if number == guess {
			fmt.Println("Correct")
		} else {
			fmt.Println("Incorrect")
		}
	}

}

func TestSwitch() {
	switch 3 {
	case 1, 5, 10:
		fmt.Println("one, five, ten")
	case 4, 6:
		fmt.Println("four, six")
	default:
		fmt.Println("none")
	}

	// syntax sugar
	switch i := 2+3; i {
	case 1, 5, 10:
		fmt.Println("one, five, ten")
	case 4, 6:
		fmt.Println("four, six")
	default:
		fmt.Println("none")
	}

	// tagless switch
	// 會允許 overlap, 但根據宣告順序, 只會執行第一個
	i := 10
	switch {
	case i <= 10:
		fmt.Print("小於 10")
		break
	case i <= 20:
		fmt.Print("小於 20")
		break
	default:
		fmt.Print("大於 20")
	}

	// 如果有用到 fallthrough 關鍵字
	// 表示將流程的控制全交由自己掌握
	// 因此, 即使 switch 不符合結果, 也會執行
}

func TestTypeInterfaceSwitch() {
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("I'm int")
		// 如果不想執行下面
		break
		fmt.Println("I'm int")
	case float64:
		fmt.Println("I'm float")
	case string:
		fmt.Println("I'm string")
	// 留意, 下面兩者的比較是會根據 capacity 比較的!!
	case [2]int:
		fmt.Println("I'm array with capacity 2")
	case [3]int:
		fmt.Println("I'm array with capacity 3")
	default:
		fmt.Println("I'm nothing")
	}
}


func TestOr() {
	var a byte = 1 << 1
	var b byte = 1 << 2
	var c byte = 1 << 3
	var d byte = 1 << 3

	var testRole byte = a | b | c
	fmt.Println(testRole & a == d)
}