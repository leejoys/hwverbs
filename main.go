package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	listener, err := net.Listen("tcp4", ":12345")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	txtFile, err := os.OpenFile("./fables.txt", os.O_RDONLY, 0555)
	if err != nil {
		log.Fatal(err)
	}
	defer txtFile.Close()

	reader := bufio.NewReader(txtFile)
	fables := make(map[int]string)
	i := 0
	for {
		fable, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		i++
		fables[i] = string(fable)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go fabler(conn, fables)
	}
}

func fabler(c net.Conn, fables map[int]string) {
	defer c.Close()
	for {
		for _, fable := range fables {
			_, err := c.Write([]byte(fable))
			if err != nil {
				return
			}
			time.Sleep(time.Second * 3)
		}

	}
}
