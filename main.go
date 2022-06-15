package main

import (
	"io"
	"log"
	"net"
)

func main() {
	SocketServer()
}

func SocketServer() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Println(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go connHandler(conn)
	}
}

func connHandler(conn net.Conn) {
	//수신버퍼 설정
	recvBuf := make([]byte, 4096)
	//값이 들어오면 읽기
	for {
		n, err := conn.Read(recvBuf)
		if nil != err {
			if io.EOF == err {
				log.Println(err)
				return
			}
			log.Println(err)
			return
		}
		//클라이언트가 던진 값을 다시 클라이언트에게 전달
		if 0 < n {
			data := recvBuf[:n]
			log.Println(string(data))
			_, err = conn.Write(data[:n])
			if err != nil {
				log.Println(err)
				return
			}
		}
		defer conn.Close()
	}
}
