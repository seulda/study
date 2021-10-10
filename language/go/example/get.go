package main
 
import (
    "fmt"
    "io/ioutil"
    "net/http"
)
 
 
// http.Get() 메서드는 쉽게 웹 페이지나 웹 데이타들 가져오는데 사용
// 한 문장으로 HTTP GET을 호출하는 장점이 있지만
// Request시 헤더를 추가한다던가, Request 스트림을 추가하는 것과 같은 세밀한 컨트롤을 할 수 없는 단점이 있음
// func main() {
//     // GET 호출
//     resp, err := http.Get("http://cobye.site/info")
//     if err != nil {
//         panic(err)
//     }
 
//     defer resp.Body.Close()
 
//     // 결과 출력
//     data, err := ioutil.ReadAll(resp.Body)
//     if err != nil {
//         panic(err)
//     }
//     fmt.Printf("%s\n", string(data))
// }



// Request시 보다 많은 컨트롤이 필요하다면, Request 객체를 (NewRequest 생성자를 통해) 직접 생성해서 http.Client 객체를 통해 호출하면 됨
// http.NewRequest() 생성자를 통해 Request 객체를 생성하고, 여기에 임의의 헤더를 추가하고, http.Client 객체를 통해 호출하는 코드
// HTTP 호출결과는 Response 객체. 이 객체의 Body 필드를 통해 실제 Request 결과를 가져올 수 있음
func main() {
	// Request 객체 생성
    req, err := http.NewRequest("GET", "http://csharp.tips/feed/rss", nil)
    if err != nil {
        panic(err)
    }
 
    //필요시 헤더 추가 가능
    req.Header.Add("User-Agent", "Crawler")
 
    // Client객체에서 Request 실행
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
 
    // 결과 출력
    bytes, _ := ioutil.ReadAll(resp.Body)
    str := string(bytes) //바이트를 문자열로
    fmt.Println(str)
}
