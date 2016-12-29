package main

import (
	"fmt"
	"time"
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

func solve(number int, cur_list []int, exit_on_find bool) []([]int){
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

				if exit_on_find {
					return output
				}
			}else {
				results := solve(number-1, new_list, exit_on_find)
				if len(results) > 0 {
					output = append(output, results...)
					if exit_on_find {
						return output
					}
				}
			}

		}


	}

	return output

}

func solveWithGraphics(num int, display_results bool){
	debugPrint("Solving for ", num)
	startTime := time.Now()
	result := solve(num, nil, display_results)
	duration := time.Since(startTime)

	debugPrint("Calculated in", duration, "seconds")
	debugPrint(len(result), "permutations found")
	if display_results{
		debugPrint(result)
		debugPrint()
	}
}
func checkForExistenceOnly(start int, end int){
	for i:= start; i<=end; i++{
		rem := i % 4;
		if rem == 0 || rem == 3 {
			if i == 36 || i == 39 {

			}else{
				solveWithGraphics(i, true);

			}


		}
	}
}
func solveForRange(min int, max int){
	for i:= min; i <= max; i++{
		if i % 4 == 0 || i % 4 == 3{
			solveWithGraphics(i, false)
		}else {
				debugPrint("Skipping for ", i)
		}
	}
}

func main(){
	//fmt.Println("\n\nHello, Damodar!")

	solveForRange(12, 15)
	//checkForExistenceOnly(1, 100)
}
