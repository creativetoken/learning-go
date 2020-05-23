package main

import (
	"fmt"
	"net"
)

func main() {

	var data []byte
	u, _ := net.ListenPacket("UDP", "127.0.0.1:53")
	fmt.Println("first step ran ")
	_, _, err := u.ReadFrom(data)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(data)

}
