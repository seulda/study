import (
	"bytes"
	"sort"
	"strconv"
)

func solution(numbers []int) (result string) {
	sort.Slice(numbers, func(i, j int) bool {
		prev := strconv.Itoa(numbers[i])
		next := strconv.Itoa(numbers[j])
		return prev+next > next+prev
	})

	if numbers[0] == 0 {
		result = "0"

	} else {
		var buf bytes.Buffer
		for _, num := range numbers {
			buf.WriteString(strconv.Itoa(num))
		}
		result = buf.String()
	}

	return
}
