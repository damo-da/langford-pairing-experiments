package main

import (
"fmt"
)


func getEmptyIndices(list []int) []int{
	return []int{}

}

func debugPrint(args... interface{}){
	for _, element := range args {
		fmt.Print(element)
		fmt.Print("\t")
	}
	fmt.Println()
}

func solve(number int, cur_list []int){
	if cur_list == nil {
		cur_list = make([]int, number*2)
	}

	debugPrint(cur_list, "batman", 1)

	//emptyIndices := getEmptyIndices(cur_list)

}
func main(){
	fmt.Println("\n\nHello, Damodar!")
	solve(4, nil)
}
