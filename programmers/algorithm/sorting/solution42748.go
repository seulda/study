import "sort"

func solution(array []int, commands [][]int) []int {    
    result := []int{}
	for _, command := range commands {
		arr := make([]int, command[1]-command[0]+1)
		copy(arr, array[command[0]-1:command[1]])
		sort.Ints(arr)
		result = append(result, arr[command[2]-1])
	}
	return result
}
