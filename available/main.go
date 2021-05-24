package main

import (
	"bufio"
	"net"
	"strings"
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