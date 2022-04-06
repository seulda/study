package main

import (
   "fmt"
   "reflect"
   "strconv"
   "time"
   "regexp"
)

func main() {
   var x, y = 1, "2"
   var z string

   fmt.Println(x, y)
   fmt.Println(z)
   fmt.Println(reflect.TypeOf(x), reflect.TypeOf(y))
   fmt.Println(reflect.TypeOf(z))

   if reflect.TypeOf(z) == reflect.TypeOf(1) {
      fmt.Println("분류가 되는가")
   } else {
      fmt.Println("아닌가")
   }

   ////////////////////////////////////////////////////////////////
  
   now := time.Now()
   date1 := now.AddDate(0, 0, 1).Format(time.RFC3339)
   date2 := now.Add(time.Hour * 1).Format(time.RFC3339)
   fmt.Println("포맷 한거 > ", now)
   fmt.Println("AddDate > ", date1)
   fmt.Println("Add > ", date2)

   d, _ := strconv.Atoi("1587870895")
   unixTime := time.UnixMilli(int64(d))
   fmt.Println("유닉스타임 형 변환 > ", unixTime)
   unixTime2 := time.UnixMilli(int64(d)).Format(time.RFC3339)
   fmt.Println("유닉스타임 형 변환 후 포맷 > ", unixTime2)

  
   ////////////////////////////////////////////////////////////////

   var matches = regexp.MustCompile(`text1:([A-Za-z-+\/]+);text2,(.+)`)
   matchesArray := matches.FindStringSubmatch(`text1:i+///////////!!@12345;text2,Here@@441w5@@@!!nnKK]2\set`)

   for i, ma := range matchesArray {
      fmt.Printf("%d번째 : %s\n", i, ma)
   }
  
}
