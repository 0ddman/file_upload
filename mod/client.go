package mod

import (
	"fmt"
	"net"
	"os"
)

func Start_client(ip string, port string, file_all_path string) {
	result, file_path := Check_file(file_all_path)
	if !result {
		return
	}
	fmt.Println("[*] Trance file path:\n" + file_all_path)
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	conn.Write([]byte(file_path))
	buf := make([]byte, 2048)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	revData := string(buf[:n])
	if revData == "ok" {
		Send_file(file_all_path, conn)
	}
}
func Send_file(path string, conn net.Conn) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	for {
		tmp := make([]byte, 2048)
		lens, err := f.Read(tmp)
		if err != nil && lens == 0 {
			fmt.Println("[!] File transfer completed!")
			break
		}
		conn.Write(tmp[:lens])
	}
}

func Check_file(file string) (bool, string) {
	fileInfo, err := os.Stat(file)
	if err != nil {
		fmt.Println("[!] The system cannot find the file specified.")
		return false, ""
	}
	filename := fileInfo.Name()
	return true, filename
}
