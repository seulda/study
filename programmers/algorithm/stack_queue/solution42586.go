func solutionNonQueue(progresses []int, speeds []int) []int {
    result := []int{}
	deploy := make([]int, len(progresses))
	deployPoint := 0
	for deployPoint < len(progresses) {
		success := false
		for i := 0; i < len(progresses); i++ {
			if progresses[i] == 0 {
				continue
			}
			progresses[i] += speeds[i]
			if progresses[i] >= 100 {
				success = true
				deploy[i] = 1
				progresses[i], speeds[i] = 0, 0
			}
		}
		if success {
			cnt := 0
			for deployPoint < len(progresses) && deploy[deployPoint] == 1 {
				deploy[deployPoint] = -1
				cnt++
				deployPoint++
			}
			if cnt > 0 {
				result = append(result, cnt)
			}
		}
	}
	return result
}
