package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func promptAndGetStr(prompt string) string {
	fmt.Printf("%s: ", prompt)

	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	str = strings.TrimSuffix(str, "\n")
	str = strings.TrimSuffix(str, "\r")

	return str
}
