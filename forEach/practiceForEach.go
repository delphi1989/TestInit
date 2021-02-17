package forEach

import "fmt"

func TestForEach()  {
	for i := 0; i < 5 ; i++ {
		fmt.Println("it's allowed")
	}

	for i := 0; i < 5 ; i = i + 2 {
		fmt.Println("it's allowed too")
	}

	// 同時宣告兩個變數
	// 第三個 argument 只能是個 statement
	// 而 i++ 與 j++ 都是個 statement
	// 因此不能宣告為 i, j = i++, j++
	for i, j := 0, 0; i < 5; i, j = i + 1, j + 2 {
		fmt.Println("it's allowed too")
	}
}

func TestForSyntax() {
	array := []int {
		1, 2, 3,
	}

	for k, v := range array {
		fmt.Println(k, v)
	}
}

func TestWhile() {
	//i := 0
	//
	//for i < 5 {
	//	i++
	//	fmt.Println(i)
	//	break
	//}

	Loop:
		for i := 1; i < 5; i++ {
			fmt.Println("i=", i)
			TestLoop:
				for j:= 1; j < 5; j ++ {
					fmt.Println("j=", j)
					for x := 1; x < 5; x ++ {
						fmt.Println("x=", x)
						break TestLoop
					}
					break Loop
				}
		}




}