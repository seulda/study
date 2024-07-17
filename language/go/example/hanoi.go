package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 기둥 번호의 합 (기둥 3개 기준)
const pillar int = 6

// 원반을 움직인 총 횟수
var moveCnt int = 0

func Hanoi() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("(하노이의 탑) 원반 개수 : ")

	if inputText, _, err := reader.ReadLine(); err == nil {
		discus, _ := strconv.Atoi(string(inputText))
		fmt.Println("입력 값 : ", discus)
		// MoveInt(discus, 1, 3)
		MoveString(discus, "A", "B", "C")

	} else {
		fmt.Println("console input error : ", err)
	}

	fmt.Println("이동 횟수 : ", moveCnt)
}

// no개의 원반을 x번 기둥에서 y번 기둥으로 옮김
func MoveInt(no int, x int, y int) {

	middle := pillar - x - y

	if no > 1 {
		MoveInt(no-1, x, middle)
	}

	moveCnt++
	fmt.Println(" (", moveCnt, " 회) ", "원반[", fmt.Sprint(no), "]을 ", fmt.Sprint(x), " 기둥에서 ", fmt.Sprint(y), " 기둥으로 옮김")

	if no > 1 {
		MoveInt(no-1, middle, y)
	}
}
func MoveString(no int, from string, by string, to string) {

	if no > 1 {
		MoveString(no-1, from, to, by)
	}

	moveCnt++
	fmt.Println(" (", moveCnt, " 회) ", "원반[", fmt.Sprint(no), "]을 ", from, " 기둥에서 ", to, " 기둥으로 옮김")

	if no > 1 {
		MoveString(no-1, by, from, to)
	}
}
