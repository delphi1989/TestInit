package array

import "fmt"

const TestValue = "測試"
const testValue2 = 1

func TestArray() {
	var students = [...]string{
		"Delphi", "Blair", "Tom", "Mark",
	}
	students[0] = "Delphi"
	students[1] = "Blair"
	students[2] = "Tom2"
	students[3] = "Mark2"
	// 下面會有 compile error
	//students[4] = "Mark2"
	fmt.Println(students)
}


func TestArray2() {
	var teachers = [...]string{
		"Delphi", "Blair", "Tom", "Mark",
	}

	fmt.Println(len(teachers))
}

/*
 測試一下文件
 */
func TestTwoDimensionArray() {
	var identityMatrix [3][3] int
	identityMatrix[0] = [3]int{ 0, 0, 1}
	identityMatrix[1] = [3]int{ 0, 1, 0}
	identityMatrix[2] = [3]int{ 1, 0, 0}

	fmt.Println(identityMatrix)
}

/**
 在複製 Array 時, 實際上是真的生成一個新的 array
 */
func TestCopyArrayButLiteral() {
	a := [3]int{1, 2, 3}
	b := a
	b[0] = 4

	fmt.Print(a)
	fmt.Print(b)
}

func TestPreventCopyWholeNewArray() {
	a := [3]int{1, 2, 3}
	b := &a
	b[0] = 4
	fmt.Print(a)
	// 印出來的 & 符號是有意義的
	fmt.Print(b)
}

func TestSlice() {
	// 建立一個 Slice
	// Slice 是個 projection of underlying array
	// 因此 slice 是個 reference data, 換句話說, 不需要上面的指針
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 如果是 [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} 則表示是建立 array
	//b := a
	//b[0] = 4

	// slice 第二個參數是 exclusive
	fmt.Println(a[:]) // 所有的 els
	fmt.Println(a[3:]) // 從 index 3 開始到最後
	fmt.Println(a[:6]) // 從 index 0 開始到 index 5
	fmt.Println(a[3:6]) // 從 index 3 開始到 index 5

	b := a[3:]
	b[0] = 999
	fmt.Println(b)
	// 上面 b 變動的結果也會影響到 a
	fmt.Println(a)
}

func TestMakeFeature() {
	// slice 的長度可以是動態的
	a := make([]int, 3, 100)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	// 做這件事時, 若宣告的 slice 沒有初始化 capacity, 則 golang 做的事情是建立一個全新的 Array, 把資料 copy 過去
	// 且 size 會是 double 的
	// len 表示有幾個元素的長度
	// cap 表示總共的空間
	a = append(a, 1)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))


	// spread slice, 展開 slice
	a = append(a, []int{ 7, 8, 9, 10, 11}...)
}

func TestPopElFromSlice() {
	a := []int{1, 2, 3, 4, 5}
	// 去掉中間的元素 3
	// a 的前兩個元素 append 尾端兩個元素
	b := append(a[:2], a[3:]...)
	fmt.Println(b)

}