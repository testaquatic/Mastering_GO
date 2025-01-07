package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("Please provide a host:port string")
		return
	}
	CONNECT := flag.Arg(0)

	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	fmt.Printf("The UDP server is %s\n", c.RemoteAddr().String())

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		data := []byte(text + "\n")
		_, err = c.Write(data)

		if strings.TrimSpace(string(data)) == "STOP" {
			fmt.Println("Exiting UDP client!")
			return
		}

		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Reply: %s\n", string(buffer[0:n]))
	}
}
