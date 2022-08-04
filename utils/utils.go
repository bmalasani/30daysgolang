package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

type empty interface{}

type basic interface {
	~int | ~string | ~float32
}

func Includes[T basic](arr []T, value T) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}

func WriteLog(msg ...string) {
	var buff bytes.Buffer
	logger := log.New(&buff, "MEB_", 3)
	logger.Println(msg)
	fmt.Printf("buff: %v\n", buff.String())
}

func ReadFromConsole(msg string) string {
	fmt.Println(msg)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
