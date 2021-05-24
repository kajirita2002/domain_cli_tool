package main

import (
	"log"
	"os"
	"os/exec"
)

var cmdChain = []*exec.Cmd{
	exec.Command("lib/synonyms"),
	exec.Command("lib/sprinkle"),
	exec.Command("lib/coolify"),
	exec.Command("lib/domainify"),
	exec.Command("lib/available"),
}

func main() {
	// synonymsの標準入力のストリームをdomainfinderにとっての標準入力に接続
	cmdChain[0].Stdin = os.Stdin
	// availableにとっての標準出力のストリームをdomainfinderにとっての標準出力に接続
	cmdChain[len(cmdChain)-1].Stdout = os.Stdout


	for i := 0; i < len(cmdChain)-1; i++ {
		thisCmd := cmdChain[i]
		nextCmd := cmdChain[i+1]
		// 標準出力
		stdout, err := thisCmd.StdoutPipe()
		if err != nil {
			log.Panicln(err)
		}
		// 次のコマンドの標準入力は現在の標準出力になる
		nextCmd.Stdin = stdout
	}

	for _, cmd := range cmdChain {
		// 実行する
		if err := cmd.Start(); err != nil {
			log.Panicln(err)
		} else {
			// コマンドのプロセスを終了させる (domainfinderプログラム終了時に全てのコマンドが終了している)
			defer cmd.Process.Kill()
		}
	}

	for _, cmd := range cmdChain {
		// startで始めたコマンドを待つことができる(Startとセットで実装)
		if err := cmd.Wait(); err != nil {
			log.Panicln(err)
		}
	}
}