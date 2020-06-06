package main

import (
	"fmt"
	"os"
	"strconv"
)

func run() int {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "引数の個数が正しくありません。\n")
		return 1
	}

	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".global main\n")
	fmt.Printf("main:\n")
	var n, err = strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "引数が整数ではありません。\n")
		return 1
	}
	fmt.Printf("  mov rax, %d\n", n)
	fmt.Printf("  ret\n")

	return 0
}

func main() {
	var status = run()
	os.Exit(status)
}
