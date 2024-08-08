import "sort"

func solution(lottos []int, win_nums []int) []int {
    sort.Ints(lottos)
	sort.Ints(win_nums)

	containCnt := 0
	zeroCnt := 0
	i, j := 0, 0
	for i < len(lottos) && j < len(lottos) {
		if lottos[j] == 0 {
			zeroCnt++
		}
		if lottos[j] > win_nums[i] {
			i++
		} else if lottos[j] < win_nums[i] {
			j++
		} else if lottos[j] == win_nums[i] {
			containCnt++
			i++
			j++
		}
	}

	result := []int{}
	if (containCnt < 2 && zeroCnt == 0) || (containCnt == 0 && zeroCnt < 2) {
		result = append(result, 6, 6)
	} else if zeroCnt == 6 {
		result = append(result, 1, 6)
	} else {
		winCnt := 7 - containCnt
		hiddenCnt := winCnt - zeroCnt
		result = append(result, hiddenCnt, winCnt)
	}

    return result
}
