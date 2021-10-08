package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n)		// int 입력받음
	if err != nil {
		stdin.ReadString('\n')		// 에러 발생시 입력스트림 비움
	}
	return n, err
}

func main() {

	// 시간 값을 random 시드 값으로 설정
	rand.Seed(time.Now().UnixNano())
	
	r := rand.Intn(100)				// 랜덤값 생성
	cnt := 1
	for {
		fmt.Printf("숫자값을 입력하세요:")
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else if n == 00 {
			fmt.Println("종료합니다.")
			break
		} else {					// 숫자 비교
			if n > r {
				fmt.Println("입력된  숫자가 더 큽니다.")
			} else if n < r {
				fmt.Println("입력된  숫자가 더 작습니다.")
			} else {
				fmt.Println("숫자를 맞췄습니다. (시도 횟수 : ", cnt, " )")
				break
			}
			cnt++
		}
	}
}
