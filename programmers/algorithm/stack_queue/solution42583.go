type Truck struct {
	Weight int
	Time   int
}

type Queue struct {
	Value   []*Truck
	Weights int
}

func NewQueue() *Queue {
	return &Queue{Value: nil, Weights: 0}
}

func (q *Queue) IsEmpty() bool {
	return len(q.Value) == 0
}

func (q *Queue) Size() int {
	return len(q.Value)
}

func (q *Queue) Enqueue(data int) {
	q.Value = append(q.Value, &Truck{Weight: data, Time: 0})
	q.Weights += data
}

func (q *Queue) Dequeue() *Truck {
	if q.IsEmpty() {
		return nil
	}
	data := q.Value[0]
	q.Value = q.Value[1:]
	q.Weights -= data.Weight
	return data
}

func solution(pbridge_length int, weight int, truck_weights []int) int {

	cnt := 1

	standby := NewQueue()
	for _, truck := range truck_weights {
		standby.Enqueue(truck)
	}
	readyTruck := 0

	onBridge := NewQueue()

	for !onBridge.IsEmpty() || !standby.IsEmpty() || readyTruck != 0{

		if readyTruck == 0 && !standby.IsEmpty() {
			readyTruck = standby.Dequeue().Weight
		}

		// 다리 위 모든 트럭의 무게와 수를 확인 후 다리로 진입
		if (onBridge.Weights + readyTruck <= weight) && (onBridge.Size() < pbridge_length) && readyTruck != 0 {
			onBridge.Enqueue(readyTruck)
			readyTruck = 0
		}
		
		// 1초 진행
		for _, onBridgeTruck := range onBridge.Value {
			onBridgeTruck.Time++
		}
		cnt++

		// 트럭이 다리를 건넜는지 확인
		for _, onBridgeTruck := range onBridge.Value {
			if onBridgeTruck.Time == pbridge_length {
				onBridge.Dequeue()
			} else {
				break
			}
		}
	}

	return cnt
}
