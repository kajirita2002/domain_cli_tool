package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func exists(domain string) (bool, error) {
	const whoisServer string = "com.whois-servers.net"
	// 指定されたサーバーのポート43に対して、net.Dialを使って接続を開く
	conn, err := net.Dial("tcp", whoisServer+":43")
	if err != nil {
		return false, err
	}
	defer conn.Close()
	// ドメイン名と復帰と改行を送信
	conn.Write([]byte(domain + "\r\n"))
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		if strings.Contains(strings.ToLower(scanner.Text()), "no match") {
			return false, nil
		}
	}
	return true, nil
}



// 返された真偽値を理解しやすいようにする
var marks = map[bool]string{true: "○", false: "×"}

func main() {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		domain := s.Text()
		fmt.Print(domain, " ")
		exist, err := exists(domain)
		if err != nil {
			log.Fatalln(err)
		}
		// 使われていないものを○にする
		fmt.Println(marks[!exist])
		// サーバーの負荷の上昇を避けるために1秒休止している
		time.Sleep(1 * time.Second)
	}
}