package main

import (
   "fmt"
   "github.com/google/uuid"
)

func main() {

   test_uuid := "a617516f-19bf-4e3b7b923-03bf58a8b3f7"
   fmt.Println("test uuid >> ", test_uuid)

   if test_uuid == "" {
      fmt.Println("빈 값")
   } else if _, uuidErr := uuid.Parse(test_uuid); uuidErr != nil {
      fmt.Println("uuid 타입이 아님 >> ", uuidErr)
   } else {
      fmt.Println("uuid 타입 ㅊㅊ")
   }

}
