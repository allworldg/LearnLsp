package main

import (
	"bufio"
	"fmt"
	"golsp/rpc"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)
	logger := getLogger("/Users/allworldg/Dev/go-lsp/log/log.txt")
	logger.Println("fine the first log is created")
	for scanner.Scan() {
		handleMessage(scanner.Text())
	}
}
func handleMessage(_ any) {}
func getLogger(filePath string) *log.Logger {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println(err)
		panic("u do not give a right filePath")
	}
	logger := log.New(file, "golsp ", log.Ldate|log.Ltime|log.Lshortfile)
	return logger

}
