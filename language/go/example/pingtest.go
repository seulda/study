package main

import (
	"github.com/go-ping/ping"
	"fmt"
)

func PingTest() {
  
	pinger, err := ping.NewPinger("www.google.com") // "142.250.66.100"
	pinger.SetPrivileged(true)
  
	if err != nil {
		fmt.Println("error check 1 : ", err.Error())
		panic(err)
	}
  
	pinger.Count = 1
  
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		fmt.Println("error check 2 : ", err.Error())
		panic(err)
	}
  
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats

	fmt.Println("\n@@@@@ ping test result : ", stats)
}
