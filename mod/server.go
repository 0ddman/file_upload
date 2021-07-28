package mod

import (
	"fmt"
	"io"
	"net"
	"os"
)

func Start_server(download_path string, port string) {
	fmt.Println("[*] Listen up to-->" + port)
	fmt.Println("[*] File download to path-->" + download_path)
	listen, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	Handler(conn, download_path)
}

func Handler(conn net.Conn, download_path string) {
	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	addr := conn.RemoteAddr().String()
	fmt.Println("[!] " + addr + " connected")
	file_name := string(buf[:n])
	fmt.Println("[*] Transfer filename:" + " " + string(buf[:n]))
	conn.Write([]byte("ok"))
	Write_file(download_path, file_name, addr, conn)
}

func Write_file(path string, file_name string, addr string, conn net.Conn) {
	f, err := os.OpenFile(path+file_name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	defer f.Close()
	_, err1 := io.Copy(f, conn)
	if err1 == nil {
		fmt.Println("[!]" + addr + " disconnected")
		return
	}
}
