package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const otherWord = "*"

var transforms = []string{
	otherWord,
	otherWord,
	otherWord,
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	"get" + otherWord,
	"go" + otherWord,
	"lets" + otherWord,
}

func main() {
	// ナノ秒にすることで必ず違くなる
	rand.Seed(time.Now().UTC().UnixNano())
	// 標準入力のストリームからデーターを読み込むbufio.Scanner
	s := bufio.NewScanner(os.Stdin)
	// 読み込みデータの有無を調べる
	for s.Scan() {
		// スライスtransformsの中から項目がランダムに呼び出される
		t := transforms[rand.Intn(len(transforms))]
		// 置き換えの結果がデフォルトの標準入力に書き出される
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))

	}
}
