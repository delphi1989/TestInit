package mapp

import "fmt"

/*
  map 的建立, key 需要實際型別
 */
func TestMapEquality() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 300000,
		"Texas": 200000,
		"Florida": 300000,
		"New York": 100000,
		"Illinois": 3000000,
		"Ohio": 50000,
	}

	// 回傳的 map 順序是 not guarantee
	statePopulations["Georgia"] = 10310371
	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations["Georgia"])

	// 第二個回傳的參數代表確認 map 中是否存在此 key
	// 如果存在會是 true, 否則 false
	pop, ok := statePopulations["Ohi"]
	fmt.Println(pop, ok)
	// 如果單純想確認, 第一個參數可以用 Blank identifier 來替代
	_, check := statePopulations["Ohio"]
	fmt.Println(check)

	// 確認 map 長度
	fmt.Println(len(statePopulations))

	// map 是 pass by value, 因此會影響原始的 map
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(len(sp))
	fmt.Println(len(statePopulations))



}