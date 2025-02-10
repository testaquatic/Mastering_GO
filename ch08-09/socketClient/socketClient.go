package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("Need socket path")
		return
	}
	socketPath := flag.Arg(0)

	c, err := net.Dial("unix", socketPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		_, err = c.Write([]byte(text))
		if err != nil {
			fmt.Println("Write:", err)
			break
		}

		buf := make([]byte, 256)

		n, err := c.Read(buf)
		if err != nil {
			fmt.Println(err, n)
			return
		}
		fmt.Print("Read:", string(buf[:n]))

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("Exiting UNIX domain socket client!")
			return
		}

		time.Sleep(5 * time.Second)
	}
}
