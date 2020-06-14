package main

import (
	"fmt"
	"os"
	"unicode"
)

func at(s string, i int) rune {
	return []rune(s)[i]
}

func strtoi(input string, pos int) (int, int) {
	var res = 0
	for pos < len(input) && unicode.IsDigit(at(input, pos)) {
		res = res*10 + int(at(input, pos)-'0')
		pos++
	}
	return res, pos
}

func run() int {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "引数の個数が正しくありません。\n")
		return 1
	}

	fmt.Printf(".intel_syntax noprefix\n")
	fmt.Printf(".global main\n")
	fmt.Printf("main:\n")

	var input = os.Args[1]
	num, pos := strtoi(input, 0)

	fmt.Printf("  mov rax, %d\n", num)
	for pos < len(input) {
		if at(input, pos) == '+' {
			pos++
			num, pos = strtoi(input, pos)
			fmt.Printf("  add rax, %d\n", num)
			continue
		}

		if at(input, pos) == '-' {
			pos++
			num, pos = strtoi(input, pos)
			fmt.Printf("  sub rax, %d\n", num)
			continue
		}

		fmt.Fprintf(os.Stderr, "予期しない文字です: %c\n", at(input, pos))
		return 1
	}

	fmt.Printf("  ret\n")

	return 0
}

func main() {
	var status = run()
	os.Exit(status)
}
