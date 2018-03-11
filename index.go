package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	choclateFeast()
}

func choclateFeast() {
	var t int
	fmt.Scanf("%d", &t)
	for i := 0; i < t; i++ {
		var money, choclateCount, wrappers int
		fmt.Scanf("%d %d %d", &money, &choclateCount, &wrappers)
		result := ClacChoclateCount(money, choclateCount, wrappers)
		fmt.Println(result)
	}
}

func ClacChoclateCount(money, price, wrappers int) int {
	totalChoc := money / price
	totalWrappers := totalChoc
	for wrappers <= totalWrappers {
		totalChoc += totalWrappers / wrappers
		totalWrappers = totalWrappers/wrappers + totalWrappers%wrappers
	}
	return totalChoc
}

func strongPassword(password string) int {
	result := 0
	numbers := "0123456789"
	lower_case := "abcdefghijklmnopqrstuvwxyz"
	upper_case := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special_characters := "!@#$%^&*()-+"

	if !strings.ContainsAny(password, numbers) {
		result += 1
	}

	if !strings.ContainsAny(password, lower_case) {
		result += 1
	}

	if !strings.ContainsAny(password, upper_case) {
		result += 1
	}

	if !strings.ContainsAny(password, special_characters) {
		result += 1
	}
	if len(password) < 6 {
		if (result + len(password)) < 6 {
			result = (6 - len(password))
		}
	}
	return result
}

func testStrongPass() {
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	result := strongPassword(text)
	fmt.Println(result)
}
func test() {
	testStrongPass()
}
