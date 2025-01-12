package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func echo(c net.Conn) {
	defer c.Close()
	for {
		buf := make([]byte, 1024)
		n, err := c.Read(buf)
		if err != nil {
			return
		}

		data := buf[:n]
		fmt.Print("Server got: ", string(data))
		_, err = c.Write(data)
		if err != nil {
			fmt.Println("Write:", err)
			return
		}
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("Need socket path")
		return
	}
	socketPath := flag.Arg(0)

	_, err := os.Stat(socketPath)
	if err == nil {
		fmt.Println("Deleting existing", socketPath)
		err := os.Remove(socketPath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	l, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	defer l.Close()

	for {
		fd, err := l.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			return
		}
		go echo(fd)
	}
}
