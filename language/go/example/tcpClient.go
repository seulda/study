package main

import (
	"fmt"
	"net"
	// "time"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:9000")
    if err != nil {
        fmt.Println("Faield to Dial : ", err)
    }

	// client addr
	fmt.Println("@@@@@ LocalAddr() print : ", conn.LocalAddr())
	fmt.Println("@@@@@ LocalAddr().Network() print : ", conn.LocalAddr().Network()) // 연결 요청 방식
	fmt.Println("@@@@@ LocalAddr().String() print : ", conn.LocalAddr().String())   // client addr string

	// server addr
	fmt.Println("@@@@@ RemoteAddr() print : ", conn.RemoteAddr())
	fmt.Println("@@@@@ RemoteAddr() print : ", conn.RemoteAddr().Network()) // 연결 요청 방식
	fmt.Println("@@@@@ RemoteAddr() print : ", conn.RemoteAddr().String())  // server addr string

	fmt.Println("@@@@@ Close() print : ", conn.Close())
	// client -> server , use of closed network connection
	fmt.Println("@@@@@ Close().Error() print : ", conn.Close().Error())

    defer conn.Close()

    // go func(c net.Conn) {
    //     send := "Hello"
    //     for {
    //         _, err = c.Write([]byte(send))
    //         if err != nil {
    //             fmt.Println("Failed to write data : ", err)
    //             break;
    //         }
 
    //         time.Sleep(1 * time.Second)
    //     }
    // }(conn)
 
    // go func(c net.Conn) {
    //     recv := make([]byte, 4096)
 
    //     for {
    //         n, err := c.Read(recv)
    //         if err != nil {
    //             fmt.Println("Failed to Read data : ", err)
    //             break
    //         }
 
    //         fmt.Println(string(recv[:n]))
    //     }
    // }(conn)
 
    // fmt.Scanln()
}

