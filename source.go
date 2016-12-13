package main

import (
	"fmt"
)


func getEmptyIndices(list []int) []int{
	cur_index := 0

	z := make([]int, len(list))

	for index,x := range list {
		if x == 0 {
			z[cur_index] = index
			cur_index++
		}
	}

	return z[:cur_index]
}

func debugPrint(args... interface{}){
	for _, element := range args {
		fmt.Print(element)
		fmt.Print("\t")
	}
	fmt.Println()
}

func solve(number int, cur_list []int) []([]int){
	if cur_list == nil {
		cur_list = make([]int, number*2)
	}

	var output []([]int)

	emptyIndices := getEmptyIndices(cur_list)

	if len(emptyIndices) == 0{ return []([]int){cur_list}}

	for _,index := range emptyIndices {
		next_index := index + number + 1

		if next_index < len(cur_list) && cur_list[next_index] == 0{

			new_list := make([]int, len(cur_list))
			copy(new_list, cur_list)

			new_list[index] = number
			new_list[next_index] = number

			if number == 1{
				output = append(output, new_list)
			}else {
				results := solve(number-1, new_list)
				if len(results) > 0 {
					output = append(output, results...)
				}
			}

		}else {
			continue
		}


	}

	return output

}

func solveForRange(min int, max int){
	for i:= min; i <= max; i++{
		debugPrint("Solving for ", i)
		result := solve(i, nil)
		debugPrint(len(result), "permutations found\n")
	}
}
func main(){
	fmt.Println("\n\nHello, Damodar!")

	solveForRange(10, 15)
}
