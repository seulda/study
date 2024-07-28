
type Queue []int

func NewQueue() *Queue {
    return &Queue{}
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue (data int) {
	*q = append(*q, data)
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		return 0
	}
	data := (*q)[0]
	*q = (*q)[1:]
	return data
}

func solution(priorities []int, location int) int {
    readyQueue := NewQueue()
	for _, priority := range priorities {
		readyQueue.Enqueue(priority)
	}
	chkLocation := location
	dequeCnt := 0
	for {
		chk := false
		dequeData := readyQueue.Dequeue()
		for _, temp := range *readyQueue {
			if dequeData < temp {
				chk = true
				break
			}
		}

		if chk {
			// 이후 대기프로세스들 중 현재 큐[0] 보다 우선순위가 높은게 있음 = 다시 돈다
			readyQueue.Enqueue(dequeData)

		} else {
			// 이후 대기프로세스들 중 현재 큐[0] 보다 우선순위가 높은게 없음 = 디큐할 상황
			dequeCnt++

			// 체크할 프로세스 순위인지 검사
			if chkLocation == 0 {
				break
			}
		}

		chkLocation--
		if chkLocation < 0 {
			chkLocation = len(*readyQueue) - 1
		}
	}

    return dequeCnt
}
