package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"tcp_test_ok/mod"
)

var (
	service       string
	port          string
	download_path string
	file          string
	ip            string
)

func checkPath(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}
func pathRule(path string) string {
	if strings.Contains(path, ":") {
		path = strings.Replace(path, "/", "\\", -1)
		if strings.HasSuffix(path, "\\") {
			return path
		} else {
			return path + "\\"
		}
	} else {
		path = strings.Replace(path, "\\", "/", -1)
		if strings.HasSuffix(path, "/") {
			return path
		} else {
			return path + "/"
		}
	}
}

func main() {
	flag.StringVar(&port, "port", "", "Listen/connect port")
	flag.StringVar(&download_path, "download_path", "", "Server download path")
	flag.StringVar(&file, "file", "", "Client file path")
	flag.StringVar(&ip, "ip", "", "Server ip")
	flag.StringVar(&service, "service", "", "Service mod")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `example :
	server mod: ./down -service server -port 8888  -download_path c:/windows/temp/
	client mod: ./down -service client -ip 8.8.8.8 -port 8888 -file c:/web/www/config.php	

`)
		flag.PrintDefaults()
	}

	flag.Parse()
	switch service {
	case "server":
		if port != "" && download_path != "" {
			if checkPath(download_path) {
				mod.Start_server(pathRule(download_path), port)
			} else {
				fmt.Println("[!] Path wrong! example:c:\\windows\\temp\\ or /tmp/")
				flag.Usage()
			}
		} else {
			flag.Usage()
		}
	case "client":
		if ip != "" && port != "" && file != "" {
			mod.Start_client(ip, port, file)
		} else {
			flag.Usage()
		}

	default:
		fmt.Println("please chose a service")
		flag.Usage()
		os.Exit(-1)
	}

}
