package main

import (
	"fmt"
	"sort"
)

func main() {
	lockTest()
}
func lockTest() {

	map1 := make(map[string]int)
	map1["3"] = 11
	map1["5"] = 5
	map1["1"] = 7
	map1["4"] = 23
	map1["2"] = 42

	//.sortMap2(map1)
	//fmt.Println()
	sortMap(map1)
}

//根据key排序
func sortMap2(mp map[string]int) {
	//1.将map1的key放到切片中
	var newMap = make([]string, 0)
	for k, _ := range mp {
		newMap = append(newMap, k)
	}
	//2.对切片排序
	sort.Strings(newMap)
	//3.遍历切片，然后按key来输出map的值
	for _, v := range newMap {
		fmt.Printf("根据key排序后的新集合为:[%v]=%v \n", v, mp[v])
	}
}

//根据value排序
func sortMap(mp map[string]int) {
	var newMp = make([]int, 0)
	var newMpKey = make([]string, 0)
	for oldk, v := range mp {
		newMp = append(newMp, v)
		newMpKey = append(newMpKey, oldk)
	}
	sort.Ints(newMp)
	fmt.Println("newMp=", newMp)

}
